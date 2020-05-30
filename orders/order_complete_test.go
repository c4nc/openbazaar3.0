package orders

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/cpacia/openbazaar3.0/database"
	"github.com/cpacia/openbazaar3.0/events"
	"github.com/cpacia/openbazaar3.0/models"
	npb "github.com/cpacia/openbazaar3.0/net/pb"
	"github.com/cpacia/openbazaar3.0/orders/pb"
	"github.com/cpacia/openbazaar3.0/orders/utils"
	iwallet "github.com/cpacia/wallet-interface"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"reflect"
	"testing"
)

func TestOrderProcessor_processOrderCompleteMessage(t *testing.T) {
	op, teardown, err := newMockOrderProcessor()
	if err != nil {
		t.Fatal(err)
	}
	defer teardown()

	vendorPriv, vendorPub, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	pubkeyBytes, err := crypto.MarshalPublicKey(vendorPub)
	if err != nil {
		t.Fatal(err)
	}
	vendor, err := peer.IDFromPublicKey(vendorPub)
	if err != nil {
		t.Fatal(err)
	}

	buyerPriv, buyerPub, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	buyerPubkeyBytes, err := crypto.MarshalPublicKey(buyerPub)
	if err != nil {
		t.Fatal(err)
	}
	buyer, err := peer.IDFromPublicKey(buyerPub)
	if err != nil {
		t.Fatal(err)
	}
	op.identity = vendor

	orderID := "1234"

	chaincode := make([]byte, 32)
	rand.Read(chaincode)

	ratingMaster, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		t.Fatal(err)
	}
	ratingKeys, err := utils.GenerateRatingPrivateKeys(ratingMaster, 1, chaincode)
	if err != nil {
		t.Fatal(err)
	}

	vendorSig := &pb.RatingSignature{
		Slug:      "abc",
		RatingKey: ratingKeys[0].PubKey().SerializeCompressed(),
	}

	ser, err := proto.Marshal(vendorSig)
	if err != nil {
		t.Fatal(err)
	}
	sig, err := vendorPriv.Sign(ser)
	if err != nil {
		t.Fatal(err)
	}
	vendorSig.VendorSignature = sig

	var (
		vendorPeerID   = vendor.Pretty()
		vendorHandle   = "abc"
		smallImageHash = "aaaa"
		tinyImageHash  = "bbbb"
	)

	hashedRatingKey := sha256.Sum256(ratingKeys[0].PubKey().SerializeCompressed())
	buyerSig, err := buyerPriv.Sign(hashedRatingKey[:])
	if err != nil {
		t.Fatal(err)
	}

	orderComplete := &pb.OrderComplete{
		Ratings: []*pb.Rating{
			{
				VendorSig: vendorSig,
				VendorID: &pb.ID{
					PeerID: vendorPeerID,
					Handle: vendorHandle,
					Pubkeys: &pb.ID_Pubkeys{
						Identity: pubkeyBytes,
					},
				},
				Timestamp: ptypes.TimestampNow(),
				BuyerID: &pb.ID{
					PeerID: buyer.Pretty(),
					Handle: "aaa",
					Pubkeys: &pb.ID_Pubkeys{
						Identity: buyerPubkeyBytes,
					},
				},
				BuyerName: "Ernie",
				BuyerSig:  buyerSig,

				Overall:         5,
				DeliverySpeed:   4,
				Description:     3,
				CustomerService: 2,
				Quality:         1,
				Review:          "sucked",
			},
		},
	}
	ser, err = proto.Marshal(orderComplete.Ratings[0])
	if err != nil {
		t.Fatal(err)
	}
	hashed := sha256.Sum256(ser)
	ratingSig, err := ratingKeys[0].Sign(hashed[:])
	if err != nil {
		t.Fatal(err)
	}
	orderComplete.Ratings[0].RatingSignature = ratingSig.Serialize()

	completeAny, err := ptypes.MarshalAny(orderComplete)
	if err != nil {
		t.Fatal(err)
	}

	orderMsg := &npb.OrderMessage{
		OrderID:     orderID,
		MessageType: npb.OrderMessage_ORDER_COMPLETE,
		Message:     completeAny,
	}

	orderOpen := &pb.OrderOpen{
		Listings: []*pb.SignedListing{
			{
				Listing: &pb.Listing{
					Slug: "abc",
					VendorID: &pb.ID{
						PeerID: vendorPeerID,
						Handle: vendorHandle,
						Pubkeys: &pb.ID_Pubkeys{
							Identity: pubkeyBytes,
						},
					},
					Item: &pb.Listing_Item{
						Images: []*pb.Listing_Item_Image{
							{
								Small: smallImageHash,
								Tiny:  tinyImageHash,
							},
						},
					},
				},
			},
		},
		Payment: &pb.OrderOpen_Payment{
			Coin:      iwallet.CtMock,
			Chaincode: hex.EncodeToString(chaincode),
		},
		RatingKeys: [][]byte{
			ratingKeys[0].PubKey().SerializeCompressed(),
		},
		Items: []*pb.OrderOpen_Item{
			{
				ListingHash: "1234",
			},
		},
		BuyerID: &pb.ID{
			PeerID: buyer.Pretty(),
			Handle: "aaa",
			Pubkeys: &pb.ID_Pubkeys{
				Identity: buyerPubkeyBytes,
			},
		},
	}

	tests := []struct {
		setup         func(order *models.Order) error
		expectedError error
		expectedEvent interface{}
	}{
		{
			// Normal case where order open exists.
			setup: func(order *models.Order) error {
				order.SetRole(models.RoleVendor)
				order.ID = models.OrderID(orderID)
				return order.PutMessage(&npb.OrderMessage{
					Signature: []byte("abc"),
					Message:   mustBuildAny(orderOpen),
				})
			},
			expectedError: nil,
			expectedEvent: &events.OrderCompletion{
				OrderID: orderID,
				Thumbnail: events.Thumbnail{
					Tiny:  tinyImageHash,
					Small: smallImageHash,
				},
				BuyerHandle: "aaa",
				BuyerID:     buyer.Pretty(),
			},
		},
		{
			// Order cancel already exists.
			setup: func(order *models.Order) error {
				order.SerializedOrderCancel = []byte{0x00}
				return nil
			},
			expectedError: ErrUnexpectedMessage,
			expectedEvent: nil,
		},
		{
			// Duplicate order complete.
			setup: func(order *models.Order) error {
				return order.PutMessage(&npb.OrderMessage{
					Signature:   []byte("abc"),
					Message:     completeAny,
					MessageType: npb.OrderMessage_ORDER_COMPLETE,
				})
			},
			expectedError: nil,
			expectedEvent: nil,
		},
		{
			// Duplicate but different.
			setup: func(order *models.Order) error {
				a := proto.Clone(orderComplete)
				a.(*pb.OrderComplete).Ratings[0].Review = "fasdfad"
				return order.PutMessage(&npb.OrderMessage{
					Signature:   []byte("abc"),
					Message:     mustBuildAny(a),
					MessageType: npb.OrderMessage_ORDER_COMPLETE,
				})
			},
			expectedError: ErrChangedMessage,
			expectedEvent: nil,
		},
		{
			// Out of order.
			setup: func(order *models.Order) error {
				order.SerializedOrderOpen = nil
				return nil
			},
			expectedError: nil,
			expectedEvent: nil,
		},
	}

	for i, test := range tests {
		order := &models.Order{}
		if err := test.setup(order); err != nil {
			t.Errorf("Test %d setup error: %s", i, err)
			continue
		}
		err := op.db.Update(func(tx database.Tx) error {
			event, err := op.processOrderCompleteMessage(tx, order, buyer, orderMsg)
			if err != test.expectedError {
				return fmt.Errorf("incorrect error returned. Expected %t, got %t", test.expectedError, err)
			}
			if !reflect.DeepEqual(event, test.expectedEvent) {
				fmt.Println(event, test.expectedEvent)
				return fmt.Errorf("incorrect event returned")
			}
			return nil
		})
		if err != nil {
			t.Errorf("Error executing db update in test %d: %s", i, err)
		}
	}

	err = op.db.View(func(tx database.Tx) error {
		index, err := tx.GetRatingIndex()
		if err != nil {
			return err
		}
		if len(index) != 1 {
			return fmt.Errorf("expected index len 1 got %d", len(index))
		}
		if index[0].Slug != "abc" {
			return fmt.Errorf("expected slug abc got %s", index[0].Slug)
		}
		if index[0].Average != 5 {
			return fmt.Errorf("expected average 5 got %f", index[0].Average)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

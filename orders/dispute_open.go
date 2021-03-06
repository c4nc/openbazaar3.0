package orders

import (
	"github.com/cpacia/openbazaar3.0/database"
	"github.com/cpacia/openbazaar3.0/events"
	"github.com/cpacia/openbazaar3.0/models"
	npb "github.com/cpacia/openbazaar3.0/net/pb"
	"github.com/cpacia/openbazaar3.0/orders/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/libp2p/go-libp2p-core/peer"
)

func (op *OrderProcessor) processDisputeOpenMessage(dbtx database.Tx, order *models.Order, peer peer.ID, message *npb.OrderMessage) (interface{}, error) {
	disputeOpen := new(pb.DisputeOpen)
	if err := ptypes.UnmarshalAny(message.Message, disputeOpen); err != nil {
		return nil, err
	}
	dup, err := isDuplicate(disputeOpen, order.SerializedDisputeOpen)
	if err != nil {
		return nil, err
	}
	if order.SerializedDisputeOpen != nil && !dup {
		log.Errorf("Duplicate DISPUTE_OPEN message does not match original for order: %s", order.ID)
		return nil, ErrChangedMessage
	} else if dup {
		return nil, nil
	}

	if order.SerializedOrderComplete != nil {
		log.Errorf("Received DISPUTE_OPEN message for order %s after ORDER_COMPLETION", order.ID)
		return nil, ErrUnexpectedMessage
	}

	if order.SerializedPaymentFinalized != nil {
		log.Errorf("Received DISPUTE_OPEN message for order %s after PAYMENT_FINALIZED", order.ID)
		return nil, ErrUnexpectedMessage
	}

	if order.SerializedOrderReject != nil {
		log.Errorf("Received DISPUTE_OPEN message for order %s after ORDER_REJECT", order.ID)
		return nil, ErrUnexpectedMessage
	}

	if order.SerializedOrderCancel != nil {
		log.Errorf("Received DISPUTE_OPEN message for order %s after ORDER_CANCEL", order.ID)
		return nil, ErrUnexpectedMessage
	}

	orderOpen, err := order.OrderOpenMessage()
	if models.IsMessageNotExistError(err) {
		return nil, order.ParkMessage(message)
	}
	if err != nil {
		return nil, err
	}

	var (
		disputer       = orderOpen.BuyerID.PeerID
		disputerHandle = orderOpen.BuyerID.Handle
		disputee       = orderOpen.Listings[0].Listing.VendorID.PeerID
		disputeeHandle = orderOpen.Listings[0].Listing.VendorID.Handle
	)
	if disputeOpen.OpenedBy == pb.DisputeOpen_VENDOR {
		disputer = orderOpen.Listings[0].Listing.VendorID.PeerID
		disputerHandle = orderOpen.Listings[0].Listing.VendorID.Handle
		disputee = orderOpen.BuyerID.PeerID
		disputeeHandle = orderOpen.BuyerID.Handle
	}

	event := &events.DisputeOpen{
		OrderID: order.ID.String(),
		Thumbnail: events.Thumbnail{
			Tiny:  orderOpen.Listings[0].Listing.Item.Images[0].Tiny,
			Small: orderOpen.Listings[0].Listing.Item.Images[0].Small,
		},
		DisputerID:     disputer,
		DisputerHandle: disputerHandle,
		DisputeeID:     disputee,
		DisputeeHandle: disputeeHandle,
	}

	if (order.Role() == models.RoleBuyer && disputeOpen.OpenedBy == pb.DisputeOpen_BUYER) ||
		(order.Role() == models.RoleVendor && disputeOpen.OpenedBy == pb.DisputeOpen_VENDOR) {

		log.Infof("Processed own DISPUTE_OPEN for orderID: %s", order.ID)
	} else {
		log.Infof("Received DISPUTE_OPEN message for order %s", order.ID)
	}

	return event, order.PutMessage(message)
}

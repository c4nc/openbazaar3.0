package core

import (
	"github.com/cpacia/openbazaar3.0/models/factory"
	"github.com/cpacia/openbazaar3.0/orders/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/ipfs/go-cid"
	"strings"
	"testing"
	"time"
)

func TestOpenBazaarNode_SaveListing(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	defer node.DestroyNode()

	listing := factory.NewPhysicalListing("ron-swanson-shirt")

	done := make(chan struct{})
	if err := node.SaveListing(listing, done); err != nil {
		t.Fatal(err)
	}
	<-done

	_, err = node.GetMyListingBySlug("ron-swanson-shirt")
	if err != nil {
		t.Fatal(err)
	}

	index, err := node.GetMyListings()
	if err != nil {
		t.Fatal(err)
	}

	if len(index) != 1 {
		t.Errorf("Returned incorrect number of listings. Expected %d, got %d", 1, len(index))
	}

	c, err := cid.Decode(index[0].Hash)
	if err != nil {
		t.Fatal(err)
	}
	_, err = node.GetMyListingByCID(c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenBazaarNode_DeleteListing(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	defer node.DestroyNode()

	listing := factory.NewPhysicalListing("ron-swanson-shirt")

	done := make(chan struct{})
	if err := node.SaveListing(listing, done); err != nil {
		t.Fatal(err)
	}
	<-done

	done2 := make(chan struct{})
	if err := node.DeleteListing(listing.Slug, done2); err != nil {
		t.Fatal(err)
	}
	<-done2

	_, err = node.GetMyListingBySlug("ron-swanson-shirt")
	if err == nil {
		t.Fatal(err)
	}

	index, err := node.GetMyListings()
	if err != nil {
		t.Fatal(err)
	}

	if len(index) != 0 {
		t.Errorf("Returned incorrect number of listings. Expected %d, got %d", 0, len(index))
	}
}

func TestOpenBazaarNode_ListingsGet(t *testing.T) {
	network, err := NewMocknet(2)
	if err != nil {
		t.Fatal(err)
	}

	listing := factory.NewPhysicalListing("ron-swanson-shirt")

	done := make(chan struct{})
	if err := network.Nodes()[0].SaveListing(listing, done); err != nil {
		t.Fatal(err)
	}
	<-done

	listing2, err := network.Nodes()[1].GetListingBySlug(network.Nodes()[0].Identity(), listing.Slug, false)
	if err != nil {
		t.Fatal(err)
	}

	if listing2.Slug != listing.Slug {
		t.Errorf("Incorrect slug returned. Expected %s, got %s", listing.Slug, listing2.Slug)
	}

	index, err := network.Nodes()[1].GetListings(network.Nodes()[0].Identity(), false)
	if err != nil {
		t.Fatal(err)
	}

	if len(index) != 1 {
		t.Errorf("Returned incorrect number of listings in index. Expected %d, got %d", 1, len(index))
	}

	c, err := cid.Decode(index[0].Hash)
	if err != nil {
		t.Fatal(err)
	}
	listing2, err = network.Nodes()[1].GetListingByCID(c)
	if err != nil {
		t.Fatal(err)
	}

	if listing2.Slug != listing.Slug {
		t.Errorf("Incorrect slug returned. Expected %s, got %s", listing.Slug, listing2.Slug)
	}
}

func Test_generateListingSlug(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	defer node.DestroyNode()

	listing := factory.NewPhysicalListing("ron-swanson-shirt")

	done := make(chan struct{})
	if err := node.SaveListing(listing, done); err != nil {
		t.Fatal(err)
	}
	<-done

	tests := []struct {
		title    string
		expected string
	}{
		{
			"test",
			"test",
		},
		{
			"test title",
			"test-title",
		},
		{
			"ron swanson shirt",
			"ron-swanson-shirt1",
		},
		{
			"💩💩💩",
			"and-x1f4a9-and-x1f4a9-and-x1f4a9",
		},
		{
			strings.Repeat("s", 65),
			strings.Repeat("s", 65),
		},
		{
			strings.Repeat("s", 66),
			strings.Repeat("s", 65),
		},
	}

	for _, test := range tests {
		slug, err := node.generateListingSlug(test.title)
		if err != nil {
			t.Fatal(err)
		}
		if slug != test.expected {
			t.Errorf("Expected slug %s, got %s", test.expected, slug)
		}
	}
}

func Test_validateCryptocurrencyListing(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	defer node.DestroyNode()

	tests := []struct {
		listing   *pb.Listing
		transform func(listing *pb.Listing)
		valid     bool
	}{
		{
			listing:   factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {},
			valid:     true,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Coupons = []*pb.Listing_Coupon{
					{
						Title: "fads",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "fasdf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions = []*pb.Listing_ShippingOption{
					{
						Name: "fasdf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Item.Condition = "terrible"
			},
			valid: false,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency.Divisibility = 10
			},
			valid: false,
		},
	}

	for i, test := range tests {
		test.transform(test.listing)
		err := node.validateCryptocurrencyListing(test.listing)
		if test.valid && err != nil {
			t.Errorf("Test %d failed when it should not have: %s", i, err)
		} else if !test.valid && err == nil {
			t.Errorf("Test %d did not fail when it should have", i)
		}
	}
}

func Test_validateMarketPriceListing(t *testing.T) {
	tests := []struct {
		listing   *pb.Listing
		transform func(listing *pb.Listing)
		valid     bool
	}{
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.Format = pb.Listing_Metadata_MARKET_PRICE
				listing.Item.Price = ""
			},
			valid: true,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.Format = pb.Listing_Metadata_MARKET_PRICE
				listing.Item.Price = ""
				listing.Metadata.PriceModifier = -99.99
			},
			valid: true,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.Format = pb.Listing_Metadata_MARKET_PRICE
				listing.Item.Price = ""
				listing.Metadata.PriceModifier = 1000
			},
			valid: true,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.Format = pb.Listing_Metadata_MARKET_PRICE
				listing.Item.Price = ""
				listing.Metadata.PriceModifier = -100
			},
			valid: false,
		},
		{
			listing: factory.NewCryptoListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.Format = pb.Listing_Metadata_MARKET_PRICE
				listing.Item.Price = ""
				listing.Metadata.PriceModifier = 1001
			},
			valid: false,
		},
	}

	for i, test := range tests {
		test.transform(test.listing)
		err := validateMarketPriceListing(test.listing)
		if test.valid && err != nil {
			t.Errorf("Test %d failed when it should not have: %s", i, err)
		} else if !test.valid && err == nil {
			t.Errorf("Test %d did not fail when it should have", i)
		}
	}
}

func Test_validatePhysicalListing(t *testing.T) {
	tests := []struct {
		listing   *pb.Listing
		transform func(listing *pb.Listing)
		valid     bool
	}{
		{
			listing:   factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {},
			valid:     true,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency = nil
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency.Code = ""
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency.Name = ""
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency.CurrencyType = ""
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency.Code = strings.Repeat("s", WordMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency.Name = strings.Repeat("s", WordMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Metadata.PricingCurrency.CurrencyType = strings.Repeat("s", WordMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.Item.Condition = strings.Repeat("s", SentenceMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				for i := 0; i < MaxListItems+1; i++ {
					listing.Item.Options = append(listing.Item.Options, &pb.Listing_Item_Option{
						Name: "fadsfa",
					})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions = []*pb.Listing_ShippingOption{}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions = []*pb.Listing_ShippingOption{}
				for i := 0; i < MaxListItems+1; i++ {
					listing.ShippingOptions = append(listing.ShippingOptions, &pb.Listing_ShippingOption{
						Name: "fadsfa",
					})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = ""
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = strings.Repeat("s", WordMaxCharacters+1)
				listing.ShippingOptions[0].Regions = []pb.CountryCode{
					pb.CountryCode_UNITED_ARAB_EMIRATES,
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions = []*pb.Listing_ShippingOption{}
				for i := 0; i < 2; i++ {
					listing.ShippingOptions = append(listing.ShippingOptions, &pb.Listing_ShippingOption{
						Name: "fadsfa",
						Regions: []pb.CountryCode{
							pb.CountryCode_UNITED_ARAB_EMIRATES,
						},
					})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Regions = []pb.CountryCode{
					pb.CountryCode_UNITED_ARAB_EMIRATES,
				}
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE + 1
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Regions = []pb.CountryCode{}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Regions = []pb.CountryCode{
					0,
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Regions = []pb.CountryCode{
					501,
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Regions = []pb.CountryCode{}
				for i := 0; i < MaxCountryCodes+1; i++ {
					listing.ShippingOptions[0].Regions = append(listing.ShippingOptions[0].Regions, pb.CountryCode_AFGHANISTAN)
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{}
				for i := 0; i < MaxListItems+1; i++ {
					listing.ShippingOptions[0].Services = append(listing.ShippingOptions[0].Services, &pb.Listing_ShippingOption_Service{
						Name: "afdsf",
					})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{
					{
						Name: "",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{
					{
						Name: strings.Repeat("s", WordMaxCharacters+1),
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{
					{
						Name:              "asdf",
						EstimatedDelivery: "adf",
					},
					{
						Name: "asdf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{
					{
						Name: "asdf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{
					{
						Name:              "asdf",
						EstimatedDelivery: strings.Repeat("s", SentenceMaxCharacters+1),
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewPhysicalListing("test-listing"),
			transform: func(listing *pb.Listing) {
				listing.ShippingOptions[0].Name = "afsdf"
				listing.ShippingOptions[0].Type = pb.Listing_ShippingOption_FIXED_PRICE
				listing.ShippingOptions[0].Services = []*pb.Listing_ShippingOption_Service{
					{
						Name:              "asdf",
						EstimatedDelivery: "asdf",
						Price:             strings.Repeat("s", WordMaxCharacters+1),
					},
				}
			},
			valid: false,
		},
	}

	for i, test := range tests {
		test.transform(test.listing)
		err := validatePhysicalListing(test.listing)
		if test.valid && err != nil {
			t.Errorf("Test %d failed when it should not have: %s", i, err)
		} else if !test.valid && err == nil {
			t.Errorf("Test %d did not fail when it should have", i)
		}
	}
}

func Test_validateListing(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	defer node.DestroyNode()

	tests := []struct {
		listing   *pb.SignedListing
		transform func(sl *pb.SignedListing)
		valid     bool
	}{
		{
			listing:   factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {},
			valid:     true,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Slug = ""
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Slug = strings.Repeat("s", SentenceMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Slug = " "
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Slug = "/"
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata = nil
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata.ContractType = pb.Listing_Metadata_CRYPTOCURRENCY + 1
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata.Format = pb.Listing_Metadata_MARKET_PRICE + 1
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata.Expiry = nil
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				ts, _ := ptypes.TimestampProto(time.Time{})
				sl.Listing.Metadata.Expiry = ts
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata.Language = strings.Repeat("s", WordMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				node.testnet = false
				sl.Listing.Metadata.EscrowTimeoutHours = 1
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata.AcceptedCurrencies = []string{}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata.AcceptedCurrencies = []string{}
				for i := 0; i < MaxListItems+1; i++ {
					sl.Listing.Metadata.AcceptedCurrencies = append(sl.Listing.Metadata.AcceptedCurrencies, "abc")
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Metadata.AcceptedCurrencies = []string{
					strings.Repeat("s", WordMaxCharacters+1),
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Title = ""
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Price = "0"
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Title = strings.Repeat("s", TitleMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Description = strings.Repeat("s", DescriptionMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.ProcessingTime = strings.Repeat("s", SentenceMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Tags = []string{}
				for i := 0; i < MaxTags+1; i++ {
					sl.Listing.Item.Tags = append(sl.Listing.Item.Tags, "asdf")
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Tags = []string{
					strings.Repeat("s", WordMaxCharacters+1),
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Tags = []string{
					"",
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{}
				for i := 0; i < MaxListItems+1; i++ {
					sl.Listing.Item.Images = append(sl.Listing.Item.Images, &pb.Listing_Item_Image{})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:     "fasdf",
						Small:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Medium:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Large:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Original: "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:  "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Small: "adsf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:  "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Small: "adsf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Small:  "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Medium: "fasdf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Small:  "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Medium: "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Large:  "adfadf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:     "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Small:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Medium:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Large:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Original: "afdsf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:     "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Small:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Medium:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Large:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Original: "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Filename: "",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Images = []*pb.Listing_Item_Image{
					{
						Tiny:     "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Small:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Medium:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Large:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Original: "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
						Filename: strings.Repeat("s", FilenameMaxCharacters+1),
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Categories = []string{}
				for i := 0; i < MaxListItems+1; i++ {
					sl.Listing.Item.Categories = append(sl.Listing.Item.Categories, "asdf")
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Categories = []string{
					"",
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Categories = []string{
					strings.Repeat("s", WordMaxCharacters+1),
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "faddf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: strings.Repeat("s", WordMaxCharacters+1),
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "faddf",
							},
							{
								Name: "asdf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name:        "asdf",
						Description: strings.Repeat("s", SentenceMaxCharacters+1),
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "faddf",
							},
							{
								Name: "asdf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "faddf",
							},
							{
								Name: "asdf",
							},
						},
					},
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "faddf",
							},
							{
								Name: "asdf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
					},
				}
				for i := 0; i < MaxListItems+1; i++ {
					sl.Listing.Item.Options[0].Variants = append(sl.Listing.Item.Options[0].Variants, &pb.Listing_Item_Option_Variant{})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: strings.Repeat("s", WordMaxCharacters+1),
							},
							{
								Name: "asdf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
							},
							{
								Name: "asdf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
								Image: &pb.Listing_Item_Image{
									Tiny: "adf",
								},
							},
							{
								Name: "adf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
								Image: &pb.Listing_Item_Image{
									Tiny:  "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Small: "adfadf",
								},
							},
							{
								Name: "adf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
								Image: &pb.Listing_Item_Image{
									Tiny:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Small:  "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Medium: "adsf",
								},
							},
							{
								Name: "adf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
								Image: &pb.Listing_Item_Image{
									Tiny:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Small:  "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Medium: "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Large:  "asdf",
								},
							},
							{
								Name: "adf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
								Image: &pb.Listing_Item_Image{
									Tiny:     "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Small:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Medium:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Large:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Original: "asdf",
								},
							},
							{
								Name: "adf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
								Image: &pb.Listing_Item_Image{
									Tiny:     "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Small:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Medium:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Large:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Original: "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Filename: strings.Repeat("s", FilenameMaxCharacters+1),
								},
							},
							{
								Name: "adf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "asdf",
								Image: &pb.Listing_Item_Image{
									Tiny:     "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Small:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Medium:   "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Large:    "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Original: "QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
									Filename: "",
								},
							},
							{
								Name: "adf",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "fffff",
							},
							{
								Name: "asdf",
							},
						},
					},
				}
				sl.Listing.Item.Skus = []*pb.Listing_Item_Sku{}
				for i := 0; i < 3; i++ {
					sl.Listing.Item.Skus = append(sl.Listing.Item.Skus, &pb.Listing_Item_Sku{})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "fffff",
							},
							{
								Name: "asdf",
							},
						},
					},
				}
				sl.Listing.Item.Skus = []*pb.Listing_Item_Sku{
					{
						Selections: []*pb.Listing_Item_Sku_Selection{{}},
						ProductID:  strings.Repeat("s", WordMaxCharacters+1),
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "asdf",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "fffff",
							},
							{
								Name: "asdf",
							},
						},
					},
				}
				sl.Listing.Item.Skus = []*pb.Listing_Item_Sku{
					{
						ProductID: "adsf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "color",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "red",
							},
							{
								Name: "blue",
							},
						},
					},
				}
				sl.Listing.Item.Skus = []*pb.Listing_Item_Sku{
					{
						ProductID: "adsf",
						Selections: []*pb.Listing_Item_Sku_Selection{
							{
								Option:  "color",
								Variant: "red",
							},
						},
					},
					{
						ProductID: "adsf",
						Selections: []*pb.Listing_Item_Sku_Selection{
							{
								Option:  "color",
								Variant: "red",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "color",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "red",
							},
							{
								Name: "blue",
							},
						},
					},
				}
				sl.Listing.Item.Skus = []*pb.Listing_Item_Sku{
					{
						ProductID: "adsf",
						Selections: []*pb.Listing_Item_Sku_Selection{
							{
								Option:  "color",
								Variant: "red",
							},
							{
								Option:  "size",
								Variant: "red",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "color",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "red",
							},
							{
								Name: "blue",
							},
						},
					},
				}
				sl.Listing.Item.Skus = []*pb.Listing_Item_Sku{
					{
						ProductID: "adsf",
						Selections: []*pb.Listing_Item_Sku_Selection{
							{
								Option:  "color",
								Variant: "yellow",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Options = []*pb.Listing_Item_Option{
					{
						Name: "color",
						Variants: []*pb.Listing_Item_Option_Variant{
							{
								Name: "red",
							},
							{
								Name: "blue",
							},
						},
					},
				}
				sl.Listing.Item.Skus = []*pb.Listing_Item_Sku{
					{
						ProductID: "adsf",
						Selections: []*pb.Listing_Item_Sku_Selection{
							{
								Option:  "size",
								Variant: "red",
							},
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Item.Price = strings.Repeat("1", SentenceMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Taxes = []*pb.Listing_Tax{}
				for i := 0; i < MaxListItems+1; i++ {
					sl.Listing.Taxes = append(sl.Listing.Taxes, &pb.Listing_Tax{})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Taxes = []*pb.Listing_Tax{
					{
						TaxType: "",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Taxes = []*pb.Listing_Tax{
					{
						TaxType: strings.Repeat("s", WordMaxCharacters+1),
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Taxes = []*pb.Listing_Tax{
					{
						TaxType: "asdf",
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Taxes = []*pb.Listing_Tax{
					{
						TaxType: "asdf",
					},
				}
				for i := 0; i < MaxCountryCodes+1; i++ {
					sl.Listing.Taxes[0].TaxRegions = append(sl.Listing.Taxes[0].TaxRegions, pb.CountryCode_ALBANIA)
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Taxes = []*pb.Listing_Tax{
					{
						TaxType:    "asdf",
						TaxRegions: []pb.CountryCode{pb.CountryCode_ALBANIA},
						Percentage: 0,
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Taxes = []*pb.Listing_Tax{
					{
						TaxType:    "asdf",
						TaxRegions: []pb.CountryCode{pb.CountryCode_ALBANIA},
						Percentage: 101,
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Coupons = []*pb.Listing_Coupon{}
				for i := 0; i < MaxListItems+1; i++ {
					sl.Listing.Coupons = append(sl.Listing.Coupons, &pb.Listing_Coupon{})
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Coupons = []*pb.Listing_Coupon{
					{
						Title: strings.Repeat("s", CouponTitleMaxCharacters+1),
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Coupons = []*pb.Listing_Coupon{
					{
						Title: "asdf",
						Code: &pb.Listing_Coupon_DiscountCode{
							DiscountCode: strings.Repeat("s", CodeMaxCharacters+1),
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Coupons = []*pb.Listing_Coupon{
					{
						Title: "asdf",
						Discount: &pb.Listing_Coupon_PercentDiscount{
							PercentDiscount: 101,
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Coupons = []*pb.Listing_Coupon{
					{
						Title: "asdf",
						Discount: &pb.Listing_Coupon_PriceDiscount{
							PriceDiscount: strings.Repeat("1", SentenceMaxCharacters+1),
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Coupons = []*pb.Listing_Coupon{
					{
						Title: "asdf",
						Discount: &pb.Listing_Coupon_PriceDiscount{
							PriceDiscount: "0",
						},
					},
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Moderators = []string{}
				for i := 0; i < MaxListItems+1; i++ {
					sl.Listing.Moderators = append(sl.Listing.Moderators, " ")
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.Moderators = []string{
					"dafd",
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.TermsAndConditions = strings.Repeat("s", PolicyMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.RefundPolicy = strings.Repeat("s", PolicyMaxCharacters+1)
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID = nil
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID = &pb.ID{
					Handle: strings.Repeat("s", SentenceMaxCharacters+1),
				}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID = &pb.ID{}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID.PeerID = "adsf"
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID.Pubkeys.Secp256K1 = []byte("asdf")
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID.Pubkeys.Secp256K1 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID.Sig = []byte{0x00}
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Listing.VendorID.Sig[25] = 0x00
			},
			valid: false,
		},
		{
			listing: factory.NewSignedListing(),
			transform: func(sl *pb.SignedListing) {
				sl.Signature[25] = 0x00
			},
			valid: false,
		},
	}

	for i, test := range tests {
		test.transform(test.listing)
		err := node.validateListing(test.listing)
		if test.valid && err != nil {
			t.Errorf("Test %d failed when it should not have: %s", i, err)
		} else if !test.valid && err == nil {
			t.Errorf("Test %d did not fail when it should have", i)
		}
	}
}
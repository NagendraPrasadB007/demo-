package tests

/*
import (
	handler "pickupmanagement/handler"
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

// raiserequest
func TestRaiserequest(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/raisepickup/"

	payload := gin.H{
		"customerid":                "12345",
		"pickupdroptype":            "Express",
		"pickuplocation":            "123 Main Street",
		"droplocation":              "456 Oak Avenue",
		"pickupscheduleslot":        "Morning",
		"pickupscheduledate":        "2022-01-01T10:00:00Z",
		"actualpickupdatetime":      "2022-01-01T10:15:00Z",
		"pickupagentid":             789,
		"pickupfacilityid":          "ABC123",
		"pickupstatus":              "Unassigned",
		"paymentstatus":             "Pending",
		"pickupaddress":             "789 Elm Road",
		"domesticforeignidentifier": "domestic",
		"pickuplong":                "12.345",
		"pickuplat":                 "67.890",
		"pickuprequestedpincode":    "12345",
		"customername":              "Nagendra",
		"customermobilenumber":      "9353844206",
		"assigneddatetime":          "2023-11-24T13:00:00Z",
		"articleid":                 "ART123",
		"articlestate":              "Shipped",
		"articletype":               "Package",
		"articlecontent":            "Electronics",
		"articleimageid":            123,
		"articlepickupcharges":      5.99,
		"ispremailing":              true,
		"isparcelpacking":           false,
		"createddatetime":           "2022-01-01T08:00:00Z",
		"modifieddatetime":          "2022-01-01T09:30:00Z",
		"customerdacpickup":         "Yes",
		"addresstype":               "Residential",
		"bkgtransactionid":          "BKG123",
		"originpin":                 54321,
		"destinationpin":            98765,
		"physicalweight":            2.5,
		"shape":                     "Rectangle",
		"dimensionlength":           10.0,
		"dimensionbreadth":          5.0,
		"dimensionheight":           3.0,
		"volumetricweight":          3.0,
		"chargedweight":             2.8,
		"mailservicetypecode":       "Standard",
		"bkgtype":                   "Regular",
		"mailform":                  "Parcel",
		"isprepaid":                 true,
		"prepaymenttype":            "Credit Card",
		"valueofprepayment":         20.0,
		"vpcodflag":                 false,
		"valueforvpcod":             0.0,
		"insuranceflag":             true,
		"insurancetype":             "Basic",
		"valueinsurance":            15.0,
		"acknowledgementpod":        false,
		"instructionsrts":           "Handle with care",
		"addressrefsender":          "Sender123",
		"addressrefreceiver":        "Receiver456",
		"addressrefsenderaltaddr":   "AltAddrSender",
		"addressrefreceiveraktaddr": "AltAddrReceiver",
		"barcodenumber":             "BAR123",
		"pickupflag":                true,
		"basetariff":                10.0,
		"tax":                       2.0,
		"totaltariff":               12.0,
		"modeofpayment":             "Cash on Delivery",
		"paymenttranid":             "PAY123",
		"status":                    "Pending",
		"createdon":                 "2022-01-01T08:30:00Z",
		"createdby":                 "User123",
		"updatedon":                 "2022-01-01T09:45:00Z",
		"updatedby":                 "User456",
		"authorisedon":              "2022-01-01T10:00:00Z",
		"authorisedby":              "Admin",
		"facilityid":                "FAC789",
		"reqipaddress":              "192.168.1.1",
		"bookingchannel":            "Web",
		"contractnumber":            "CON123",
		"isparcel":                  true,
		"iscod":                     false,
		"totalamount":               150.5,
		"pickupcharges":             10.0,
		"registrationfee":           5.0,
		"postage":                   20.5,
		"ackorpodfee":               15.0,
		"doordeliverycharges":       12.0,
		"packingfee":                8.0,
		"cgst":                      2.5,
		"sgst":                      2.5,
		"othercharges":              5.0,
		"paymenttype":               "Online",
		"paymentdatetime":           "2023-12-15T14:30:00Z",
		"paidamount":                50.0,
	}

	var userResponse handler.RaisedPickupRequestResponseDom
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "post")

}
*/

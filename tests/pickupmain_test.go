package tests

import (
	handler "pickupmanagement/handler"
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

// createpickupmainrequest
func TestCreatepickupmainrequest(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/pickupmain/singlereq"

	payload := gin.H{
		"customerid":                "129",
		"pickupdroptype":            "Pickup",
		"pickuplocation":            "chennai",
		"droplocation":              "chennai",
		"pickupscheduleslot":        "10:00am to 1:00pm",
		"pickupscheduledate":        "2023-11-24T13:00:00Z",
		"actualpickupdatetime":      "2023-11-24T13:30:00Z",
		"pickupagentid":             87654321,
		"pickupfacilityid":          "HO123456789",
		"pickupstatus":              "Pickedup",
		"paymentstatus":             "Paid",
		"pickupaddress":             "Mysore",
		"domesticforeignidentifier": "domestic",
		"pickuplong":                "28.123",
		"pickuplat":                 "34.455",
		"pickuprequestedpincode":    "570010",
		"customername":              "Nagendra",
		"customermobilenumber":      "9353844206",
		"assigneddatetime":          "2023-11-24T13:00:00Z",
	}

	var userResponse handler.PickupmainResponse
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "post")

}

/*
//createbulkpickupmainrequest

func TestCreatebulkpickupmainrequest(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/pickupmain/bulk"

	payload := gin.H{
		"customerid":                "129",
		"pickupdroptype":            "Pickup",
		"pickuplocation":            "chennai",
		"droplocation":              "chennai",
		"pickupscheduleslot":        "10:00am to 1:00pm",
		"pickupscheduledate":        "2023-11-24T13:00:00Z",
		"actualpickupdatetime":      "2023-11-24T13:30:00Z",
		"pickupagentid":             87654321,
		"pickupfacilityid":          "HO123456789",
		"pickupstatus":              "Pickedup",
		"paymentstatus":             "Paid",
		"pickupaddress":             "Mysore",
		"domesticforeignidentifier": "domestic",
		"pickuplong":                "28.123",
		"pickuplat":                 "34.455",
		"pickuprequestedpincode":    "570010",
		"customername":              "Nagendra",
		"customermobilenumber":      "9353844206",
		"assigneddatetime":          "2023-11-24T13:00:00Z",
	}

	var userResponse []domain.Pickupmain
	// var userResponse handler.PickupmainResponse
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "post")

}

*/

/*
// AssigningProcess
func TestAssigningProcess(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/pickupmain/assigning"

	payload := gin.H
	[
		{
		"pickuprequestid": 90,
		"pickupagentid":   999,
		},
	]
	response := "Pickup requests assigned successfully"
	sendAndAssertPostRequest(t, test, url, payload, &response, "put")

}
*/

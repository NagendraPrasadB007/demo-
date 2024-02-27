package tests

import (
	handler "pickupmanagement/handler"
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

// getdetailsbypickuprequestid
func TestGetdetailsbypickuprequestid(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/getdetails/210"
	payload := gin.H{}
	var userResponse handler.RaisedPickupRequestResponseInt1
	// var userResponse handler.RaisedPickupRequestResponseDom
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

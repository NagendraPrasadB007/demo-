package tests

import (
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

// updatedetailsbypickuprequestid
func TestUpdatedetailsbypickuprequestid(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/updatedetails/9"

	payload := gin.H{
		"customermobilenumber": "8888888888",
	}
	response := "Pickuprequest details Updated successfully"
	sendAndAssertPostRequest(t, test, url, payload, &response, "put")

}

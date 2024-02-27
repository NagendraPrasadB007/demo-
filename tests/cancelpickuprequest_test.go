package tests

import (
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

//cancelpickuprequestbypickuprequestid

func TestCancelpickuprequestbypickuprequestid(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/cancel/61"
	payload := gin.H{}
	response := "Pickuprequest Cancelled successfully"
	sendAndAssertPostRequest(t, test, url, payload, &response, "put")
}

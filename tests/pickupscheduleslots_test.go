package tests

import (
	"pickupmanagement/core/domain"
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetallscheduleslots(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/pickupscheduleslots/"
	payload := gin.H{}
	var userResponse []domain.Pickupscheduleslots
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

// getallpickupagent
func TestGetallpickupagent(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/pickupagentlist/"
	payload := gin.H{}
	var userResponse []domain.Pickupagent
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

// getallremarks
func TestGetallremarks(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/remarkslist/"
	payload := gin.H{}
	var userResponse []domain.Remarks
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

// getalladdresstypelist
func TestGetalladdresstypelist(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/addresstypelist/"
	payload := gin.H{}
	var userResponse []domain.Addresstype
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

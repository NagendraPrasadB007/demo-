package tests

import (
	handler "pickupmanagement/handler"
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

// count
func TestCountdetails(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/dashboard/count/FacilityA"
	payload := gin.H{}
	var userResponse handler.CountDetails
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

// Assignedlist
func TestAssignedlist(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/dashboard/Assigned/HO123456789"
	payload := gin.H{}
	var userResponse []int
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

/*
// Unassignedlist
func TestUnassignedlist(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/dashboard/Unassigned/HO123456789"
	payload := gin.H{}
	var userResponse handler.UnassignedDom
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}
*/

/*
// Pickuprequestlist
func TestPickuprequestlist(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/dashboard/pickuprequest/BS123"
	payload := gin.H{}
	var userResponse handler.PickuprequestBasicfac
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}
*/

// Pickuprequestlistbypincode
func TestPickuprequestlistbypincode(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/dashboard/pickuprequests/570010"
	payload := gin.H{}
	var userResponse []int
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

// Assignedrequesttoagent
func TestAssignedrequesttoagent(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/dashboard/getpickuprequest/1000"
	payload := gin.H{}
	var userResponse []int
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

/*
// Customerrequests
func TestCustomerrequests(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/dashboard/customerpickuprequests/12345"
	payload := gin.H{}
	var userResponse handler.PickuprequestBasiccus
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}
*/

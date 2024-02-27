package tests

import (
	handler "pickupmanagement/handler"
	"pickupmanagement/supertest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetaddressdetailsbyid(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/addressdetails/666"
	payload := gin.H{}
	var userResponse handler.AddressdetailsResponseTest
	// var userResponse handler.AddressdetailsResponse
	// var userResponse []domain.Addressdetails
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "get")
}

func TestCreateaddressforcustomer(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/addressdetails/"

	payload := gin.H{
		"customerid":   "91",
		"firstname":    "Tarun",
		"lastname":     "Pant",
		"addressline1": "door no 5",
		"addressline2": "B block 1st cross",
		"landmark":     "nanjungud",
		"city":         "Mysuru",
		"state":        "karnataka",
		"country":      "India",
		"pincode":      "570015",
		"mobilenumber": "9483700509",
		"emailid":      "puni@gmail.com",
		"geocode":      "85.126 200.123",
		"addresstype":  "lab",
		"fromtopickup": "yes",
		"isverified":   true,
	}

	var userResponse handler.AddressdetailsResponse
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "post")

}

/*
// not working
func TestUpdateadress(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/addressdetails/6"

	payload := gin.H{
		"customerid":   "999",
		"firstname":    "Rishab",
		"lastname":     "Pant",
		"addressline1": "door no 5",
		"addressline2": "B block 1st cross",
		"landmark":     "nanjungud",
		"city":         "Mysuru",
		"state":        "karnataka",
		"country":      "India",
		"pincode":      "570015",
		"mobilenumber": "9483700509",
		"emailid":      "puni@gmail.com",
		"geocode":      "85.126 200.123",
		"addresstype":  "lab",
		"fromtopickup": "yes",
		"isverified":   true,
	}
	var userResponse handler.AddressdetailsResponse
	sendAndAssertPostRequest(t, test, url, payload, &userResponse, "put")

}
*/

// func TestDeleteadress(t *testing.T) {
// 	test := supertest.NewSuperTest(router, t)
// 	url := "/pickup/v1/addressdetails/40"
// 	payload := gin.H{}

// 	var responseMessage map[string]string
// 	sendAndAssertPostRequest(t, test, url, payload, &responseMessage, "delete")

// 	// Assert the response message
// 	expectedMessage := "AddressID 40 deleted successfully"
// 	assert.Equal(t, expectedMessage, responseMessage["message"])
// }

func TestDeleteadress(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/addressdetails/54"
	payload := gin.H{}
	response := "Address deleted succesfully"
	sendAndAssertPostRequest(t, test, url, payload, &response, "delete")
}

/*
func TestUserInvalidError(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/v1/users/invalid/"
	payload := gin.H{}
	sendAndAssertErrorRequest(t, test, url, payload, "", "get", http.StatusNotFound)

}
*/

/*
func TestUserdatanotfoundError(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	//url := "/pickup/v1/addressdetails/3200"
	url := "/v1/users/123344/"
	payload := gin.H{}

	sendAndAssertErrorRequest(t, test, url, payload, "", "get", http.StatusNotFound)
}
*/

/*
// 415 expected 400 actual
func TestUserUnsupportedError(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	//url := "/v1/users/123344/"
	url := "/pickup/v1/addressdetails/32"
	payload := gin.H{}

	sendAndAssertErrorRequest(t, test, url, payload, "", "unsupported", http.StatusUnsupportedMediaType)

}
*/

/*
func TestUserUnprocessedError(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/addressdetails/"
	payload := gin.H{
		"customerid":   "91",
		"firstname":    "Tarun",
		"lastname":     "Pant",
		"addressline1": "door no 5",
		"addressline2": "B block 1st cross",
		"landmark":     "nanjungud",
		"city":         "Mysuru",
		"state":        "karnataka",
		"country":      "India",
		"pincode":      "570015",
		"mobilenumber": "9483700509",
		"emailid":      "puni@gmail.com",
		"geocode":      "85.126 200.123",
		"addresstype":  "lab",
		"fromtopickup": "yes",
		"isverified":   "true",
	}

	sendAndAssertErrorRequest(t, test, url, payload, "", "post", http.StatusUnprocessableEntity)
}
*/

/*
func TestUserdbError(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/pickup/v1/addressdetails/"
	payload := gin.H{
		"customerid":   "91",
		"firstname":    "Tarun",
		"lastname":     "Pant",
		"addressline1": "door no 5",
		"addressline2": "B block 1st cross",
		"landmark":     "nanjungud",
		"city":         "Mysuru",
		"state":        "karnataka",
		"country":      "India",
		"pincode":      "570015",
		"mobilenumber": "9483700509",
		"emailid":      "puni@gmail.com",
		"geocode":      "85.126 200.123",
		"addresstype":  "lab",
		"fromtopickup": "yes",
		"isverified":   true,
		//"check":        10, //extra added
	}

	sendAndAssertErrorRequest(t, test, url, payload, "", "post", http.StatusInternalServerError)

}
*/

/*

func TestGetUsers(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/v1/users/?skip=1&limit=5"
	payload := gin.H{}
	type GetAll struct {
		meta  `json:"meta"`
		Users []domain.User `json:"users"`
	}
	var userall GetAll
	sendAndAssertPostRequest(t, test, url, payload, &userall, "get")
}

func TestUserDelete(t *testing.T) {

	test := supertest.NewSuperTest(router, t)
	url := "/v1/users/373"
	payload := gin.H{}

	sendAndAssertPostRequest(t, test, url, payload, "delete", "delete")

}


func TestUserdbError(t *testing.T) {
	test := supertest.NewSuperTest(router, t)
	url := "/v1/users/"
	payload := gin.H{"email": "3248r11test288990@gmail.com",
		"password": "fghjklhjgf",
		"name":     "sawerr",
		"check":    10}

	sendAndAssertErrorRequest(t, test, url, payload, "", "post", http.StatusInternalServerError)

}
*/

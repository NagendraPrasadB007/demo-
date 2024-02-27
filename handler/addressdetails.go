package handler

import (
	"pickupmanagement/core/domain"
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AddressdetailsHandler struct {
	svc repo.AddressdetailsRepository
	log *logger.Logger
}

func NewAddressdetailsHandler(svc repo.AddressdetailsRepository, log *logger.Logger) *AddressdetailsHandler {
	return &AddressdetailsHandler{
		svc,
		log,
	}
}

type getaddressdetailsbyid struct {
	Customerid string `uri:"customerid" validate:"required,min=1"`
}

// GetAddress godoc
//
//	@Summary		Get a address
//	@Description	get a address by customerid
//	@Tags			Addressdetails
//	@Accept			json
//	@Produce		json
//	@Param			customerid	path		string				true	"Customerid ID"
//	@Success		200	{object}	AddressdetailsResponse	"Category retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/addressdetails/{customerid} [get]
func (ad *AddressdetailsHandler) getaddressdetailsbyid(ctx *gin.Context) {

	var add getaddressdetailsbyid
	// below line not needed when you need to fetch single record
	//var addressdetailsList []addressdetailsResponse

	if err := ctx.ShouldBindUri(&add); err != nil {
		ad.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	//addressdetails, err := ad.svc.GetAddressdetailsByID(ctx, add.Customerid)
	//below line not needed when you need to fetch single record // above line add
	addressdetailss, err := ad.svc.GetAddressdetailsByID(ctx, add.Customerid)
	if err != nil {
		ad.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	//for loop for fetching ,multiple record not needed for single
	//not needed to append
	// for _, addressdetails := range addressdetailss {
	// 	addressdetailsList = append(addressdetailsList, newAddressdetailsResponse(&addressdetails))
	// }

	//rsp := addressdetailsList

	rsp := addressdetailss

	//below line for fetching single record
	//rsp := newAddressdetailsResponse(addressdetails)

	handleSuccess(ctx, rsp)

}

// added insert tag

type createAddress struct {
	Customerid   string `json:"customerid" validate:"required" insert:"customer_id"`
	Firstname    string `json:"firstname" validate:"required" insert:"firstname"`
	Lastname     string `json:"lastname" validate:"required" insert:"lastname"`
	Addressline1 string `json:"addressline1" validate:"required"  insert:"addressline1"`
	Addressline2 string `json:"addressline2" validate:"required"  insert:"addressline2"`
	Landmark     string `json:"landmark" validate:"required"  insert:"landmark"`
	City         string `json:"city" validate:"required"  insert:"city"`
	State        string `json:"state" validate:"required"  insert:"state"`
	Country      string `json:"country" validate:"required"  insert:"country"`
	Pincode      string `json:"pincode" validate:"required"  insert:"pincode"`
	Mobilenumber string `json:"mobilenumber" validate:"required"  insert:"mobilenumber"`
	Emailid      string `json:"emailid" validate:"required" insert:"email_id"`
	Geocode      string `json:"geocode" validate:"required"  insert:"geo_code"`
	Addresstype  string `json:"addresstype" validate:"required"  insert:"address_type"`
	Fromtopickup string `json:"fromtopickup" validate:"required"  insert:"fromtopickup"`
	Isverified   bool   `json:"isverified" validate:"required" insert:"is_verified"`
}

// CreateAddress godoc
//
//	@Summary        Create a new address for a customer
//	@Description    create a new address  for a customer
//	@Tags           Addressdetails
//	@Accept         json
//	@Produce        json
//	@Param          createAddress   body        createAddress   true    "Create address request"
//	@Success        200                     {object}    AddressdetailsResponse        "address created"
//	@Failure        400                     {object}    errorValidResponse          "Validation error"
//	@Failure        401                     {object}    errorValidResponse          "Unauthorized error"
//	@Failure        403                     {object}    errorValidResponse          "Forbidden error"
//	@Failure        404                     {object}    errorValidResponse          "Data not found error"
//	@Failure        409                     {object}    errorValidResponse          "Data conflict error"
//	@Failure        500                     {object}    errorValidResponse          "Internal server error"
//	@Router         /addressdetails/ [post]
func (ad *AddressdetailsHandler) createaddressforcustomer(ctx *gin.Context) {
	var post createAddress
	if err := ctx.ShouldBindJSON(&post); err != nil {
		handleError(ctx, err)
		ad.log.Debug("error occured during binding:", err)
		return
	}
	if !handleValidation(ctx, post) {
		return
	}

	//var address domain.Addressdetails

	address := domain.Addressdetails{
		Customerid:   post.Customerid,
		Firstname:    post.Firstname,
		Lastname:     post.Lastname,
		Addressline1: post.Addressline1,
		Addressline2: post.Addressline2,
		Landmark:     post.Landmark,
		City:         post.City,
		State:        post.State,
		Country:      post.Country,
		Pincode:      post.Pincode,
		Mobilenumber: post.Mobilenumber,
		Emailid:      post.Emailid,
		Geocode:      post.Geocode,
		Addresstype:  post.Addresstype,
		Fromtopickup: post.Fromtopickup,
		Isverified:   post.Isverified,
	}

	// //above mapping commented below copier code added
	// err := copier.Copy(&address, &post)
	// if err != nil {
	// 	ad.log.Debug("error: ", err)
	// }
	// ad.log.Debug("customerid:", address.Customerid)

	//ends here

	addresscreated, err := ad.svc.CreateAddressForCustomer(ctx, &address)
	if err != nil {
		ad.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	ad.log.Debug("addresscreated", addresscreated)
	rsp := newCreatedAddressdetailsResponse(addresscreated)

	handleSuccess(ctx, rsp)

}

// type updateAddressDetails struct {
// 	Customerid   string `json:"customerid" validate:"omitempty,required,min=1" update:"customer_id"`
// 	Firstname    string `json:"firstname" validate:"omitempty,required,min=1" update:"firstname"`
// 	Lastname     string `json:"lastname" validate:"omitempty,required,min=1" update:"lastname"`
// 	Addressline1 string `json:"addressline1" validate:"omitempty,required,min=1" update:"addressline1"`
// 	Addressline2 string `json:"addressline2" validate:"omitempty,required,min=1" update:"addressline2"`
// 	Landmark     string `json:"landmark" validate:"omitempty,required,min=1" update:"landmark"`
// 	City         string `json:"city" validate:"omitempty,required,min=1" update:"city"`
// 	State        string `json:"state" validate:"omitempty,required,min=1" update:"state"`
// 	Country      string `json:"country" validate:"omitempty,required,min=1" update:"country"`
// 	Pincode      string `json:"pincode" validate:"omitempty,required,min=1" update:"pincode"`
// 	Mobilenumber string `json:"mobilenumber" validate:"omitempty,required,min=1" update:"mobilenumber"`
// 	Emailid      string `json:"emailid" validate:"omitempty,required,min=1" update:"email_id"`
// 	Geocode      string `json:"geocode" validate:"omitempty,required,min=1" update:"geo_code"`
// 	Addresstype  string `json:"addresstype" validate:"omitempty,required,min=1" update:"address_type"`
// 	Fromtopickup string `json:"fromtopickup" validate:"omitempty,required,min=1" update:"fromtopickup"`
// 	Isverified   bool   `json:"isverified" validate:"omitempty,required" update:"is_verified"`
// }

type updateAddressDetails struct {
	Customerid   string `json:"customerid" validate:"omitempty,required,min=1"`
	Firstname    string `json:"firstname" validate:"omitempty,required,min=1"`
	Lastname     string `json:"lastname" validate:"omitempty,required,min=1"`
	Addressline1 string `json:"addressline1" validate:"omitempty,required,min=1"`
	Addressline2 string `json:"addressline2" validate:"omitempty,required,min=1"`
	Landmark     string `json:"landmark" validate:"omitempty,required,min=1"`
	City         string `json:"city" validate:"omitempty,required,min=1"`
	State        string `json:"state" validate:"omitempty,required,min=1"`
	Country      string `json:"country" validate:"omitempty,required,min=1"`
	Pincode      string `json:"pincode" validate:"omitempty,required,min=1"`
	Mobilenumber string `json:"mobilenumber" validate:"omitempty,required,min=1"`
	Emailid      string `json:"emailid" validate:"omitempty,required,min=1"`
	Geocode      string `json:"geocode" validate:"omitempty,required,min=1"`
	Addresstype  string `json:"addresstype" validate:"omitempty,required,min=1"`
	Fromtopickup string `json:"fromtopickup" validate:"omitempty,required,min=1"`
	Isverified   bool   `json:"isverified" validate:"omitempty,required"`
}

// UpdateCategory godoc
//
//	@Summary		Update a address
//	@Description	update a address's based on addressid
//	@Tags			Addressdetails
//	@Accept			json
//	@Produce		json
//	@Param			addressid				path		int					true	"update address"
//	@Param			updateAddressDetails	body		updateAddressDetails	true	"updateAddressDetails"
//	@Success		200						{object}	AddressdetailsResponse		"Category updated"
//	@Failure		400						{object}	errorValidResponse			"Validation error"
//	@Failure		401						{object}	errorValidResponse			"Unauthorized error"
//	@Failure		403						{object}	errorValidResponse			"Forbidden error"
//	@Failure		404						{object}	errorValidResponse			"Data not found error"
//	@Failure		409						{object}	errorValidResponse			"Data conflict error"
//	@Failure		500						{object}	errorValidResponse			"Internal server error"
//	@Router			/addressdetails/{addressid} [put]
func (ad *AddressdetailsHandler) updateadress(ctx *gin.Context) {
	var req updateAddressDetails
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, err)
		ad.log.Debug("error occured during binding:", err)
		return
	}
	if !handleValidation(ctx, req) {
		return
	}

	idStr := ctx.Param("addressid")
	// here idStr is of type string i am converting string into int because adressid in my DB is int
	addressID, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(ctx, err)
		ad.log.Debug("error occured while converting addressid:", err)
		return
	}

	addressdetails, err := ad.svc.GetAddressdetailsByaddressid(ctx, addressID)
	if err != nil {
		ad.log.Debug("error occured while converting addressid:", err)
		return
	}

	ad.log.Debug("address from repo ", addressdetails)

	var address domain.Addressdetails

	err = copier.Copy(&address, &addressdetails)
	if err != nil {
		ad.log.Debug("error occured while converting addressid:", err)
		return
	}

	ad.log.Debug("city req req ", req.City)
	ad.log.Debug("city req ad", address.City)

	err = copier.Copy(&address, &req)
	if err != nil {
		ad.log.Debug("error occured while converting addressid:", err)
		return
	}

	// address := domain.Addressdetails{
	// 	Customerid:   req.Customerid,
	// 	Addressid:    addressID,
	// 	Firstname:    req.Firstname,
	// 	Lastname:     req.Lastname,
	// 	Addressline1: req.Addressline1,
	// 	Addressline2: req.Addressline2,
	// 	Landmark:     req.Landmark,
	// 	City:         req.City,
	// 	State:        req.State,
	// 	Country:      req.Country,
	// 	Pincode:      req.Pincode,
	// 	Mobilenumber: req.Mobilenumber,
	// 	Emailid:      req.Emailid,
	// 	Geocode:      req.Geocode,
	// 	Addresstype:  req.Addresstype,
	// 	Fromtopickup: req.Fromtopickup,
	// 	Isverified:   req.Isverified,
	// }

	// ad.log.Debug("after copying ", addressdetails)
	ad.log.Debug("after copying ", address)

	_, err = ad.svc.UpdateAddress(ctx, &address)
	if err != nil {
		handledbError(ctx, err)
		ad.log.Debug("error occured while repo calling:", err)
		return
	}

	rsp := newAddressdetailsResponse(&address)

	handleSuccess(ctx, rsp)

}

type deleteaddressbyaddressid struct {
	Addressid int `uri:"addressid"`
}

// DeleteAddress godoc
//
//	@Summary		Delete a address
//	@Description	Delete a address based on addressid
//	@Tags			Addressdetails
//	@Accept			json
//	@Produce		json
//	@Param			addressid	path		int			true	"addressid"
//
// @Success 200 {string} string "Address deleted succesfully"
//
//	@Failure		400	{object}	errorValidResponse	"Validation error"
//	@Failure		401	{object}	errorValidResponse	"Unauthorized error"
//	@Failure		403	{object}	errorValidResponse	"Forbidden error"
//	@Failure		404	{object}	errorValidResponse	"Data not found error"
//	@Failure		500	{object}	errorValidResponse	"Internal server error"
//	@Router			/addressdetails/{addressid} [delete]
func (ad *AddressdetailsHandler) deleteadress(ctx *gin.Context) {
	var del deleteaddressbyaddressid

	if err := ctx.ShouldBindUri(&del); err != nil {
		ad.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, del) {
		return
	}

	err := ad.svc.DeleteAddressByAddressID(ctx, del.Addressid)
	if err != nil {
		ad.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "AddressID " + strconv.Itoa(del.Addressid) + " deleted successfully",
	// })

	handleSuccess(ctx, "Address deleted succesfully")

}

/*
//RR sir template

type updateAddressDetails struct {
	Customerid   string `json:"customerid" validate:"omitempty,required" update:"customer_id"`
	Firstname    string `json:"firstname" validate:"omitempty,required" update:"firstname"`
	Lastname     string `json:"lastname" validate:"omitempty,required" update:"lastname"`
	Addressline1 string `json:"addressline1" validate:"omitempty,required"  update:"addressline1"`
	Addressline2 string `json:"addressline2" validate:"omitempty,required"  update:"addressline2"`
	Landmark     string `json:"landmark" validate:"omitempty,required"  update:"landmark"`
	City         string `json:"city" validate:"omitempty,required"  update:"city"`
	State        string `json:"state" validate:"omitempty,required"  update:"state"`
	Country      string `json:"country" validate:"omitempty,required"  update:"country"`
	Pincode      string `json:"pincode" validate:"omitempty,required"  update:"pincode"`
	Mobilenumber string `json:"mobilenumber" validate:"omitempty,required"  update:"mobilenumber"`
	Emailid      string `json:"emailid" validate:"omitempty,required" update:"email_id"`
	Geocode      string `json:"geocode" validate:"omitempty,required"  update:"geo_code"`
	Addresstype  string `json:"addresstype" validate:"omitempty,required"  update:"address_type"`
	Fromtopickup string `json:"fromtopickup" validate:"omitempty,required"  update:"fromtopickup"`
	Isverified   bool   `json:"isverified" validate:"omitempty,required" update:"is_verified"`
}

func (ad *AddressdetailsHandler) updateadress(ctx *gin.Context) {
	var req updateAddressDetails
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, err)
		ad.log.Debug("error occured during binding:", err)
		return
	}
	if !handleValidation(ctx, req) {
		return
	}

	idStr := ctx.Param("addressid")
	addressID, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(ctx, err)
		ad.log.Debug("error occured while converting addressid:", err)
		return
	}
	ad.log.Debug("addressId :", addressID)

	var address domain.Addressdetails
	err = copier.Copy(&address, &req)
	if err != nil {
		ad.log.Debug("error: ", err)
	}
	ad.log.Debug("customerid:", address.Customerid)

	_, err = ad.svc.UpdateAddress(ctx, &address)
	if err != nil {
		handledbError(ctx, err)
		ad.log.Debug("error occured while repo calling:", err)
		return
	}

	rsp := newAddressdetailsResponse(&address)

	handleSuccess(ctx, rsp)

}
*/

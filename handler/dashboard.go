package handler

import (
	"net/http"
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	svc repo.DashboardRepository
	log *logger.Logger
}

func NewDashboardHandler(svc repo.DashboardRepository, log *logger.Logger) *DashboardHandler {
	return &DashboardHandler{
		svc,
		log,
	}
}

//count API starts here

type CountApiResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    CountDetails `json:"data"`
}

type CountDetails struct {
	UnassignedCount int `json:"unassigned_count"`
	AssignedCount   int `json:"assigned_count"`
	PickedupCount   int `json:"pickedup_count"`
	CancelledCount  int `json:"cancelled_count"`
	Sum             int `json:"sum"`
}

type getcountdetailsbyid struct {
	Pickupfacilityid string `uri:"facilityid" validate:"required,min=1"`
}

// Handler function to get the count of Unassigned , Assigned , pickedup, cancelled , total sum based on facilityid

// Getcount godoc
//
//	@Summary		Get the count of unassigned , assigned , pickedup , cancelled , total sum based on facilityid
//	@Description	Get the count of unassigned , assigned , pickedup , cancelled , total sum based on facilityid
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			facilityid	path		string				true	"Pickupfacilityid"
//	@Success		200	{object}	CountApiResponse	"count retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/dashboard/count/{facilityid} [get]
func (dh *DashboardHandler) countdetails(ctx *gin.Context) {

	var add getcountdetailsbyid

	if err := ctx.ShouldBindUri(&add); err != nil {
		dh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	counts, err := dh.svc.GetCountByID(ctx, add.Pickupfacilityid)
	if err != nil {
		dh.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	dh.log.Debug(counts)

	totalSum := counts.UnassignedCount + counts.AssignedCount + counts.PickedupCount + counts.CancelledCount

	response := CountApiResponse{
		Success: true,
		Message: "Success",
		Data: CountDetails{
			UnassignedCount: counts.UnassignedCount,
			AssignedCount:   counts.AssignedCount,
			PickedupCount:   counts.PickedupCount,
			CancelledCount:  counts.CancelledCount,
			Sum:             totalSum,
		},
	}

	ctx.JSON(http.StatusOK, response)

}

// Get the Assigned Pickuprequest based on facilityid
type getassigneddetailsbyid struct {
	Pickupfacilityid string `uri:"facilityid" validate:"required,min=1"`
}

// Handler function to get all Assigned pickuprequest based on facilityid

// GetAssignedlist godoc
//
//	@Summary		Get the Assigned pickuprequest based on facilityid
//	@Description	Get the Assigned pickuprequest based on facilityid
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			facilityid	path		string				true	"Pickupfacilityid"
//
// @Success 200 {object} []int
//
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/dashboard/Assigned/{facilityid} [get]
func (dh *DashboardHandler) Assignedlist(ctx *gin.Context) {

	var add getassigneddetailsbyid

	if err := ctx.ShouldBindUri(&add); err != nil {
		dh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	pickupRequestIDs, err := dh.svc.GetAssignedlistByID(ctx, add.Pickupfacilityid)
	if err != nil {
		dh.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	//rsp := newResponse(gin.H{"AssignedList": pickupRequestIDs})
	//handleSuccess(rsp)

	handleSuccess(ctx, pickupRequestIDs)
	//ctx.JSON(http.StatusOK, gin.H{"assignedList": pickupRequestIDs})

}

// Get the Unassigned Pickuprequest based on facilityid
type getunassigneddetailsbyid struct {
	Pickupfacilityid string `uri:"facilityid" validate:"required,min=1"`
}

// Handler function to get all Unassigned pickuprequest based on facilityid

// GetUnassignedlist godoc
//
//	@Summary		Get the Unassigned pickuprequest based on facilityid
//	@Description	Get the Unassigned pickuprequest based on facilityid
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			facilityid	path		string				true	"Pickupfacilityid"
//
// @Success		200	{object}	UnassignedDom	"Unassigned retrieved"
//
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/dashboard/Unassigned/{facilityid} [get]
func (dh *DashboardHandler) Unassignedlist(ctx *gin.Context) {

	var add getunassigneddetailsbyid

	if err := ctx.ShouldBindUri(&add); err != nil {
		dh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	dh.log.Debug("after validation ")

	//calling repo function to get all unassigned pickuprequestid
	pickupDetailsList, err := dh.svc.GetUnAssignedlistByID(ctx, add.Pickupfacilityid)
	if err != nil {
		dh.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	dh.log.Debug("after getting pickuprequest ids")

	// Create a slice to store the responses
	var responses []interface{}
	for _, pickupRequestID := range pickupDetailsList {
		// Call the Identify function to determine if the pickup request is domestic or international
		pickupMain, err := dh.svc.Identify(ctx, pickupRequestID)
		if err != nil {
			dh.log.Debug("error occured while repo calling:", err)
		}

		dh.log.Debug("after identify")
		dh.log.Debug("domestic or international :", pickupMain.Domesticforeignidentifier)

		if pickupMain.Domesticforeignidentifier == "domestic" {
			dh.log.Debug("entered domestic loop")
			pickupmain, domestic, err := dh.svc.GetcompletedetailsDom(ctx, pickupRequestID)
			if err != nil {
				dh.log.Debug("error occured while repo calling domestic:", err)
				handledbError(ctx, err)
				return
			}
			dh.log.Debug("after dom details ")
			// rsp := newFetchPickupRequestResponseDom(pickupmain, domestic)
			// handleSuccess(ctx, rsp)
			responses = append(responses, newFetchPickupRequestResponseDom(pickupmain, domestic))

		} else {
			dh.log.Debug("entered international loop")
			pickupmain, international, err := dh.svc.GetcompletedetailsInt(ctx, pickupRequestID)
			if err != nil {
				dh.log.Debug("error occured while repo calling international :", err)
				handledbError(ctx, err)
				return
			}
			dh.log.Debug("after international")

			// rsp := newFetchPickupRequestResponseInt(pickupmain, international)
			// handleSuccess(ctx, rsp)
			responses = append(responses, newFetchPickupRequestResponseInt(pickupmain, international))
		}
	}
	handleSuccess(ctx, responses)
}

// This API is used to get pickuprequestid based on facilityid
type getpickuprequestbyfacilityid struct {
	Pickupfacilityid string `uri:"facilityid" validate:"required,min=1"`
}

// Handler Function to get all pickuprequest based on facilityid

// GetPickuprequestlist godoc
//
// 	@Summary		Get the  pickuprequest list based on facilityid
// 	@Description	Get the  pickuprequest list based on facilityid
// 	@Tags			Dashboard
// 	@Accept			json
// 	@Produce		json
// 	@Param			facilityid	path		string				true	"Pickupfacilityid"

// @Success		200	{object}	PickuprequestBasicfac	"All pickuprequest retrieved"

// @Failure		400	{object}	errorValidResponse		"Validation error"
// @Failure		404	{object}	errorValidResponse		"Data not found error"
// @Failure		500	{object}	errorValidResponse		"Internal server error"
// @Router			/dashboard/pickuprequest/{facilityid} [get]
func (dh *DashboardHandler) Pickuprequestlist(ctx *gin.Context) {

	var add getpickuprequestbyfacilityid

	if err := ctx.ShouldBindUri(&add); err != nil {
		dh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	// calling repo function to get list of pickuprequestid based on facilityid
	pickupDetailsList, err := dh.svc.GetPickuprequestlistByFacilityid(ctx, add.Pickupfacilityid)
	if err != nil {
		dh.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	var responses []interface{}

	for _, pickupRequestID := range pickupDetailsList {
		pickupmain, err := dh.svc.Getcompletedetailsbasedobidfac(ctx, pickupRequestID)
		if err != nil {
			dh.log.Debug("error occured while repo calling:", err)
			handledbError(ctx, err)
			return
		}
		responses = append(responses, newPickuprequestBasicfac(pickupmain))
	}

	handleSuccess(ctx, responses)

}

// This API is used to get pickuprequestid based on pincode
type getpickuprequestbypincode struct {
	Pickuprequestedpincode string `uri:"pincode" validate:"required,min=1"`
}

// Handler Function to get all pickuprequest based on pincode

// GetPickuplistbypincode godoc
//
//	@Summary		Get the pickuprequest based on pincode
//	@Description	Get the pickuprequest based on pincode
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			pincode	path		string				true	"Pickuprequestedpincode"
//
// @Success 200 {object} []int
//
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/dashboard/pickuprequests/{pincode} [get]
func (dh *DashboardHandler) Pickuprequestlistbypincode(ctx *gin.Context) {

	var add getpickuprequestbypincode

	if err := ctx.ShouldBindUri(&add); err != nil {
		dh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	pickupRequestIDs, err := dh.svc.GetPickuprequestlistBypincode(ctx, add.Pickuprequestedpincode)
	if err != nil {
		dh.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	handleSuccess(ctx, pickupRequestIDs)

}

// This api gives all pickuprequestid assigned to pickupagentid
type getassignedrequesttoagent struct {
	Pickupagentid int `uri:"pickupagentid" validate:"required,min=1"`
}

// Handler Function to get all pickuprequestid assigned to pickupagentid based on pickupagentid

// GetAssignedrequesttoagent godoc
//
//	@Summary		Get the pickuprequest assigned to the pickupagent
//	@Description	Get the pickuprequest assigned to the pickupagent
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			pickupagentid	path		int				true	"Pickupagentid"
//
// @Success 200 {object} []int
//
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/dashboard/getpickuprequest/{pickupagentid} [get]
func (dh *DashboardHandler) Assignedrequesttoagent(ctx *gin.Context) {

	var add getassignedrequesttoagent

	if err := ctx.ShouldBindUri(&add); err != nil {
		dh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	pickupRequestIDs, err := dh.svc.GetPickuprequestbyPickupagentid(ctx, add.Pickupagentid)
	if err != nil {
		dh.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	handleSuccess(ctx, pickupRequestIDs)
}

// This API will give all pickuprequestid raised by the customer
type getcustomerrequests struct {
	Customerid string `uri:"customerid" validate:"required,min=1"`
}

// Handler Function to get all pickuprequestid raised by the customer based on customerid

// GetCustomerrequests godoc
//
//	@Summary		Get the pickuprequest raised by the custoner
//	@Description	Get the pickuprequest raised by the customer
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			customerid	path		string				true	"Customerid"
//
// @Success		200	{object}	PickuprequestBasiccus	"Unassigned retrieved"
//
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/dashboard/customerpickuprequests/{customerid} [get]
func (dh *DashboardHandler) Customerrequests(ctx *gin.Context) {

	dh.log.Debug("Inside Customerrequests handler")
	var add getcustomerrequests

	if err := ctx.ShouldBindUri(&add); err != nil {
		dh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	dh.log.Debug("before calling repo layer from handler")
	// calling repo function to get list of pickuprequestid based on facilityid
	pickupDetailsList, err := dh.svc.GetPickuprequestbyCustomerid(ctx, add.Customerid)
	if err != nil {
		dh.log.Debug("error occured while repo calling:", err)
		handledbError(ctx, err)
		return
	}

	var responses []interface{}

	for _, pickupRequestID := range pickupDetailsList {
		pickupmain, err := dh.svc.Getcompletedetailsbasedobidcus(ctx, pickupRequestID)
		if err != nil {
			dh.log.Debug("error occured while repo calling:", err)
			handledbError(ctx, err)
			return
		}
		responses = append(responses, newPickuprequestBasiccus(pickupmain))
	}

	handleSuccess(ctx, responses)

}

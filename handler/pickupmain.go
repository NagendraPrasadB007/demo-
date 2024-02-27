package handler

import (
	"pickupmanagement/core/domain"

	"pickupmanagement/logger"
	repo "pickupmanagement/repo"
	"time"

	"github.com/gin-gonic/gin"
)

type PickupmainHandler struct {
	svc repo.PickupmainRepository
	log *logger.Logger
}

func NewPickupmainHandler(svc repo.PickupmainRepository, log *logger.Logger) *PickupmainHandler {
	return &PickupmainHandler{
		svc,
		log,
	}
}

type createPickupmain struct {
	Customerid                string    `json:"customerid" validate:"required"`
	Pickupdroptype            string    `json:"pickupdroptype" validate:"required"`
	Pickuplocation            string    `json:"pickuplocation" validate:"required"`
	Droplocation              string    `json:"droplocation" validate:"required"`
	Pickupscheduleslot        string    `json:"pickupscheduleslot" validate:"required"`
	Pickupscheduledate        time.Time `json:"pickupscheduledate" validate:"required"`
	Actualpickupdatetime      time.Time `json:"actualpickupdatetime" validate:"omitempty,required"`
	Pickupagentid             int       `json:"pickupagentid" vaidate:"omitempty,required"`
	Pickupfacilityid          string    `json:"pickupfacilityid" validate:"omitempty,required"`
	Pickupstatus              string    `json:"pickupstatus" validate:"omitempty,required"`
	Paymentstatus             string    `json:"paymentstatus" validate:"omitempty,required"`
	Pickupaddress             string    `json:"pickupaddress" validate:"required"`
	Domesticforeignidentifier string    `json:"domesticforeignidentifier" validate:"required"`
	Pickuplong                string    `json:"pickuplong" validate:"omitempty,required"`
	Pickuplat                 string    `json:"pickuplat" validate:"omitempty,required"`
	Pickuprequestedpincode    string    `json:"pickuprequestedpincode" validate:"omitempty,required"`
	//Customername              sql.NullString `json:"customername" validate:"required"`
	//Customername         null.String `json:"customername" validate:"required"`
	//Customername         pgtype.NullString `json:"customername" validate:"required"`
	//Customername null.String `json:"customername" validate:"required"`
	Customername string `json:"customername" validate:"required"`
	//Customername         pgtype.Text `json:"customername" validate:"required"`
	Customermobilenumber string    `json:"customermobilenumber" validate:"required"`
	Assigneddatetime     time.Time `json:"assigneddatetime"`
	//Createddatetime           time.Time `json:"createddatetime"`
	//Modifieddatetime          time.Time `json:"modifieddatetime"`
}

// Handler function to create a new request in pickupmain table

// CreatePickupmain godoc
//
//	@Summary        Insert data into Pickupmain
//	@Description    Insert data into Pickupmain
//	@Tags           Pickupmain
//	@Accept         json
//	@Produce        json
//	@Param          createPickupmain   body        createPickupmain   true    "createPickupmain"
//	@Success        200                     {object}    PickupmainResponse        "Pickupmain created"
//	@Failure        400                     {object}    errorValidResponse          "Validation error"
//	@Failure        401                     {object}    errorValidResponse          "Unauthorized error"
//	@Failure        403                     {object}    errorValidResponse          "Forbidden error"
//	@Failure        404                     {object}    errorValidResponse          "Data not found error"
//	@Failure        409                     {object}    errorValidResponse          "Data conflict error"
//	@Failure        500                     {object}    errorValidResponse          "Internal server error"
//	@Router         /pickupmain/singlereq [post]
func (ph *PickupmainHandler) createpickupmainrequest(ctx *gin.Context) {
	var req createPickupmain
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ph.log.Debug("error occurred while binding:", err)
		handleError(ctx, err)
		return
	}

	if !handleValidation(ctx, req) {
		return
	}

	pickupmain := domain.Pickupmain{
		Customerid:           req.Customerid,
		Pickupdroptype:       req.Pickupdroptype,
		Pickuplocation:       req.Pickuplocation,
		Droplocation:         req.Droplocation,
		Pickupscheduleslot:   req.Pickupscheduleslot,
		Pickupscheduledate:   req.Pickupscheduledate,
		Actualpickupdatetime: req.Actualpickupdatetime,
		Pickupagentid:        req.Pickupagentid,
		Pickupfacilityid:     req.Pickupfacilityid,
		Pickupstatus:         req.Pickupstatus,
		Paymentstatus:        req.Paymentstatus,
		//Createddatetime:           req.Createddatetime,
		Pickupaddress:             req.Pickupaddress,
		Domesticforeignidentifier: req.Domesticforeignidentifier,
		Pickuplong:                req.Pickuplong,
		Pickuplat:                 req.Pickuplat,
		Pickuprequestedpincode:    req.Pickuprequestedpincode,
		Customername:              req.Customername,
		Customermobilenumber:      req.Customermobilenumber,
		Assigneddatetime:          req.Assigneddatetime,
		//Modifieddatetime:          req.Modifieddatetime,
	}

	ph.log.Debug("finished mapping with pickupmain")

	_, err := ph.svc.Pickupmain(ctx, &pickupmain)
	//pickup, err := ph.svc.Pickupmain(ctx, &pickupmain)
	if err != nil {
		ph.log.Debug("error occurred while calling repo function:", err)
		handledbError(ctx, err)
	}

	rsp := newPickupmainResponse(&pickupmain)
	handleSuccess(ctx, rsp)

}

// bulk request

type createBulkPickupmain struct {
	Customerid           string    `json:"customerid" validate:"required"`
	Pickupdroptype       string    `json:"pickupdroptype" validate:"required"`
	Pickuplocation       string    `json:"pickuplocation" validate:"required"`
	Droplocation         string    `json:"droplocation" validate:"required"`
	Pickupscheduleslot   string    `json:"pickupscheduleslot" validate:"required"`
	Pickupscheduledate   time.Time `json:"pickupscheduledate" validate:"required"`
	Actualpickupdatetime time.Time `json:"actualpickupdatetime" validate:"omitempty,required"`
	Pickupagentid        int       `json:"pickupagentid" vaidate:"omitempty,required"`
	Pickupfacilityid     string    `json:"pickupfacilityid" validate:"omitempty,required"`
	//Pickupstatus              string    `json:"pickupstatus" validate:"omitempty,required"`
	Pickupstatus              string `json:"pickupstatus" validate:"omitempty,required,oneof=Unassigned Assigned Pickedup Cancelled"`
	Paymentstatus             string `json:"paymentstatus" validate:"omitempty,required"`
	Pickupaddress             string `json:"pickupaddress" validate:"required"`
	Domesticforeignidentifier string `json:"domesticforeignidentifier" validate:"required,oneof=domestic international"`
	Pickuplong                string `json:"pickuplong" validate:"omitempty,required"`
	Pickuplat                 string `json:"pickuplat" validate:"omitempty,required"`
	Pickuprequestedpincode    string `json:"pickuprequestedpincode" validate:"omitempty,required"`
	// Customername              sql.NullString `json:"customername" validate:"required"`
	// Customername         null.String `json:"customername" validate:"required"`
	Customername         string    `json:"customername" validate:"required"`
	Customermobilenumber string    `json:"customermobilenumber" validate:"required"`
	Assigneddatetime     time.Time `json:"assigneddatetime"`
}

// Handler function to create bulk insertion in pickupmain table

// CreateBulkPickupmain godoc
//
//	@Summary        Insert Bulk data into Pickupmain
//	@Description    Insert Bulk data into Pickupmain
//	@Tags           Pickupmain
//	@Accept         json
//	@Produce        json
//	@Param          createBulkPickupmain   body        createBulkPickupmain   true    "createBulkPickupmain"
//	@Success        200                     {object}    PickupmainResponse        "Pickupmain created"
//	@Failure        400                     {object}    errorValidResponse          "Validation error"
//	@Failure        401                     {object}    errorValidResponse          "Unauthorized error"
//	@Failure        403                     {object}    errorValidResponse          "Forbidden error"
//	@Failure        404                     {object}    errorValidResponse          "Data not found error"
//	@Failure        409                     {object}    errorValidResponse          "Data conflict error"
//	@Failure        500                     {object}    errorValidResponse          "Internal server error"
//	@Router         /pickupmain/bulk [post]
func (ph *PickupmainHandler) createbulkpickupmainrequest(ctx *gin.Context) {
	var req []createBulkPickupmain
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ph.log.Debug("error occurred while binding:", err)
		handleError(ctx, err)
		return
	}

	var pickupmainSlice []domain.Pickupmain

	for _, reqElement := range req {
		pickupmain := domain.Pickupmain{
			Customerid:                reqElement.Customerid,
			Pickupdroptype:            reqElement.Pickupdroptype,
			Pickuplocation:            reqElement.Pickuplocation,
			Droplocation:              reqElement.Droplocation,
			Pickupscheduleslot:        reqElement.Pickupscheduleslot,
			Pickupscheduledate:        reqElement.Pickupscheduledate,
			Actualpickupdatetime:      reqElement.Actualpickupdatetime,
			Pickupagentid:             reqElement.Pickupagentid,
			Pickupfacilityid:          reqElement.Pickupfacilityid,
			Pickupstatus:              reqElement.Pickupstatus,
			Paymentstatus:             reqElement.Paymentstatus,
			Pickupaddress:             reqElement.Pickupaddress,
			Domesticforeignidentifier: reqElement.Domesticforeignidentifier,
			Pickuplong:                reqElement.Pickuplong,
			Pickuplat:                 reqElement.Pickuplat,
			Pickuprequestedpincode:    reqElement.Pickuprequestedpincode,
			Customername:              reqElement.Customername,
			Customermobilenumber:      reqElement.Customermobilenumber,
			Assigneddatetime:          reqElement.Assigneddatetime,
		}

		pickupmainSlice = append(pickupmainSlice, pickupmain)
	}

	ph.log.Debug("finished mapping with pickupmain")

	_, err := ph.svc.PickupmainBatch(ctx, pickupmainSlice)
	//repoout, err := ph.svc.PickupmainBatch(ctx, pickupmainSlice)
	if err != nil {
		ph.log.Debug("error occurred while calling repo function:", err)
		handledbError(ctx, err)
		return
	}

	rsp := pickupmainSlice
	//rsp := newBulkPickupmainResponse(pickupmainSlice)
	//rsp := repoout
	ph.log.Debug("rsp:", rsp)

	handleSuccess(ctx, rsp)

}

/*

type Assigning struct {
	//Pickuprequestid int `json:"pickuprequestid" validate:"required"`
	Pickuprequestid []int `json:"pickuprequestid" validate:"required"`
}

func (ph *PickupmainHandler) AssignPickuprequest(ctx *gin.Context) {
	var req Assigning
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ph.log.Debug("error occurred while binding:", err)
		handleError(ctx, err)
		return
	}

	if !handleValidation(ctx, req) {
		return
	}

	idStr := ctx.Param("pickupagentid")
	Pickupagentid, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(ctx, err)
		ph.log.Debug("error occured while converting addressid:", err)
		return
	}

	//for loop for taking multiple pickuprequestid in the body for the same pickupagentid in the url
	for _, pickupRequestID := range req.Pickuprequestid {
		assigned := domain.Pickupmain{
			Pickuprequestid: pickupRequestID,
			Pickupagentid:   Pickupagentid,
		}

		err = ph.svc.AssigningPickuprequests(ctx, &assigned)
		if err != nil {
			ph.log.Debug("error occurred while calling repo function:", err)
			handledbError(ctx, err)
			return
		}
	}

	handleSuccess(ctx, "Pickuprequest Assigned successfully")

}
*/

//AssigningProcess

type Assigning struct {
	Pickuprequestid int `json:"pickuprequestid" validate:"required"`
	Pickupagentid   int `json:"pickupagentid" validate:"required"`
}

// Handler function to assign pickuprequest to pickupagent

// Assiging pickuprequest to pickupagent godoc
//
//	@Summary		Assiging pickuprequest to pickupagent
//	@Description	Assiging pickuprequest to pickupagent
//	@Tags			Pickupmain
//	@Accept			json
//	@Produce		json
//	@Param			Assigning	body		Assigning	true	"Assigning"
//	@Success 200 {string} string "Pickup requests assigned successfully"
//	@Failure		400						{object}	errorValidResponse			"Validation error"
//	@Failure		401						{object}	errorValidResponse			"Unauthorized error"
//	@Failure		403						{object}	errorValidResponse			"Forbidden error"
//	@Failure		404						{object}	errorValidResponse			"Data not found error"
//	@Failure		409						{object}	errorValidResponse			"Data conflict error"
//	@Failure		500						{object}	errorValidResponse			"Internal server error"
//	@Router			/pickupmain/assigning [put]
func (ph *PickupmainHandler) AssigningProcess(ctx *gin.Context) {

	var req []Assigning // Array to hold multiple assignments

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ph.log.Debug("error occurred while binding:", err)
		handleError(ctx, err)
		return
	}
	// Create empty slices to hold multiple request IDs and agent IDs
	var requestIDs []int
	var agentIDs []int

	for _, assignment := range req {

		requestIDs = append(requestIDs, assignment.Pickuprequestid)
		agentIDs = append(agentIDs, assignment.Pickupagentid)
	}

	// Call the service function with the accumulated slices
	err := ph.svc.AssignPickupAgents(ctx, requestIDs, agentIDs)
	if err != nil {
		ph.log.Debug(err)
		handledbError(ctx, err)
		return
	}
	handleSuccess(ctx, "Pickup requests assigned successfully")
}

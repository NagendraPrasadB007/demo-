package handler

import (
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"

	"github.com/gin-gonic/gin"
)

type FetchdetailsHandler struct {
	svc repo.FetchdetailsRepository
	log *logger.Logger
}

func NewFetchdetailsHandler(svc repo.FetchdetailsRepository, log *logger.Logger) *FetchdetailsHandler {
	return &FetchdetailsHandler{
		svc,
		log,
	}
}

//getdetailsbypickuprequestid

type getpickuprequestid struct {
	Pickuprequestid int `uri:"pickuprequestid"`
}

/*
func (fh *FetchdetailsHandler) getdetailsbypickuprequestid(ctx *gin.Context) {

	var add getpickuprequestid

	if err := ctx.ShouldBindUri(&add); err != nil {
		fh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	returnvalue, err := fh.svc.Identify(ctx, add.Pickuprequestid)
	if err != nil {
		fh.log.Debug(err)
		return
	}

	result := returnvalue.Domesticforeignidentifier

	if result == "domestic" {

		pickupmain, domestic, tariff, payment, err := fh.svc.GetdetailsByPickuprequestID(ctx, add.Pickuprequestid)
		if err != nil {
			fh.log.Debug("error occured while repo calling:", err)
			handledbError(ctx, err)
			return
		}

		rsp := newFetchPickupRequestResponse(pickupmain, domestic, tariff, payment)
		handleSuccess(ctx, rsp)
	} else {
		pickupmain, international, subpiece, tariff, payment, err := fh.svc.GetIntdetailsByPickuprequestID(ctx, add.Pickuprequestid)
		if err != nil {
			fh.log.Debug("error occured while repo calling:", err)
			handledbError(ctx, err)
			return
		}

		rsp := newIntFetchPickupRequestResponse(pickupmain, international, subpiece, tariff, payment)
		handleSuccess(ctx, rsp)

	}

}
*/

// Handler function to fetch pickuprequest details based on pickuprequestid

// GetPickuprequestdetails godoc
//
//	@Summary		Get a Pickuprequestdetails
//	@Description	Get a Pickuprequestdetails
//	@Tags			Pickuprequest
//	@Accept			json
//	@Produce		json
//	@Param			pickuprequestid	path		int				true	"Pickuprequestid"
//	@Success		200	{object}	RaisedPickupRequestResponseDom	"Category retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/getdetails/{pickuprequestid} [get]
func (fh *FetchdetailsHandler) getdetailsbypickuprequestid(ctx *gin.Context) {

	var add getpickuprequestid

	if err := ctx.ShouldBindUri(&add); err != nil {
		fh.log.Debug("error occured during binding:", err)
		handleError(ctx, err)
		return
	}
	if !handleValidation(ctx, add) {
		return
	}

	// to get to know wheather it is domestic article or international article
	fh.log.Debug("before calling identifyPR")
	returnvalue, err := fh.svc.Identify(ctx, add.Pickuprequestid)
	if err != nil {
		fh.log.Debug(err)
		return
	}
	fh.log.Debug("after getting values from identify")
	result := returnvalue.Domesticforeignidentifier

	if result == "domestic" {

		pickupmain, domestic, tariff, payment, err := fh.svc.GetdetailsByPickuprequestID(ctx, add.Pickuprequestid)
		if err != nil {
			fh.log.Debug("error occured while repo calling:", err)
			handledbError(ctx, err)
			return
		}

		rsp := newFetchPickupRequestResponse(pickupmain, domestic, tariff, payment)
		handleSuccess(ctx, rsp)
	} else {
		pickupmain, international, subpiece, tariff, payment, err := fh.svc.GetIntdetailsByPickuprequestID(ctx, add.Pickuprequestid)
		fh.log.Debug("subpiece:", subpiece)
		if err != nil {
			fh.log.Debug("error occured while repo calling:", err)
			handledbError(ctx, err)
			return
		}
		rsp := newFetchPickupRequestResponseInt1(pickupmain, international, tariff, payment)
		handleSuccess(ctx, rsp)

	}

}

package handler

import (
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CancelpickuprequestHandler struct {
	svc repo.CancelpickuprequestRepository
	log *logger.Logger
}

func NewCancelpickuprequestHandler(svc repo.CancelpickuprequestRepository, log *logger.Logger) *CancelpickuprequestHandler {
	return &CancelpickuprequestHandler{
		svc,
		log,
	}
}

// Handler function used to cancel pickuprequest based on pickuprequestid

// CancelPickuprequest godoc
//
//	@Summary		cancel a pickup request
//	@Description	cancel a pickup request based on pickuprequestid
//	@Tags			Pickuprequest
//	@Accept			json
//	@Produce		json
//	@Param			pickuprequestid				path		int					true	"cancel pickuprequest"
//
// @Success 200 {string} string "Pickuprequest Cancelled successfully"
//
//	@Failure		400						{object}	errorValidResponse			"Validation error"
//	@Failure		401						{object}	errorValidResponse			"Unauthorized error"
//	@Failure		403						{object}	errorValidResponse			"Forbidden error"
//	@Failure		404						{object}	errorValidResponse			"Data not found error"
//	@Failure		409						{object}	errorValidResponse			"Data conflict error"
//	@Failure		500						{object}	errorValidResponse			"Internal server error"
//	@Router			/cancel/{pickuprequestid} [put]
func (ch *CancelpickuprequestHandler) cancelpickuprequestbypickuprequestid(ctx *gin.Context) {

	idStr := ctx.Param("pickuprequestid")
	// here idStr is of type string i am converting string into int because pickuprequestid in my DB is int
	pickuprequestID, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(ctx, err)
		ch.log.Debug("error occured while converting addressid:", err.Error())
		return
	}
	err = ch.svc.Cancelpickuprequestdetails(ctx, pickuprequestID)
	if err != nil {
		handledbError(ctx, err)
		ch.log.Debug("error occured while repo calling:", err.Error())
		return
	}
	handleSuccess(ctx, "Pickuprequest Cancelled successfully")
}

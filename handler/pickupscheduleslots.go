package handler

import (
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"

	"github.com/gin-gonic/gin"
)

type PickupscheduleslotsHandler struct {
	svc repo.PickupscheduleslotsRepository
	log *logger.Logger
}

func NewpickupscheduleslotsHandler(svc repo.PickupscheduleslotsRepository, log *logger.Logger) *PickupscheduleslotsHandler {
	return &PickupscheduleslotsHandler{
		svc,
		log,
	}
}

// Handler function to get all pickupschedulslots

// GetSchedules godoc
//
//	@Summary		Get a PickupScheduleslots
//	@Description	Get a PickupScheduleslots
//	@Tags			Common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pickupscheduleslotsResponse 	"Schedule retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/pickupscheduleslots/ [get]
func (ps *PickupscheduleslotsHandler) getallscheduleslots(ctx *gin.Context) {

	slots, err := ps.svc.Listslots(ctx)
	if err != nil {
		ps.log.Debug("error occurred while calling repo function:", err)
		handledbError(ctx, err)
		return
	}

	ps.log.Debug("slots:", slots)

	//no need to append in repo only append happened
	// for _, slot := range slots {
	// 	slotslists = append(slotslists, newPickupscheduleslotsResponse(&slot))
	// }

	//rsp := slotslists
	rsp := slots

	handleSuccess(ctx, rsp)

}

// Hanlder function to get all pickup agents

// GetPickupagent godoc
//
//	@Summary		Get a PickupagentList
//	@Description	Get a PickupagentList
//	@Tags			Common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.Pickupagent 	"Pickupagent list retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/pickupagentlist/ [get]
func (ps *PickupscheduleslotsHandler) getallpickupagent(ctx *gin.Context) {

	lists, err := ps.svc.Pickupagentlist(ctx)
	if err != nil {
		ps.log.Debug("error occurred while calling repo function:", err)
		handledbError(ctx, err)
		return
	}

	rsp := lists

	handleSuccess(ctx, rsp)

}

// Hanlder function to get remarks

// GetRemarks godoc
//
//	@Summary		Get a Remarks
//	@Description	Get a Remarks
//	@Tags			Common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.Remarks 	"Remarks list retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/remarkslist/ [get]
func (ps *PickupscheduleslotsHandler) getallremarks(ctx *gin.Context) {

	lists, err := ps.svc.Remarkslist(ctx)
	if err != nil {
		ps.log.Debug("error occurred while calling repo function:", err)
		handledbError(ctx, err)
		return
	}

	rsp := lists

	handleSuccess(ctx, rsp)

}

//getalladdresstypelist

// GetAddresstypelist godoc
//
//	@Summary		Get a Addresstypelist
//	@Description	Get a Addresstypelist
//	@Tags			Common
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.Addresstype 	"Addresstype list retrieved"
//	@Failure		400	{object}	errorValidResponse		"Validation error"
//	@Failure		404	{object}	errorValidResponse		"Data not found error"
//	@Failure		500	{object}	errorValidResponse		"Internal server error"
//	@Router			/addresstypelist/ [get]
func (ps *PickupscheduleslotsHandler) getalladdresstypelist(ctx *gin.Context) {

	lists, err := ps.svc.Addresstypelist(ctx)
	if err != nil {
		ps.log.Debug("error occurred while calling repo function:", err)
		handledbError(ctx, err)
		return
	}

	rsp := lists

	handleSuccess(ctx, rsp)

}

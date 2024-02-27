package route

import (
	handler "pickupmanagement/handler"
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"
)

func LoggerInit(loglevel string) *logger.Logger {
	log := logger.New(loglevel)
	return log
}
func Routes(db *repo.DB, log *logger.Logger) (router *handler.Router, err1 error) {

	addressdetailsRepo := repo.NewAddressdetailsRepository(db, log)
	addressdetailsHandler := handler.NewAddressdetailsHandler(*addressdetailsRepo, log)

	pickupmainRepo := repo.NewPickupmainRepository(db, log)
	pickupmainHandler := handler.NewPickupmainHandler(*pickupmainRepo, log)

	raisepickuprequestRepo := repo.NewRaisepickuprequestRepository(db, log)
	raisepickuprequestHandler := handler.NewRaisepickuprequestHandler(*raisepickuprequestRepo, log)

	pickupscheduleslotsRepo := repo.NewPickupscheduleslotsRepository(db, log)
	pickupscheduleslotsHandler := handler.NewpickupscheduleslotsHandler(*pickupscheduleslotsRepo, log)

	imageRepo := repo.NewImageRepository(db, log)
	imageHandler := handler.NewimageHandler(*imageRepo, log)

	fetchRepo := repo.NewFetchdetailsRepository(db, log)
	fetchHandler := handler.NewFetchdetailsHandler(*fetchRepo, log)

	updateRepo := repo.NewUpdatedetailsRepository(db, log)
	updateHandler := handler.NewUpdatedetailsHandler(*updateRepo, log)

	cancelRepo := repo.NewCancelpickuprequestRepository(db, log)
	cancelHandler := handler.NewCancelpickuprequestHandler(*cancelRepo, log)

	dashboardRepo := repo.NewDashboardRepository(db, log)
	dashboardHandler := handler.NewDashboardHandler(*dashboardRepo, log)

	router, err1 = handler.NewRouter(
		*addressdetailsHandler,
		*pickupmainHandler,
		*pickupscheduleslotsHandler,
		*imageHandler,
		*raisepickuprequestHandler,
		*fetchHandler,
		*updateHandler,
		*cancelHandler,
		*dashboardHandler,
	)
	return router, err1

}

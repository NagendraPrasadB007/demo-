// @version 1.0
// @description This is a  Demo Template API with Swagger documentation
// @host localhost:3030
// @BasePath /pickup/v1

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	config "pickupmanagement/config"
	"pickupmanagement/logger"
	repo "pickupmanagement/repo"
	r "pickupmanagement/route"
	"syscall"
	"time"
)

func init() {
	config.Load()
}

func LoggerInit(loglevel string) *logger.Logger {
	log := logger.New(loglevel)
	return log
}

func main() {

	log := LoggerInit(os.Getenv("LOG_LEVEL"))
	listenAddr := ":" + os.Getenv("HTTP_PORT")
	ctx := context.Background()

	db, err := repo.NewDB(ctx)
	if err != nil {
		log.Warn("error in db connection %s", err)
	}
	defer db.Close()
	log.Info("Successfully connected to the database %s", os.Getenv("DB_CONNECTION"))

	router, err := r.Routes(db, log)

	//Handlers and Repo initialization

	/*
		pickuprequestRepo := repo.NewPickuprequestRepository(db, log)
		pickuprequestHandler := handler.NewPickuprequestHandler(*pickuprequestRepo, log)

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

		router, err := handler.NewRouter(
			*pickuprequestHandler,
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
	*/

	if err != nil {
		log.Warn("Error initializing router %s", err)
		os.Exit(1)
	}

	log.Info("Starting the HTTP server: %s", listenAddr)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("HTTP_PORT"),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Info("Received signal...%s", sig)

	duration, err := time.ParseDuration(os.Getenv("SHUTDOWN_TIME"))
	if err != nil {
		log.Fatal("Error in parsing duration", err)

	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Info("Server exiting")

}

package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	_ "pickupmanagement/docs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	addressdetailsHandler AddressdetailsHandler,
	pickupmainHandler PickupmainHandler,
	pickupscheduleslotsHandler PickupscheduleslotsHandler,
	imageHandler ImageHandler,
	raisepickuprequestHandler RaisepickuprequestHandler,
	fetchpickuprequestHandler FetchdetailsHandler,
	updatedetailsHandler UpdatedetailsHandler,
	cancelpickuprequestHandler CancelpickuprequestHandler,
	dashboardHandler DashboardHandler,
) (*Router, error) {
	// Disable debug mode and write logs to file in production
	env := os.Getenv("APP_ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)

		logFile, _ := os.Create("gin.log")
		gin.DefaultWriter = io.Writer(logFile)
	}

	// CORS
	config := cors.DefaultConfig()
	//allowedOrigins := os.Getenv("HTTP_ALLOWED_ORIGINS")
	//originsList := strings.Split(allowedOrigins, ",")
	//config.AllowOrigins = originsList

	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"*"}
	config.AllowBrowserExtensions = true
	config.AllowMethods = []string{"*"}

	router := gin.New()
	router.RedirectTrailingSlash = false
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": []string{"Invalid Path"},
			"errorno": []string{"INV1"},
		})

	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"success": false,
			"message": []string{"Method not Allowed"},
			"errorno": []string{"MD01"},
		})

	})

	if env == "production" {
		router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config), ValidateContentType([]string{"application/json"}))
	}
	//This is working
	if env == "development" {
		router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config), ValidateContentType([]string{"application/json", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7", "text/css,*/*;q=0.1", "application/json,*/*", "*/*", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8"}))
	}

	//Pavani mam
	// if env == "development" {
	// 	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config), ValidateContentType([]string{"application/json", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,/;q=0.8,application/signed-exchange;v=b3;q=0.7", "image/avif,image/webp,image/apng,image/svg+xml,image/,/;q=0.8", "application/json,/", "/*"}))
	// }

	// if env == "development" {
	// 	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config), ValidateContentType([]string{"application/json", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8", "text/css,*/*;q=0.1", "application/json,*/*", "*/*"}))
	// }

	//router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config))

	// router.Use(
	// 	gin.LoggerWithFormatter(customLogger),
	// 	gin.Recovery(),
	// 	cors.New(config),
	// 	ValidateContentType([]string{"application/json", "image/png"}),
	// )

	/*	// Custom validators
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			if err := v.RegisterValidation("user_role", userRoleValidator); err != nil {
				return nil, err
			}

			if err := v.RegisterValidation("payment_type", paymentTypeValidator); err != nil {
				return nil, err
			}

		}*/

	// Swagger
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/pickup/v1")
	{
		addressdetails := v1.Group("/addressdetails")
		{
			addressdetails.GET("/:customerid", addressdetailsHandler.getaddressdetailsbyid)
			addressdetails.POST("/", addressdetailsHandler.createaddressforcustomer)
			addressdetails.PUT("/:addressid", addressdetailsHandler.updateadress)
			addressdetails.DELETE("/:addressid", addressdetailsHandler.deleteadress)
		}

		pickupmain := v1.Group("/pickupmain")
		{
			pickupmain.POST("/singlereq", pickupmainHandler.createpickupmainrequest)
			pickupmain.POST("/bulk", pickupmainHandler.createbulkpickupmainrequest)
			//pickupmain.PUT("/assign/:pickupagentid", pickupmainHandler.AssignPickuprequest)
			pickupmain.PUT("/assigning", pickupmainHandler.AssigningProcess)
		}

		pickupscheduleslots := v1.Group("/pickupscheduleslots")
		{
			pickupscheduleslots.GET("/", pickupscheduleslotsHandler.getallscheduleslots)
		}

		pickupagent := v1.Group("/pickupagentlist")
		{
			pickupagent.GET("/", pickupscheduleslotsHandler.getallpickupagent)
		}

		remarks := v1.Group("/remarkslist")
		{
			remarks.GET("/", pickupscheduleslotsHandler.getallremarks)
		}

		addresstype := v1.Group("/addresstypelist")
		{
			addresstype.GET("/", pickupscheduleslotsHandler.getalladdresstypelist)
		}

		image := v1.Group("/image")
		{
			image.POST("/", imageHandler.uploadimage)
			image.GET("/:imageid", imageHandler.getimagebyimageid)
		}

		raisepickuprequest := v1.Group("/raisepickup")
		{
			raisepickuprequest.POST("/", raisepickuprequestHandler.raiserequest)
		}

		fetchpickuprequest := v1.Group("/getdetails")
		{
			fetchpickuprequest.GET("/:pickuprequestid", fetchpickuprequestHandler.getdetailsbypickuprequestid)
		}

		updatepickuprequest := v1.Group("/updatedetails")
		{
			updatepickuprequest.PUT("/:pickuprequestid", updatedetailsHandler.updatedetailsbypickuprequestid)
		}

		cancelpickuprequest := v1.Group("/cancel")
		{
			cancelpickuprequest.PUT("/:pickuprequestid", cancelpickuprequestHandler.cancelpickuprequestbypickuprequestid)
		}

		dashboardrequest := v1.Group("/dashboard")
		{
			dashboardrequest.GET("/count/:facilityid", dashboardHandler.countdetails)
			dashboardrequest.GET("/Assigned/:facilityid", dashboardHandler.Assignedlist)
			dashboardrequest.GET("/Unassigned/:facilityid", dashboardHandler.Unassignedlist) //pincode
			dashboardrequest.GET("/pickuprequest/:facilityid", dashboardHandler.Pickuprequestlist)
			dashboardrequest.GET("/pickuprequests/:pincode", dashboardHandler.Pickuprequestlistbypincode)
			dashboardrequest.GET("/getpickuprequest/:pickupagentid", dashboardHandler.Assignedrequesttoagent)
			dashboardrequest.GET("/customerpickuprequests/:customerid", dashboardHandler.Customerrequests)
		}
	}
	//}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

// customLogger is a custom Gin logger
func customLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s [%s]\"\n",
		param.TimeStamp.Format(time.RFC1123),
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency.Round(time.Millisecond),
		param.Request.UserAgent(),
	)
}

func ValidateContentType(allowedTypes []string) gin.HandlerFunc {
	//acceptedTypes = append(acceptedTypes, "image/png")
	return func(c *gin.Context) {

		//contentType := c.GetHeader("Content-Type")
		contentType := c.GetHeader("Accept")
		//log.Println("Received Content-Type:", contentType)

		// Check if the Content-Type is in the allowedTypes
		validContentType := false
		for _, allowedType := range allowedTypes {
			if contentType == allowedType {
				validContentType = true
				break
			}
		}

		if !validContentType {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{
				"success": false,
				"message": []string{"Invalid Content-Type. Supported types are: " + fmt.Sprintf("%v", allowedTypes)},
				"errorno": []string{"USP1"},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// func ValidateContentType(acceptedTypes []string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	   contentType := c.Request.Header.Get("Content-Type")
// 	   log.Println("Received Content-Type:", contentType)

// 	   // Your validation logic here

// 	   if !isValidContentType(contentType, acceptedTypes) {
// 		  c.JSON(http.StatusBadRequest, gin.H{
// 			 "errorno":  ["USP1"]
// 			 "message":  ["Invalid Content-Type. Supported types are: " + strings.Join(acceptedTypes, " ")],
// 			 "success":  false,
// 		  })
// 		  c.Abort()
// 		  return
// 	   }

// 	   // Continue with the next middleware or handlers
// 	   c.Next()
// 	}
//  }

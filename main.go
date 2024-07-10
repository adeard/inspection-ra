package main

import (
	"inspection-ra/config"
	"inspection-ra/docs"
	"inspection-ra/middlewares"
	"inspection-ra/modules/attachment"
	"inspection-ra/modules/auth"
	"inspection-ra/modules/configuration"
	"inspection-ra/modules/mob"
	"inspection-ra/modules/objpart"
	"inspection-ra/modules/plan"
	"inspection-ra/modules/runacct"
	auth_v2 "inspection-ra/modules/v2/auth"
	"inspection-ra/modules/vehicletype"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @Host https://service.indoagri.co.id/InspectionRA
// @title API SWAGGER FOR INSPECTION RA API SERVICE
// @version 1.0.0
// @description INSPECTION RA API SERVICE
// @termsOfService http://swagger.io/terms/

// @contact.name ICT INDOAGRI
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath

func main() {
	db := config.Connect()

	docs.SwaggerInfo.BasePath = "/InspectionRA"

	router := gin.Default()
	router.Use(cors.AllowAll())
	router.Use(middlewares.LoggerMiddleware())
	router.GET("InspectionRA/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"title":         "Inspection RA API Service",
			"documentation": "/swagger/index.html",
		})
	})

	router.GET("InspectionRA/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("InspectionRA/api")

	v1 := api.Group("v1")
	mob.NewMobHandler(v1, mob.MobRegistry(db))
	auth.NewAuthHandler(v1, auth.AuthRegistry(db))
	plan.NewPlanHandler(v1, plan.PlanRegistry(db))
	runacct.NewRunAcctHandler(v1, runacct.RunAcctRegistry(db))
	objpart.NewObjPartHandler(v1, objpart.ObjPartRegistry(db))
	attachment.NewAttachmentHandler(v1, attachment.AttachmentRegistry(db))
	vehicletype.NewVehicleTypeHandler(v1, vehicletype.VehicleTypeRegistry(db))
	configuration.NewConfigurationHandler(v1, configuration.ConfigurationRegistry(db))

	v2 := api.Group("v2")
	auth_v2.NewAuthHandler(v2, auth_v2.AuthRegistry(db))

	// router.Run(":86")

	// Mengatur mode GIN menjadi release
	gin.SetMode(gin.ReleaseMode)

	//Penyesuaian Port ke IIS
	port := "86"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	// Menampilkan log koneksi sukses
	log.Println("App Service run in port:", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		// Menampilkan log ketika koneksi gagal
		log.Fatal("Connection Fail -> port "+port+":", err)
	}
}

package http

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	dbconfig "github.com/hallex-abreu/landing-page-ms/api/infra/database/mysql/db-config"
	healthController "github.com/hallex-abreu/landing-page-ms/api/infra/http/controllers/health-controller"
	leadhController "github.com/hallex-abreu/landing-page-ms/api/infra/http/controllers/lead-controller"
	"github.com/joho/godotenv"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/health", healthController.Alive)
	router.GET("/lead", leadhController.FindAllLeads)
	router.POST("/lead", leadhController.RegisterLead)

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconfig.Connection()

	port := os.Getenv("PORT")

	log.Println("Running server in port " + port)
	router.Run(":" + port)
}

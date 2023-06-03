package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/superjantung/pet-adoption-api/controllers"
	"github.com/superjantung/pet-adoption-api/models"
)

const rootPath = "/"

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	db, err := models.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET(rootPath, func(c *gin.Context) {
		response := ResponseData{
			Message: "Hello, world!",
		}
		c.JSON(http.StatusOK, response)
	})

	r.GET("/pets", controllers.FindPets)
	r.GET("/pets/:id", controllers.FindPet)
	r.POST("/pets", controllers.CreatePet)
	r.PATCH("/pets/:id", controllers.UpdatePet)
	r.DELETE("/pets/:id", controllers.DeletePet)

	r.GET("/adopters", controllers.FindAdopters)
	r.GET("/adopters/:id", controllers.FindAdopter)
	r.POST("/adopters", controllers.CreateAdopter)
	r.PATCH("/adopters/:id", controllers.UpdateAdopter)
	r.DELETE("/adopters/:id", controllers.DeleteAdopter)

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = ":8080"
	}

	err = r.Run(serverAddress)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

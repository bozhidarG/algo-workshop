package main

import (
	"github.com/gin-gonic/gin"
	"workshops/dependencyInjection/go/database"
	"workshops/dependencyInjection/go/handlers"
)

func main() {
	router := gin.Default()

	responseHandler := InitializeApp()
	//responseHandlerMock := InitializeApp()

	auth := router.Group("/auth")
	{
		auth.GET("/", responseHandler.IndexEndpoint)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func InitializeApp() handlers.ResponseHandler {
	realDatabase := database.NewDatabase()
	responseHandler := handlers.NewResponseHandler(realDatabase)
	return responseHandler
}

func InitializeMockApp() handlers.ResponseHandler {
	mockDatabase := database.NewMockDatabase()
	responseHandler := handlers.NewResponseHandler(mockDatabase)
	return responseHandler
}
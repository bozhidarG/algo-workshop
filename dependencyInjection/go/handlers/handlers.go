package handlers

import (
	"github.com/gin-gonic/gin"
	"workshops/dependencyInjection/go/database"
)

type ResponseHandler struct {
	database database.Database
}

func NewResponseHandler(db database.Database) ResponseHandler {
	return ResponseHandler{database: db}
}

//show login/registration form index file, maybe won't be needed
func (r *ResponseHandler) IndexEndpoint(c *gin.Context) {
	r.database.GetUser()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

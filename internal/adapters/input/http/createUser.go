package controller

import (
	"api/internal/infrastructure/config/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Invalid request body")
	c.JSON(err.Code, err)
}

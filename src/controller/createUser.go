package controller

import (
	"api/src/configuration/logger"
	"api/src/configuration/validation"
	"api/src/controller/model/request"
	"api/src/controller/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser.controller",
		zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error in user info validation", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User create successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}

package controller

import (
	"api/src/configuration/rest_err"
	"api/src/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Invalid request body %s", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

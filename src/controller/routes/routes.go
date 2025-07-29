package routes

import (
	"api/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:user_id", controller.FindUserById)
	r.GET("/user/:user_email", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:user_id", controller.UpdateUser)
	r.DELETE("/user/:user_id", controller.DeleteUser)
}

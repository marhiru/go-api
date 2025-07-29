package main

import (
	"api/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

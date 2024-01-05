package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit-jaisinghani/gin-gorm-rest/config"
	"github.com/rohit-jaisinghani/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.ArtistRoute(router)
	router.Run(":8080")
}

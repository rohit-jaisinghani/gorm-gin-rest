package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit-jaisinghani/gin-gorm-rest/controller"
)

// Created different routes for artist search
func ArtistRoute(router *gin.Engine) {
	router.GET("/getallartist", controller.GetArtists)
	router.GET("/getartist/:name", controller.GetArtist)
	router.GET("/getartistIsrc/:isrc", controller.GetArtistByIsrc)
	router.POST("/getartist", controller.CreateArtist)

}

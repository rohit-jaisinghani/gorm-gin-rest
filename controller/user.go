package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rohit-jaisinghani/gin-gorm-rest/config"
	"github.com/rohit-jaisinghani/gin-gorm-rest/models"
	"github.com/rohit-jaisinghani/gin-gorm-rest/processor"
)

// GetArtists will return all the artist from DB
func GetArtists(c *gin.Context) {
	users := []models.Artist{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

// GetArtist will return the artist from DB based on name.
func GetArtist(c *gin.Context) {
	name := c.Param("name")
	users := []models.Artist{}
	config.DB.Where("artist_name LIKE ?", "%"+name+"%").Find(&users)
	c.JSON(200, &users)
}

// GetArtistByIsrc will return the artist from DB based on isrc.
func GetArtistByIsrc(c *gin.Context) {
	isrc := c.Param("isrc")

	users := []models.Artist{}
	config.DB.Where("isrc LIKE ?", "%"+isrc+"%").Find(&users)
	c.JSON(200, &users)
}

// CreateArtist will store the artist into db
func CreateArtist(c *gin.Context) {
	var track models.Track
	c.BindJSON(&track)
	token := processor.GetToken()
	ss := processor.SearchForArtist(token, track.TrackName)

	if len(ss) == 0 {
		c.JSON(501, "No records exist for this ISRC")
	} else {
		for _, val := range ss {
			config.DB.Create(&val)
		}
		c.JSON(200, "Records Inserted")
	}

}

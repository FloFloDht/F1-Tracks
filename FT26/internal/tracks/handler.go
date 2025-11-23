package tracks

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/tracks", handleGetTracks)
	r.GET("/tracks/:id", handleGetTrack)
}

func handleGetTracks(c *gin.Context) {
	tracks, err := GetAllTracks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tracks)
}

func handleGetTrack(c *gin.Context) {
	id := c.Param("id")

	track, err := GetTrackByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Circuit non trouvé"})
		return
	}

	c.JSON(200, track)
}

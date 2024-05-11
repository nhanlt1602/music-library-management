package routes

import (
	"music-library-management/controllers"
	"music-library-management/validators"

	"github.com/gin-gonic/gin"
)

func MusicTrackRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	auth := router.Group("/music-track")
	{
		auth.POST(
			"",
			validators.CreateMusicTrackValidator(),
			controllers.CreateMusicTrack,
		)
	}
}

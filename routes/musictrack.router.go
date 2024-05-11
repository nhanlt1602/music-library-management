package routes

import (
	"music-library-management/controllers"
	"music-library-management/validators"

	"github.com/gin-gonic/gin"
)

func MusicTrackRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	auth := router.Group("/music-tracks")
	{
		auth.POST(
			"",
			validators.CreateMusicTrackValidator(),
			controllers.CreateMusicTrack,
		)
		auth.PUT(
			"/:id",
			validators.UpdateMusicTrackValidator(),
			controllers.UpdateMusicTrack,
		)
		auth.GET(
			"/:id",
			controllers.GetMusicTrackById,
		)
		auth.GET(
			"",
			controllers.GetMusicTracks,
		)
		auth.DELETE(
			"/:id",
			controllers.DeleteMusicTrack,
		)
	}
}

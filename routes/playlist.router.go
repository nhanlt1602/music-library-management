package routes

import (
	"music-library-management/controllers"
	"music-library-management/validators"

	"github.com/gin-gonic/gin"
)

func PlaylistRoute(router *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	auth := router.Group("/playlists")
	{
		auth.POST(
			"",
			validators.CreatePlaylistValidator(),
			controllers.CreatePlaylist,
		)
		auth.PUT(
			"/:id",
			validators.UpdatePlaylistValidator(),
			controllers.UpdatePlaylist,
		)
		auth.GET(
			"/:id",
			controllers.GetPlaylistById,
		)
		auth.GET(
			"",
			controllers.GetPlaylists,
		)
		auth.DELETE(
			"/:id",
			controllers.DeletePlaylist,
		)
	}
}

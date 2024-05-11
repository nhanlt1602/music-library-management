package validators

import (
	"music-library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreatePlaylistValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var createPlaylist models.PlaylistRequest
		_ = c.ShouldBindBodyWith(&createPlaylist, binding.JSON)

		if err := createPlaylist.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func UpdatePlaylistValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updatePlaylist models.PlaylistRequest
		_ = c.ShouldBindBodyWith(&updatePlaylist, binding.JSON)

		if err := updatePlaylist.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func GetPlaylistsValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		page := c.DefaultQuery("page", "0")
		size := c.DefaultQuery("size", "0")

		if page == "0" || size == "0" {
			models.SendErrorResponse(c, http.StatusBadRequest, "page and size are required")
			return
		}

		c.Next()
	}
}

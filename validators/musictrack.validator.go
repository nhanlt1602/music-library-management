package validators

import (
	"music-library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateMusicTrackValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var createMusicTrack models.MusicTrackRequest
		_ = c.ShouldBindBodyWith(&createMusicTrack, binding.JSON)

		if err := createMusicTrack.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

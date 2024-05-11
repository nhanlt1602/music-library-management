package validators

import (
	"music-library-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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

func UpdateMusicTrackValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updateMusicTrack models.MusicTrackRequest
		_ = c.ShouldBindBodyWith(&updateMusicTrack, binding.JSON)

		if err := updateMusicTrack.Validate(); err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Next()
	}
}

func GetMusicTracksValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		page := c.DefaultQuery("page", "0")
		err := validation.Validate(page, is.Int)
		if err != nil {
			models.SendErrorResponse(c, http.StatusBadRequest, "invalid page: "+page)
			return
		}

		c.Next()
	}
}

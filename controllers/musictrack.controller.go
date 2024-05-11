package controllers

import (
	"music-library-management/models"
	"music-library-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateMusicTrack godoc
// @Summary      Create Music Track
// @Description  create a new music track
// @Tags         music-tracks
// @Accept       json
// @Produce      json
// @Param        req body models.MusicTrackRequest true "Music Track Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /music-track [post]
func CreateMusicTrack(c *gin.Context) {
	var requestBody models.MusicTrackRequest

	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	// userId, exists := c.Get("id")
	// if !exists {
	// 	response.Message = "cannot get user"
	// 	response.SendResponse(c)
	// 	return
	// }

	musictrack, err := services.CreateMusicTrack(primitive.NewObjectID(), &requestBody)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{"MusicTrack": musictrack}
	response.SendResponse(c)
}

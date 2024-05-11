package controllers

import (
	"music-library-management/models"
	"music-library-management/services"
	"music-library-management/utils"
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

	musictrack, err := services.CreateMusicTrack(primitive.NewObjectID(), &requestBody)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{"Music Track": musictrack}
	response.SendResponse(c)
}

// UpdateMusicTrack godoc
// @Summary      Update a music track
// @Description  update a music track by id
// @Tags         music-tracks
// @Accept       json
// @Produce      json
// @Param        id path string true "Music Track ID"
// @Param        req body models.MusicTrackRequest true "Music Track Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /music-track/{id} [put]
func UpdateMusicTrack(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	idHex := c.Param("id")
	musicTrackId, _ := primitive.ObjectIDFromHex(idHex)

	var requestBody models.MusicTrackRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	err := services.UpdateMusicTrack(musicTrackId, &requestBody)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.SendResponse(c)
}

// DeleteMusicTrack godoc
// @Summary      Delete a music track
// @Description  delete a music track by id
// @Tags         music-tracks
// @Accept       json
// @Produce      json
// @Param        id path string true "Music Track ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /music-track/{id} [delete]
func DeleteMusicTrack(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	idHex := c.Param("id")
	musicTrackId, _ := primitive.ObjectIDFromHex(idHex)

	err := services.DeleteMusicTrack(musicTrackId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.SendResponse(c)
}

// GetMusicTrackById godoc
// @Summary      Get a music track by id
// @Description  get a music track by id
// @Tags         music-tracks
// @Accept       json
// @Produce      json
// @Param        id path string true "Music Track ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /music-track/{id} [get]
func GetMusicTrackById(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	idHex := c.Param("id")
	musicTrackId, _ := primitive.ObjectIDFromHex(idHex)

	musicTrack, err := services.GetMusicTrackById(musicTrackId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"Music Track": musicTrack}
	response.SendResponse(c)
}

// GetMusicTracks godoc
// @Summary      Get all music tracks
// @Description  get all music tracks, have paging, sorting, and filtering
// @Tags         music-tracks
// @Accept       json
// @Produce      json
// @Param        page query int false "Page Number"
// @Param        size query int false "Page Size"
// @Param        sort query string false "Sort Field"
// @Param        filter query string false "Filter Field"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /music-tracks [get]
func GetMusicTracks(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	page, size, sort, filterMap, shouldReturn := utils.Pagination(c, response)
	if shouldReturn {
		return
	}

	musicTracks, err := services.GetMusicTracks(page, size, sort, filterMap)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"Music Tracks": musicTracks}
	response.SendResponse(c)
}

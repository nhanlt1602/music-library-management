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
// @Router       /music-tracks [post]
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
// @Router       /music-tracks/{id} [put]
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
	response.Message = "Music Track updated successfully"
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
// @Router       /music-tracks/{id} [delete]
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
	response.Message = "Music Track deleted successfully"
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
// @Router       /music-tracks/{id} [get]
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
// @Param 	  	 paging_ignore query bool false "Ignore Paging"
// @Param        search query string false "search Field"
// @Param        title query string false "Title"
// @Param        artist query string false "Artist"
// @Param        album query string false "Album"
// @Param        genre query string false "Genre"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /music-tracks [get]
func GetMusicTracks(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	request := buildListMusicTrackRequest(c)

	musicTracks, err := services.GetMusicTracks(request)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	hasPrev := request.Paging.Page > 0
	hasNext := len(musicTracks) > request.Paging.Size

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"items":          musicTracks,
		"page":           request.Paging.Page,
		"size":           request.Paging.Size,
		"sort":           request.Paging.Sort,
		"paging_ignore":  request.Paging.PagingIgnore,
		"hasPrev":        hasPrev,
		"hasNext":        hasNext,
		"total_elements": len(musicTracks),
	}
	response.SendResponse(c)
}

func buildListMusicTrackRequest(c *gin.Context) *models.GetMusicTrackRequest {
	pagingReq := utils.AsPageable(c)

	return &models.GetMusicTrackRequest{
		Title:  c.DefaultQuery("title", ""),
		Artist: c.DefaultQuery("artist", ""),
		Album:  c.DefaultQuery("album", ""),
		Genre:  c.DefaultQuery("genre", ""),
		Paging: *pagingReq,
	}
}

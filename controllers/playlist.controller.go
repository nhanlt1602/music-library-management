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

// CreatePlaylist godoc
// @Summary      Create Playlist
// @Description  create a new playlist
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param        req body models.PlaylistRequest true "Playlist Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /playlists [post]
func CreatePlaylist(c *gin.Context) {
	var requestBody models.PlaylistRequest

	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	playlist, err := services.CreatePlaylist(primitive.NewObjectID(), &requestBody)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusCreated
	response.Success = true
	response.Data = gin.H{"Playlist": playlist}
	response.SendResponse(c)
}

// UpdatePlaylist godoc
// @Summary      Update a playlist
// @Description  update a playlist by id
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param        id path string true "Playlist ID"
// @Param        req body models.PlaylistRequest true "Playlist Request"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /playlists/{id} [put]
func UpdatePlaylist(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	playlistId := c.Param("id")
	playlistObjectId, err := primitive.ObjectIDFromHex(playlistId)
	if err != nil {
		response.Message = "Invalid Playlist ID"
		response.SendResponse(c)
		return
	}

	var requestBody models.PlaylistRequest
	_ = c.ShouldBindBodyWith(&requestBody, binding.JSON)

	err = services.UpdatePlaylist(playlistObjectId, &requestBody)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}

// DeletePlaylist godoc
// @Summary      Delete a playlist
// @Description  delete a playlist by id
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param        id path string true "Playlist ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /playlists/{id} [delete]
func DeletePlaylist(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	playlistId := c.Param("id")
	playlistObjectId, err := primitive.ObjectIDFromHex(playlistId)
	if err != nil {
		response.Message = "Invalid Playlist ID"
		response.SendResponse(c)
		return
	}

	err = services.DeletePlaylist(playlistObjectId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.SendResponse(c)
}

// GetPlaylistById godoc
// @Summary      Get a playlist by id
// @Description  get a playlist by id
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param        id path string true "Playlist ID"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /playlists/{id} [get]
func GetPlaylistById(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	playlistId := c.Param("id")
	playlistObjectId, err := primitive.ObjectIDFromHex(playlistId)
	if err != nil {
		response.Message = "Invalid Playlist ID"
		response.SendResponse(c)
		return
	}

	playlist, err := services.GetPlaylistById(playlistObjectId)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"Playlist": playlist}
	response.SendResponse(c)
}

// GetPlaylists godoc
// @Summary      Get all playlists
// @Description  get all playlists
// @Tags         playlists
// @Accept       json
// @Produce      json
// @Param        page query int false "Page Number"
// @Param        size query int false "Page Size"
// @Param        sort query string false "Sort"
// @Param        filter query string false "Filter"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Router       /playlists [get]
func GetPlaylists(c *gin.Context) {
	response := &models.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	page, size, sort, filter, shouldReturn := utils.Pagination(c, response)
	if shouldReturn {
		return
	}

	playlists, err := services.GetPlaylists(page, size, sort, filter)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{"Playlists": playlists}
	response.SendResponse(c)
}

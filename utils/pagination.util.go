package utils

import (
	"music-library-management/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AsPageable(c *gin.Context) *models.PagingRequest {
	pageStr := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(pageStr)
	sizeStr := c.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeStr)
	sort := c.DefaultQuery("sort", "")
	pagingIgnoreStr := c.DefaultQuery("paging_ignore", "false")
	pagingIgnore, _ := strconv.ParseBool(pagingIgnoreStr)

	return &models.PagingRequest{Page: page, Size: size, Sort: sort, PagingIgnore: pagingIgnore}
}

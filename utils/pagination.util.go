package utils

import (
	"encoding/json"
	"music-library-management/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context, response *models.Response) (int, int, string, map[string]interface{}, bool) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	sizeStr := c.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeStr)
	sort := c.DefaultQuery("sort", "")
	filter := c.DefaultQuery("filter", "")

	filterMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(filter), &filterMap)
	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return 0, 0, "", nil, true
	}
	return page, size, sort, filterMap, false
}

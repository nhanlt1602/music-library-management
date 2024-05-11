package routes

import (
	"music-library-management/docs"
	"music-library-management/models"
	"music-library-management/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	r := gin.New()
	initRoute(r)

	// r.Use(gin.LoggerWithWriter(middlewares.LogWriter()))
	// r.Use(gin.CustomRecovery(middlewares.AppRecovery()))
	// r.Use(middlewares.CORSMiddleware())
	v1 := r.Group("/v1")
	{
		PingRoute(v1)
	}

	docs.SwaggerInfo.BasePath = v1.BasePath()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func initRoute(r *gin.Engine) {
	_ = r.SetTrustedProxies(nil)
	r.RedirectTrailingSlash = true
	r.HandleMethodNotAllowed = true

	r.NoRoute(func(c *gin.Context) {
		models.SendErrorResponse(c, http.StatusNotFound, c.Request.RequestURI+" not found")
	})

	r.NoMethod(func(c *gin.Context) {
		models.SendErrorResponse(c, http.StatusMethodNotAllowed, c.Request.Method+" method not allowed")
	})

}

func InitGin() {
	gin.DisableConsoleColor()
	gin.SetMode(repository.Config.Mode)
}

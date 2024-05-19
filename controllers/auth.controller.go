package controllers

import "github.com/gin-gonic/gin"

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body RegisterRequest true "Register Request"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /auth/register [post]
func Register(c *gin.Context) {
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body LoginRequest true "Login Request"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /auth/login [post]
func Login(c *gin.Context) {
}

// Refresh godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body RefreshRequest true "Refresh Request"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /auth/refresh [post]
func Refresh(c *gin.Context) {
}

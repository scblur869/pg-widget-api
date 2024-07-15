package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// api/v1/health
func HealthCheck(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "service up")
}

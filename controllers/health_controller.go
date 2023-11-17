package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (hc *HealthController) GetHealth(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

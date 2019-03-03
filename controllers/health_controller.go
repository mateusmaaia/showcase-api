package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h *HealthController) Check(ctx *gin.Context) {
	res := string("All engines running. Liftoff!")
	ctx.JSON(http.StatusOK, res)
}

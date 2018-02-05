package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func MapiWebhook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "done")
}

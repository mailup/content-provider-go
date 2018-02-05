package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func MapiHealthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Mapi test webhook is up!")
}

func MapiWebhook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "done")
}

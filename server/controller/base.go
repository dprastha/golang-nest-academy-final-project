package controller

import (
	"final-project/server/view"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponseGin(c *gin.Context, payload *view.Response) {
	c.JSON(payload.Status, payload)
}

func WriteErrorJsonResponseGin(c *gin.Context, payload *view.Response) {
	c.AbortWithStatusJSON(payload.Status, payload)
}

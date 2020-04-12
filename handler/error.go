package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestError(c *gin.Context, message string) {
	SendErrorResponse(c, http.StatusBadRequest, message)
}

func NotFoundError(c *gin.Context, message string) {
	SendErrorResponse(c, http.StatusNotFound, message)
}

func UnauthorizedError(c *gin.Context, message string) {
	SendErrorResponse(c, http.StatusUnauthorized, message)
}

func SendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
	c.Abort()
}

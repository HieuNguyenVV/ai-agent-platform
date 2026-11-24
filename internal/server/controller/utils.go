package controller

import (
	appErrors "agent/ai-agent-platform/internal/errors"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	_ = c.Error(err)

	appError, ok := err.(appErrors.AppError)
	if ok {
		c.AbortWithStatusJSON(appError.StatusCode, err)
		return
	}
	c.AbortWithStatusJSON(appErrors.ErrInternalServerError.StatusCode, err)
}
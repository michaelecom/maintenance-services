package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newResponseError(c *gin.Context, err error, code int, message string) {
	fields := logrus.Fields{
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
		"error":  err,
	}

	input, _ := c.Get("input")
	if input != nil {
		fields["input"] = input
	}

	logrus.WithFields(fields).Error(message)

	c.AbortWithStatusJSON(code, errorResponse{message})
}

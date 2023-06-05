package status

import (
	"github.com/Onelvay/HL-architecture/pkg/loggerlocal"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func NewResponse(c *gin.Context, statusCode int, message string) {
	loggerlocal.Error(message)
	c.AbortWithStatusJSON(statusCode, Response{message})
}

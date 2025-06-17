package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Responder struct {
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponderNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
func ResponderApiOk(c *gin.Context, data interface{}) {
	res := Responder{
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func ResponderApiCreated(c *gin.Context, data interface{}) {
	res := Responder{
		Data: data,
	}
	c.JSON(http.StatusCreated, res)
}

func ResponderApiError(c *gin.Context, status int, err error, msg string) {
	res := Responder{
		Error:   err.Error(),
		Message: msg,
	}
	c.JSON(status, res)
}

package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "ping")
}

package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transportHTTP struct {
	server *gin.Engine
}

// RegisterHTTPEndpoints registers the HTTP endpoints for the ping feature
func RegisterHTTPEndpoints(server *gin.Engine) {
	transport := &transportHTTP{
		server: server,
	}

	server.GET("/ping", transport.pingHandler)
}

func (t *transportHTTP) pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

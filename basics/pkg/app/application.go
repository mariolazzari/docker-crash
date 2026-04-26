package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/leon-devs/course-docker_intro/basics/pkg/config"
	"github.com/leon-devs/course-docker_intro/basics/pkg/features/ping"
)

type Application interface {
	Start(ctx context.Context) error
}

type app struct {
	conf       *config.Config
	httpServer *gin.Engine
}

// NewApplication returns a new instance of Application
func NewApplication(ctx context.Context) (Application, error) {
	conf, err := config.ParseConfig(ctx)
	if err != nil {
		return nil, err
	}

	// Create the HTTP server:
	httpServer := gin.Default()

	// Register the 'ping' feature:
	ping.RegisterHTTPEndpoints(httpServer)

	// TODO: You can create new features and then register their HTTP / gRPC handlers in here.

	// Configure and return the final application:
	application := &app{
		conf:       conf,
		httpServer: httpServer,
	}

	return application, nil
}

// Start attempts to start the given instance of Application
func (a *app) Start(_ context.Context) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", a.conf.HTTP.Port), a.httpServer)
}

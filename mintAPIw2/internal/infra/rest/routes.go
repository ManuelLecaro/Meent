package rest

import (
	"fmt"
	"mintapi/internal/infra/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes loads the handlers for the urls on the application
func (a *Server) SetupRoutes() {
	baseGroup := a.engine.Group(fmt.Sprintf("/api/%s", config.GetNameConfigs()))

	go setupHealthCheck(baseGroup)

	a.setupIntegrationRoutes(baseGroup)
}

func setupHealthCheck(g *gin.RouterGroup) {
	g.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

func (a *Server) setupIntegrationRoutes(g *gin.RouterGroup) {
	integrationGroup := g.Group("/integration")

	a.setupIntegrationEventRoutes(integrationGroup)
}

func (a *Server) setupIntegrationEventRoutes(g *gin.RouterGroup) {
	g.POST("/event", a.Handler.Event.HandleCreate)
	g.GET("/event/:id", a.Handler.Event.HandleGet)
	g.GET("/event/:id/mint", nil)
}

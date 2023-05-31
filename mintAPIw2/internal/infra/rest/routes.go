package rest

import (
	"fmt"
	"mintapi/internal/infra/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes loads the handlers for the urls on the application
func (a *Server) SetupRoutes() {
	baseGroup := a.engine.Group(fmt.Sprintf("/api/%s", a.Config.GetString(config.NameConfig)))

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

	a.setupIntegrationTicketRoutes(integrationGroup)
	a.setupIntegrationEventRoutes(integrationGroup)
}

func (a *Server) setupIntegrationTicketRoutes(g *gin.RouterGroup) {
	g.POST("/event", nil)
	g.GET("/event", nil)
}

func (a *Server) setupIntegrationEventRoutes(g *gin.RouterGroup) {
	g.POST("/ticket", nil)
	g.GET("/ticket/:id", nil)
}

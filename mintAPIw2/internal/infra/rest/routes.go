package rest

import "github.com/gin-gonic/gin"

// SetupRoutes loads the handlers for the urls on the application
func (a *Server) SetupRoutes(g *gin.RouterGroup) {
	a.setupIntegrationRoutes(g)
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

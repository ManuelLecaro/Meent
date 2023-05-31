package rest

import (
	"fmt"
	"log"
	"mintapi/internal/infra/config"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	engine       *gin.Engine
	StatusReady  chan bool
	StatusKilled chan bool
	Config       *viper.Viper
}

func NewServer(appName string) *Server {
	gin.SetMode(gin.ReleaseMode)

	server := gin.New()

	server.Use(
		requestid.New(),
		gin.Recovery(),
		gin.LoggerWithConfig(GetLoggerConfig(appName)),
	)

	newSrv := &Server{
		engine:       server,
		StatusReady:  make(chan bool),
		StatusKilled: make(chan bool),
		Config:       config.LoadConfiguration(),
	}

	newSrv.initInfrastructure()

	return newSrv
}

// GetLoggerConfig defines gin logs format and skip paths
func GetLoggerConfig(appName string) gin.LoggerConfig {
	var formatter = func(param gin.LogFormatterParams) string {
		if param.Latency > time.Minute {
			param.Latency = param.Latency - param.Latency%time.Second
		}

		return fmt.Sprintf(
			"[time:%s][method:%s][uri:%s][status:%d][latency:%s][client_ip:%s][err_message:%s]",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.Method,
			param.Request.URL,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.ErrorMessage,
		)
	}

	return gin.LoggerConfig{
		Formatter: formatter,
		SkipPaths: []string{
			fmt.Sprintf("/api/%s/health-check", appName),
		},
	}
}

func (a *Server) initInfrastructure() {
	configurations := config.LoadConfiguration()
	a.Config = configurations
}

func (a *Server) runServer() {
	log.Println("running app server at port: ", a.Config.GetString(config.PortConfig))
	if err := a.engine.Run(fmt.Sprintf(":%s", a.Config.GetString(config.PortConfig))); err != nil {
		log.Println("Shutting down the server")
	}
}

func (a *Server) KillApp() {
	if a.StatusKilled != nil {
		a.StatusKilled <- true
	}
}

// RunApp initializes the server and consumers
func (a *Server) RunApp() {
	shutDown := make(chan bool)
	go a.runServer()
	if a.StatusReady != nil {
		a.StatusReady <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	close(shutDown)
	a.KillApp()
}

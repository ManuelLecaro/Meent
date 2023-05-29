package rest

import (
	"fmt"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(appName string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.New()

	server.Use(
		requestid.New(),
		gin.Recovery(),
		gin.LoggerWithConfig(GetLoggerConfig(appName)),
	)

	return server
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

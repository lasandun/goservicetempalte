package main

import (
	"log/slog"
	"strconv"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"test.com/microservice/config"
	"test.com/microservice/internal/logging"
	"test.com/microservice/internal/middleware"
	"test.com/microservice/routes"
)

func main() {

	// Load service configs
	config.LoadConfig()

	// Initialize JSON logger (stdout)
	logging.InitJSONLogger(slog.LevelInfo)

	r := gin.New()

	// Panic recovery middleware first
	r.Use(gin.Recovery())

	// Request ID (sets "req_id" into context and response header X-Request-Id)
	r.Use(requestid.New())

	// Replace gin.Logger() with our custom access log
	skip := map[string]struct{}{
		"/healthz": {},
		"/metrics": {},
	}
	r.Use(middleware.AccessLog(skip))

	// Register routes
	routes.RegisterRoutes(r)

	// Start server
	err := r.Run(config.Cfg.Server.Host + ":" + strconv.Itoa(config.Cfg.Server.Port))

	// Start server
	if err != nil {
		slog.Error("server_start_failed", "err", err)
	}
}

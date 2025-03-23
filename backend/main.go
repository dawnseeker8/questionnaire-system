package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dawnseeker8/questionnaire-system/app/model"
	"github.com/dawnseeker8/questionnaire-system/config"
	_ "github.com/dawnseeker8/questionnaire-system/docs" // Import swagger docs
	"github.com/dawnseeker8/questionnaire-system/router"
	hertzlogger "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// Initialize logger
	logger := hertzlogger.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)

	// Initialize database
	err = model.InitDB(&cfg.Database)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize database: %v", err))
	}
	hlog.Info("Database connection established successfully")

	// Create Hertz server
	h := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%d", cfg.Server.Port)),
	)

	// Setup Swagger
	// Note: docs are imported in the import section with _ "github.com/dawnseeker8/questionnaire-system/docs"
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	hlog.Infof("Swagger UI available at: http://localhost:%d/swagger/index.html", cfg.Server.Port)

	// Register routes
	router.RegisterRoutes(h)

	// Start server
	hlog.Infof("Starting server on port %d", cfg.Server.Port)
	h.Spin()
}

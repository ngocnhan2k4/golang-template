package main

import (
	"Template/configs"
	_ "Template/docs"
	"Template/internal/faculty"
	"Template/internal/status"
	"Template/internal/student"
	"Template/pkg/accesslog"
	"Template/pkg/dbcontext"
	"Template/pkg/log"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

// @title			Student Management API
// @version		1.0
// @description	This is a server for a student management system.
// @host			localhost:8081
// @BasePath		/api/v1
// @schemes		http
func main() {
	configs.LoadConfig()
	db, err := gorm.Open(postgres.Open(configs.AppConfig.DSN), &gorm.Config{})
	logger := log.New().With(context.TODO(), "version", Version)
	if err != nil {
		panic("failed to connect database")
	}

	dbContext := dbcontext.New(db)
	addr := fmt.Sprintf(":%d", configs.AppConfig.Server.Port)

	//go routing.GracefullyShutdown(addr)

	handler := buildHandler(dbContext, logger)
	if err := http.ListenAndServe(addr, handler); err != nil && err != http.ErrServerClosed {
		logger.Errorf("failed to start server: %v", err)
		os.Exit(1)
	}

}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(db *dbcontext.DB, logger log.Logger) http.Handler {
	router := gin.New()

	router.Use(accesslog.Handler(logger))
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	v1 := router.Group("/api")

	studentGroup := v1.Group("/students")
	facultyGroup := v1.Group("/Faculty")
	statusGroup := v1.Group("/StudentStatus")
	student.RegisterHandlers(studentGroup, student.NewService(student.NewRepository(db)))
	faculty.RegisterHandlers(facultyGroup, faculty.NewService(faculty.NewRepository(db), logger), logger)
	status.RegisterHandlers(statusGroup, status.NewService(status.NewRepository(db), logger), logger)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

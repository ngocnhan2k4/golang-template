package main

import (
	"Template/configs"
	"Template/internal/student"
	"Template/pkg/dbcontext"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configs.LoadConfig()
	db, err := gorm.Open(postgres.Open(configs.AppConfig.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	dbContext := dbcontext.New(db)
	addr := fmt.Sprintf(":%d", configs.AppConfig.Server.Port)

	//go routing.GracefullyShutdown(addr)

	handler := buildHandler(dbContext)
	if err := http.ListenAndServe(addr, handler); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
		os.Exit(1)
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(db *dbcontext.DB) http.Handler {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")

	studentGroup := v1.Group("/students")
	studentService := student.NewRepository(db)
	student.RegisterHandlers(studentGroup, studentService)

	return router
}

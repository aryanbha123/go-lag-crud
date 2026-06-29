package main

import (
	"crud-go/internal/config"
	"crud-go/internal/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Cortexone internal assessment initialise ....")

	fmt.Println("Database connecting....")
	cfg := config.Load()
	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	// CORS: allow a separate frontend (e.g. React on :3000) to call this API.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// // Serve the UI: ./ui is exposed at /ui, and "/" redirects to the index.
	// router.Static("/ui", "./ui")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui/index.html")
	})

	api := router.Group("/api/v1")

	routes.RegisterRoutes(api, db)

	router.Run(":8080")

}

package main

import (
	"context"
	"fmt"

	handler "github.com/scblur869/pg-widget-api/services"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//setting up a 32 bit key if we want to later protect an endpoint
	//  api key security its simple and could also be used as a seed for a bearer token / JWT

	key := handler.SetEncryptionKeyEnv()
	handler.CreateKeyFile("key.txt", key)
	os.Setenv("KEY", key)
	fmt.Println("key :", os.Getenv("KEY"))
}

func main() {

	// allowedHost := os.Getenv("ALLOWED")
	appAddr := "0.0.0.0:" + os.Getenv("PORT")

	// setting up the handler for reciever functions and connects to the database
	db := new(handler.PgHandler)
	db.Ctx = context.Background()
	db.Config = handler.ParseConfiguration()
	conn, err := db.Config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Connect = conn // sets up the db connection for the handler

	// gin.SetMode(gin.ReleaseMode)
	// routes and CORS configuration
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Accept", "X-Requested-With", "Content-Type", "Authorization", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/api/v1/health", handler.HealthCheck)                                 // tested
	r.GET("/api/v1/getAll", db.GetAllWidgets)                                    // tested
	r.GET("/api/v1/getByColor/:color", db.GetWidgetsByColor)                     // tested
	r.GET("/api/v1/getByCategory/:category", db.GetWidgetsByCategory)            // tested
	r.POST("/api/v1/addNew", db.AddNewWidget)                                    // tested
	r.PUT("/api/v1/updateWidgetColor/:id/:color", db.UpdateWidgetColor)          // tested
	r.PUT("/api/v1/updateWidgetCategory/:id/:category", db.UpdateWidgetCategory) // tested
	r.DELETE("/api/v1/deleteById/:id", db.DeleteWidget)                          // tested
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Not Found"})
	})

	// http server config
	srv := &http.Server{
		Addr:    appAddr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdowning Server and closing database connections ...")
	defer db.Connect.Close(db.Ctx) // closes database handler
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

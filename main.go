package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gofhir/db"
	"github.com/gofhir/generated"
	"github.com/joho/godotenv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//go:generate go run ./generator

func main() {
	binding.EnableDecoderDisallowUnknownFields = true

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	db.SetupDatabase()
	defer db.DisconnectDatabase()

	r := gin.Default()

	generated.Routes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    "127.0.0.1:" + port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	done := <-ctx.Done()
	log.Println("Server exiting with code:", done)

	log.Println("Server exiting")
}

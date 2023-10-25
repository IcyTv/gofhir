package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gofhir/db"
	"github.com/gofhir/db/versioning"
	"github.com/gofhir/generated"
	"github.com/joho/godotenv"
)

//go:generate go run ./generator

func main() {
	binding.EnableDecoderDisallowUnknownFields = true

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	db.SetupDatabase()
	defer db.DisconnectDatabase()

	// Setup collections
	generated.CreateCollections(db.FhirDatabase)

	var waitGroup sync.WaitGroup

	cancelFn, err := versioning.SetupChangeStream(&waitGroup, db.FhirDatabase)
	if err != nil {
		log.Fatal(err)
	}

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
	log.Println("Server exiting with code:", done, " Waiting for change stream to close...")

	cancelFn()
	waitGroup.Wait()

	log.Println("Server exiting")

}

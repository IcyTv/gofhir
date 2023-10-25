package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var FhirDatabase *mongo.Database
var VersioningDatabase *mongo.Database

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SetupDatabase() {
	fmt.Println("Setting up database")
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set the 'MONGODB_URI' environment variable")
	}
	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	check(err)

	FhirDatabase = Client.Database("fhir")
	VersioningDatabase = Client.Database("fhir-delta")
}

func DisconnectDatabase() {
	check(Client.Disconnect(context.TODO()))
}

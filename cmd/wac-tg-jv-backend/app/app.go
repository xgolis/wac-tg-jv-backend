package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Server *http.Server
}

var (
	ClientDB *mongo.Client
	DB       *mongo.Database
)

func NewApp() *App {
	mux := MakeHandlers()

	uri := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	// defer client.Disconnect(ctx)

	DB = client.Database("wac-tg-jv")

	return &App{
		Server: &http.Server{
			Addr:           "0.0.0.0:8080",
			Handler:        mux,
			ReadTimeout:    100 * time.Second,
			WriteTimeout:   100 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

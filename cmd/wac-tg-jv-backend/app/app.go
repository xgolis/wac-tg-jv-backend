package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Server *http.Server
	Config Config
}

type Config struct {
	MongoDBConn string
	Server      string
	Port        string
}

var (
	ClientDB *mongo.Client
	DB       *mongo.Database
)

func NewApp() *App {
	mux := MakeHandlers()

	conf, err := getConfig()
	if err != nil {
		panic(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.MongoDBConn))
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
			Addr:           conf.Server + ":" + conf.Port,
			Handler:        mux,
			ReadTimeout:    100 * time.Second,
			WriteTimeout:   100 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func getConfig() (*Config, error) {
	mongoURL := os.Getenv("MongoDBURL")
	if mongoURL == "" {
		return nil, fmt.Errorf("no URL provided for MongoDB")
	}
	mongoPort := os.Getenv("MongoDBPort")
	if mongoPort == "" {
		return nil, fmt.Errorf("no Port provided for MongoDB")
	}
	mongoUser := os.Getenv("MongoDBUsername")
	if mongoUser == "" {
		return nil, fmt.Errorf("no User provided for MongoDB")
	}
	mongoPassword := os.Getenv("MongoDBPassword")
	if mongoPassword == "" {
		return nil, fmt.Errorf("no Password provided for MongoDB")
	}
	serverURL := os.Getenv("serverURL")
	if serverURL == "" {
		return nil, fmt.Errorf("no User provided for MongoDB")
	}
	serverPort := os.Getenv("serverPort")
	if serverPort == "" {
		return nil, fmt.Errorf("no Password provided for MongoDB")
	}

	// mongodb://username:password@host:port/database
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/wac-tg-jv", mongoUser, mongoPassword, mongoURL, mongoPort)
	return &Config{
		MongoDBConn: url,
		Server:      serverURL,
		Port:        serverPort,
	}, nil
}

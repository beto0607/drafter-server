package core

import (
	"context"
	"log"

	"drafter/src/config"
	project_slice "drafter/src/slices/project"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_NAME = "drafter"

var DBClient *mongo.Client
var DBHandler *mongo.Database

func ConnectToDB() {
	uri, err := config.GetEnvVar("MONGODB_URI")

	if err != nil {
		log.Panic("Couldn't get DB connection string")
	}

	client, err := mongo.Connect(context.Background(), options.Client().
		ApplyURI(uri))

	if err != nil {
		log.Panic("Couldn't create DB client")
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	DBClient = client

	dbHandler := client.Database(DB_NAME)
	DBHandler = dbHandler
	log.Println("DB connected :D")
	initSlices()
}

func initSlices() {
	project_slice.InitSliceDB(DBHandler)
}

func DisconnectDB() {
	if err := DBClient.Disconnect(context.Background()); err != nil {
		log.Panic("Couldn't disconnectDB client")
	}
	log.Println("DB disconnected D:")
}

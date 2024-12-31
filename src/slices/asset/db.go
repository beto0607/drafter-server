package asset_slice

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database
var coll *mongo.Collection

func initAssetSliceDB(handler *mongo.Database) {
	db = handler
	err := handler.CreateCollection(context.TODO(), "assets")
	if err != nil {
		log.Panicln("Couldn't create collection 'assets'")
	}
	log.Println("Collection 'assets' created")
	coll = handler.Collection("assets")
}

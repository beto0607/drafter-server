package project_slice

import (
	"context"
	"drafter/src/utils"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database
var coll *mongo.Collection

func InitSliceDB(handler *mongo.Database) {
	db = handler
	err := handler.CreateCollection(context.TODO(), "projects")
	if err != nil {
		log.Panicln("Couldn't create collection 'projects'")
	}
	log.Println("Collection 'projects' created")
	coll = handler.Collection("projects")
}

func SaveProject(project ProjectModel) (*ProjectModel, error) {
	project.CreatedAt = utils.UTCTimestamp()
	project.Id = primitive.NewObjectID()
	result, err := coll.InsertOne(context.TODO(), project)
	if err != nil {
		return nil, err
	}
	if _, ok := result.InsertedID.(primitive.ObjectID); ok {
		return &project, nil
	}
	return nil, errors.New("Couldn't retrieve InsertedID")
}

func FindProject(projectId string) (*ProjectModel, error) {
	objectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	result := coll.FindOne(context.TODO(), filter)

	if err = result.Err(); err != nil {
		return nil, err
	}
	var project ProjectModel
	result.Decode(&project)
	return &project, nil
}

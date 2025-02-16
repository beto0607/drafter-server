package project_slice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectDto struct {
	Data ProjectModel `json:"data"`
}

type ProjectModel struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	Elements        []ElementModel     `json:"elements" bson:"elements"`
	BackgroundColor string             `json:"backgroundColor" bson:"backgroundColor"`
	Tags            []TagModel         `json:"tags" bson:"tags"`
	Name            string             `json:"name" bson:"name"`
	CreatedAt       string             `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt       string             `json:"updatedAt,omitempty" bson:"updatedAt"`
	DeletedAt       string             `json:"deletedAt,omitempty" bson:"deletedAt"`
}

type ElementModel struct {
	Id       string        `json:"id" bson:"id"`
	Name     string        `json:"name" bson:"name"`
	Title    string        `json:"title" bson:"title"`
	Locked   bool          `json:"locked" bson:"locked"`
	Caption  string        `json:"caption" bson:"caption"`
	Size     SizeModel     `json:"size" bson:"size"`
	Position PositionModel `json:"position" bson:"position"`
	Rotation RotationModel `json:"rotation" bson:"rotation"`
	Tags     []TagModel    `json:"tags" bson:"tags"`
	Assets   []AssetModel  `json:"assets" bson:"assets"`
}

type AssetModel struct {
	Id         string `json:"id" bson:"id"`
	Url        string `json:"url" bson:"url"`
	Type       string `json:"type" bson:"type"`
	Extenstion string `json:"extension" bson:"extension"`
	MimeType   string `json:"mimeType" bson:"mimeType"`
}

type TagModel struct {
	Color string `json:"color" bson:"color"`
	Text  string `json:"text" bson:"text"`
}

type SizeModel struct {
	Width  int32 `json:"width" bson:"width"`
	Height int32 `json:"height" bson:"height"`
}

type RotationModel struct {
	X float32 `json:"x" bson:"x"`
	Y float32 `json:"y" bson:"y"`
	Z float32 `json:"z" bson:"z"`
}

type PositionModel struct {
	X int32 `json:"x" bson:"x"`
	Y int32 `json:"y" bson:"y"`
}

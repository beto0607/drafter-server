package asset_slice

import "go.mongodb.org/mongo-driver/bson/primitive"

type AssetModel struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt string             `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt string             `json:"updatedAt,omitempty" bson:"updatedAt"`
	DeletedAt string             `json:"deletedAt,omitempty" bson:"deletedAt"`
	BlobId    string             `json:"blobId" bson:"blobId"`
	Url       string             `json:"url"  bson:"url"`
	Type      string             `json:"type" bson:"type"`
	MimeType  string             `json:"mimeType" bson:"mimeType"`
	Extension string             `json:"extension" bson:"extension"`
}

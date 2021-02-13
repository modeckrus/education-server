package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Image image with crops
type Image struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UID      string             `json:"uid"`
	Name     string             `json:"name"`
	URL      string             `json:"url"`
	BlurHash string             `json:"blurHash"`
	Crops    []*Crop            `json:"crops"`
}

//Crop ...
type Crop struct {
	URL  string `json:"url"`
	Size int    `json:"size"`
}

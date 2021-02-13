package database

import (
	"education/model"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func defaultPhoto() *model.Image {
	id, err := primitive.ObjectIDFromHex("601a90e9c4ca61d59f16ad4a")
	if err != nil {
		log.Println("Error while generate default photo")
	}
	return &model.Image{
		ID:       id,
		UID:      "0OVoJ17pNMOvMQFaxmQsory2Fek2",
		Name:     "2.jpg",
		URL:      "https://firebasestorage.googleapis.com/v0/b/education-modeck.appspot.com/o/0OVoJ17pNMOvMQFaxmQsory2Fek2/2.jpg?alt=media&token=64f9699d-9145-4c59-8890-255abd812a39",
		BlurHash: "LuLXY.?b~qj[xuWBj[Rj?bM{M{Rj",
		Crops: []*model.Crop{
			{
				URL:  "https://firebasestorage.googleapis.com/v0/b/education-modeck.appspot.com/o/0OVoJ17pNMOvMQFaxmQsory2Fek2/100_2.jpg?alt=media&token=badaf427-4f0b-40d0-85c7-206f1fa95901",
				Size: 100,
			},
		},
	}
}

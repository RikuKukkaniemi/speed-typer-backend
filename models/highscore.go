package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Highscore struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  *string            `json:"name"`
	Score *int64             `json:"score"`
	Date primitive.DateTime  `json:"date"`
}

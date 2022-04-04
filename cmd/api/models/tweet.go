package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tweet struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullText string             `json:"full_text,omitempty" bson:"full_text,omitempty"`
	User     struct {
		ScreenName []string `json:"screen_name" bson:"screen_name"`
	} `json:"user,omitempty" bson:"user,omitempty"`
}

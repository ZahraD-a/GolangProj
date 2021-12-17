package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tasks struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Titre      string             `json:"titre,omitempty"`
	DateDebut  time.Time          `json:"DateDebut,omitempty"`
	Estimation string             `json:"estimation,omitempty"`
	Status     string             `json:"status,omitempty"`
}

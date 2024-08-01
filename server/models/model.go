package models


import "go.mongodb.org/mongo-driver/bson/primitive"


type ToDoList struct {
	ID        `json:"_id,omitempty"   bson:"_id,omitempty"`
	Task		`json:"task,omitempty"`
	Status		`json:"status,omitempty"`
}


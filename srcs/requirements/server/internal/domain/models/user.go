package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const UserCollectionName = "users"

type Friends struct {
	Id   int32  `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

type UserHashedPassword struct {
	Content string `json:"password,omitempty" bson:"password,omitempty"`
}

type UserFileData struct {
	Data string `json:"data,omitempty" bson:"data,omitempty"`
}

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId     string             `json:"id,omitempty" bson:"id,omitempty"`
	Password   string             `json:"password,omitempty" bson:"password,omitempty"`
	IsActive   bool               `json:"isActive,omitempty" bson:"isActive,omitempty"`
	Balance    string             `json:"balance,omitempty" bson:"balance,omitempty"`
	Age        int32              `json:"age,omitempty" bson:"age,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Gender     string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Compagny   string             `json:"company,omitempty" bson:"company,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	About      string             `json:"about,omitempty" bson:"about,omitempty"`
	Registered string             `json:"registered,omitempty" bson:"registered,omitempty"`
	Latitude   float64            `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude  float64            `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Tags       []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Friends    []Friends          `json:"friends,omitempty" bson:"friends,omitempty"`
	Data       string             `json:"data,omitempty" bson:"data,omitempty"`
}

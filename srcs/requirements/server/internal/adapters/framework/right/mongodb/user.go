package mongodb

import (
	"context"
	"data_impact/srcs/requirements/server/internal/domain/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserQuery interface {
	// CREATE
	CreateUser(user *models.User) error
	// GET
	GetUsersList() ([]*models.User, error)
	GetUser(user_id string) (*models.User, error)
	GetUserHashedPassword(user_id string) (*string, error)
	// UPDATE
	UpdateUser(user_id string, user *models.User) error
	// DELETE
	DeleteUser(user_id string) error
}

type userQuery struct{}

var userCollection *mongo.Collection

func (u *userQuery) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{
		{"id", user.UserId},
		{"password", user.Password},
		{"isActive", false},
		{"balance", user.Balance},
		{"age", user.Age},
		{"name", user.Name},
		{"gender", user.Gender},
		{"company", user.Compagny},
		{"email", user.Email},
		{"phone", user.Phone},
		{"address", user.Address},
		{"about", user.About},
		{"registered", user.Registered},
		{"latitude", user.Latitude},
		{"longitude", user.Longitude},
		{"tags", user.Tags},
		{"friends", user.Friends},
		{"data", user.Data},
	}

	insertResult, err := userCollection.InsertOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("user cannot be created")
	}
	fmt.Sprintln(insertResult)

	return nil
}

func (u *userQuery) GetUsersList() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	filter := bson.M{}

	opts := options.Find()

	// We don't need _id and password, so we exclude them
	opts.SetProjection(bson.M{
		"_id":      0,
		"password": 0,
	})

	var users []*models.User

	cursor, err := userCollection.Find(ctx, filter, opts) // add opts here
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var u models.User

		err := cursor.Decode(&u)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(users) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return users, nil
}

func (u *userQuery) GetUserHashedPassword(user_id string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"id": user_id,
	}

	opts := options.FindOne()

	// We need only user hashed password, so we only include him
	opts.SetProjection(bson.M{
		"password": 1,
	})

	var hashedPassword models.UserHashedPassword

	err := userCollection.FindOne(ctx, filter, opts).Decode(&hashedPassword)
	if err != nil {
		return nil, err
	}

	return &hashedPassword.Content, nil
}

func (u *userQuery) GetUser(user_id string) (*models.User, error) {
	log.Println("GetUser function has been launched")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"id": user_id,
	}

	opts := options.FindOne()

	// we don't need _id and password, so we exclude them
	opts.SetProjection(bson.M{
		"_id":      0,
		"password": 0,
	})

	var user models.User

	err := userCollection.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		// user does not exist
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userQuery) UpdateUser(user_id string, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"id": user_id,
	}

	// i cant find the way to change field conditionnaly
	updated := bson.D{
		{"$set", bson.D{
			{"id", user.UserId},
			{"password", user.Password},
			{"isActive", user.IsActive},
			{"balance", user.Balance},
			{"age", user.Age},
			{"name", user.Name},
			{"gender", user.Gender},
			{"company", user.Compagny},
			{"email", user.Email},
			{"phone", user.Phone},
			{"address", user.Address},
			{"about", user.About},
			{"registered", user.Registered},
			{"latitude", user.Latitude},
			{"longitude", user.Longitude},
			{"tags", user.Tags},
			{"friends", user.Friends},
			{"data", user.Data},
		}}}

	updateResult, err := userCollection.UpdateOne(ctx, filter, updated)
	if err != nil || updateResult == nil {
		return fmt.Errorf("user not found : cannot delete him")
	}

	return nil
}

func (u *userQuery) DeleteUser(user_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"id": user_id,
	}

	err := userCollection.FindOneAndDelete(ctx, filter)

	return err.Err()
}

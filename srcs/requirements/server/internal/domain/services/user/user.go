package user

import (
	"data_impact/srcs/requirements/server/internal/adapters/framework/right/mongodb"
	"data_impact/srcs/requirements/server/internal/domain/models"
)

type User interface {
	AddUsers(file []byte) error
	DeleteUser(user_id string) error
	UpdateUser(user_id string, user_updated *models.User) error
	GetUser(user_id string) (*models.User, error)
	GetUsersList() ([]*models.User, error)
}

type user struct {
	repo mongodb.Repository
}

func NewServiceUser(repo mongodb.Repository) User {
	return &user{repo: repo}
}

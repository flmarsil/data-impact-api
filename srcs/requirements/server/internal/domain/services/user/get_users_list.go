package user

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"log"
)

func (u *user) GetUsersList() ([]*models.User, error) {
	log.Println("GetUsersList function has been launched")

	users, err := u.repo.UserQuery().GetUsersList()
	if err != nil {
		return nil, err
	}

	return users, nil
}

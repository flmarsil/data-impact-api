package user

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"log"
)

func (u *user) GetUser(user_id string) (*models.User, error) {
	log.Println("GetUser function has been launched")

	user, err := u.repo.UserQuery().GetUser(user_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

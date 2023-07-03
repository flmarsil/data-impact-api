package user

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"log"
)

func (u *user) updateUserDataFile(updated_user, current_user *models.User) error {
	log.Println("updateUserDataFile function has been launched")

	err := u.DeleteUserDataFile(current_user.UserId)
	if err != nil {
		return err
	}

	err = u.CreateUserDataFile(updated_user.UserId, updated_user.Data)
	if err != nil {
		return err
	}

	return nil
}

/*
	Not being specified in the subject, I assumed that the modification of a user will be
	done with a complete json object. With all fields filled in at the same time.
*/

func (u *user) UpdateUser(user_id string, updated_user *models.User) error {
	log.Println("UpdateUser function has been launched")

	var err error

	// hash new password before update in db
	updated_user.Password, err = u.PasswordHasher([]byte(updated_user.Password))
	if err != nil {
		return err
	}

	// get user document before updating with id of the request
	current_user, err := u.repo.UserQuery().GetUser(user_id)
	if err != nil {
		return err
	}

	// process to change file data user if id or data change
	if current_user.UserId != updated_user.UserId || current_user.Data != updated_user.Data {
		err = u.updateUserDataFile(updated_user, current_user)
		if err != nil {
			return err
		}
	}

	// change document user in db
	err = u.repo.UserQuery().UpdateUser(user_id, updated_user)
	if err != nil {
		return err
	}

	return nil
}

package user

import (
	"fmt"
	"log"
	"os"
	"path"
)

func (u *user) DeleteUserDataFile(user_id string) error {
	dir := "/go/bin/files"

	nameFile := path.Join(dir, user_id+".json")

	err := os.Remove(nameFile)
	if err != nil {
		return fmt.Errorf("delete user file function has encountered a problem")
	}

	return nil
}

func (u *user) DeleteUser(user_id string) error {
	log.Println("DeleteUser function has been launched")

	// delete document in mongodb
	err := u.repo.UserQuery().DeleteUser(user_id)
	if err != nil {
		return err
	}

	// delete data file generated separatly
	err = u.DeleteUserDataFile(user_id)
	if err != nil {
		return err
	}

	return nil
}

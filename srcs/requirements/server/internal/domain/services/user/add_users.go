package user

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// hash the plaintext password of a user form the data set file and compares it to verify the success of the hashing
// returns the newly generated hash as string for inserting it in database
func (u *user) PasswordHasher(password []byte) (string, error) {
	log.Println("password hasher has been launched")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return "", err
	}
	log.Println("password user has been hashed successfully")
	return string(hash), nil
}

// generates one file per user with its id, and containing only the "data" field
func (u *user) CreateUserDataFile(user_id, user_data string) error {
	log.Println("create user data file")

	// create directory if not exist for put user file in
	dir := "/go/bin/files"
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error when creating dir for store file data user %v : %v", user_id, err)
		}
	}

	// fill content of file with user data received in dataset
	contentFile := models.UserFileData{
		Data: user_data,
	}

	// serialize file with content of data user
	file, err := json.MarshalIndent(contentFile, "", " ")
	if err != nil {
		return fmt.Errorf("error when marshaling json file for user %v : %v", user_id, err)
	}

	// put user id in file name + path of directory storage
	nameFile := path.Join(dir, user_id+".json")

	// write file
	err = ioutil.WriteFile(nameFile, file, 0666) // read & write perms
	if err != nil {
		return fmt.Errorf("error when writing json file for user %v : %v", user_id, err)
	}

	log.Println("user data file has been created successfully")
	return nil
}

// launche a goroutine for each user of the.json file
func (u *user) concurrencyHandler(users []*models.User) <-chan error {
	var wg sync.WaitGroup

	errs := make(chan error, 1)

	// iterates through all users
	for _, user := range users {
		wg.Add(1)

		// launch go routine for each user
		go func(user *models.User) {
			defer wg.Done()

			// verifying that the user does not already exist
			exist, err := u.repo.UserQuery().GetUser(user.UserId)
			if exist == nil && err == nil {
				log.Println("adding user to the database")

				// change password with hashed password for security
				user.Password, err = u.PasswordHasher([]byte(user.Password))
				if err != nil {
					errs <- fmt.Errorf("password hasher has encountered a problem")
				}

				// create one file per user with his id and data
				err := u.CreateUserDataFile(user.UserId, user.Data)
				if err != nil {
					errs <- fmt.Errorf("user data file creator has encountered a problem")
				}

				// create user document in mongodb
				err = u.repo.UserQuery().CreateUser(user)
				if err != nil {
					errs <- fmt.Errorf("document user creator in database has encountered a problem")
				}

			} else if exist != nil && err == nil {
				log.Println("user already exists in the database, do nothing")
			}
		}(user)
	}
	wg.Wait()
	close(errs)

	return errs
}

// add users to the database based on a data set file passed as a parameter
func (u *user) AddUsers(file []byte) error {
	log.Println("AddUsers function has been launched")

	var users []*models.User

	// deserialize file data
	json.Unmarshal(file, &users)

	// try to add users in database
	add := u.concurrencyHandler(users)

	if err := <-add; err != nil {
		return err
	}

	return nil
}

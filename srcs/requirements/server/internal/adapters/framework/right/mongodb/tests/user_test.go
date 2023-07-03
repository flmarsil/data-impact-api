package tests

import (
	"fmt"
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	user_id := "lalasa"
	user, err := testRepo.UserQuery().GetUser(user_id)
	if err != nil {
		log.Printf("mongodb: Test GetUser has been failed : %v\n", err)
	}

	log.Println(user)
}

func TestGetUsersList(t *testing.T) {
	users, err := testRepo.UserQuery().GetUsersList()
	if err != nil {
		log.Printf("mongodb: Test GetUsersList has been failed : %v\n", err)
	}

	for _, user := range users {
		fmt.Println(*user)
	}
}

func TestGetUserHashedPassword(t *testing.T) {
	user_id := "1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd"
	hashedPassword, err := testRepo.UserQuery().GetUserHashedPassword(user_id)
	if err != nil {
		log.Printf("mongodb: Test GetUserHashedPassword has been failed : %v\n", err)
	}

	log.Println(*hashedPassword)
}

func TestDeleteUser(t *testing.T) {
	user_id := "buhnVfyzGfe4qfuRsqwjGhzoF8ubXcM4JXqmLRmvPmJdLXtITBY99SI0rER0bX0INj2XkAyMkLUEkQrxNG46PDXHt00jGURsnjHy"
	err := testRepo.UserQuery().DeleteUser(user_id)
	if err != nil {
		log.Printf("mongodb: Test DeleteUser has been failed : %v\n", err)
	}

	log.Printf("mongodb: Test DeleteUser : user(%v) has been deleted", user_id)
}

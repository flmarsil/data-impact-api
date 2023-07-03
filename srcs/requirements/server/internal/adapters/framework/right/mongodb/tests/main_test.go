package tests

import (
	"data_impact/srcs/requirements/server/internal/adapters/framework/right/mongodb"
	"fmt"
	"log"
	"os"
	"testing"
)

const (
	dbURL = "mongodb://flmarsil:845581EF64F3D5F1C0F17EE7D22524B165A7160FB28828B8420232D8F11D2A40@db:27017/"
)

var testRepo mongodb.Repository

func TestMain(m *testing.M) {
	db, err := mongodb.NewDB(dbURL)
	if err != nil {
		log.Fatalf("Test connection db failure : %v\n", err)
	}
	testRepo = mongodb.NewRepository(db)

	fmt.Sprintln(testRepo)

	os.Exit(m.Run())
}

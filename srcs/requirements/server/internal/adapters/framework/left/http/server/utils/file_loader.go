package utils

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// Retrieves the file, opens it, transforms it into an array of bytes, closes the file and returns the array.
func FileLoader(c *gin.Context) ([]byte, error) {
	log.Println("File uploader has been launched ...")
	var obj models.User

	err := c.ShouldBind(&obj)
	if err != nil {
		return nil, fmt.Errorf("file format is not authorized : it must be `json` file")
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("form file function has been failed")
	}

	openedFile, err := formFile.Open()
	if err != nil {
		return nil, fmt.Errorf("opening file has been failed")
	}
	defer openedFile.Close()

	byteFile, _ := ioutil.ReadAll(openedFile)
	if err != nil {
		return nil, fmt.Errorf("read all file has been failed")

	}

	log.Println("File has been uploaded successfully ...")
	return byteFile, nil
}

package api

import (
	"data_impact/srcs/requirements/server/internal/adapters/framework/left/http/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) CreateUsers(c *gin.Context) {
	status := http.StatusOK
	message := ""

	byteFile, err := utils.FileLoader(c)
	if err != nil {
		status = http.StatusInternalServerError
		message = "File loader has been failed"
	}

	err = m.userService.AddUsers(byteFile)
	if err != nil {
		status = http.StatusInternalServerError
		message = "Add users has been failed"
	}

	c.JSON(status, gin.H{
		"method": "/add/users",
		"error":  message,
	})
}

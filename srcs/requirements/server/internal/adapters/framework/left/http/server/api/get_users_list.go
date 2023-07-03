package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) GetUsersList(c *gin.Context) {
	status := http.StatusOK
	message := ""

	users, err := m.userService.GetUsersList()
	if err != nil {
		status = http.StatusInternalServerError
		message = "Get user has been failed"
		users = nil
	}

	c.JSON(status, gin.H{
		"method": "/user/:id",
		"error":  message,
		"user":   users,
	})
}

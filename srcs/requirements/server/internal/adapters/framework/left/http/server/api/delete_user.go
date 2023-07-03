package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) DeleteUser(c *gin.Context) {
	status := http.StatusOK
	message := ""

	userId := c.Param("id")

	err := m.userService.DeleteUser(userId)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	c.JSON(status, gin.H{
		"method": "/delete/user/:id",
		"error":  message,
	})
}

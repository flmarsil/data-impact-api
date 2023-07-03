package api

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) UpdateUser(c *gin.Context) {
	status := http.StatusOK
	message := ""

	userId := c.Param("id")

	var updatedUser *models.User

	err := json.NewDecoder(c.Request.Body).Decode(&updatedUser)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	// TODO : check here if user is correctly logged in with token
	// but it is not precise in subject

	err = m.userService.UpdateUser(userId, updatedUser)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	c.JSON(status, gin.H{
		"method": "/user/:id",
		"error":  message,
	})
}

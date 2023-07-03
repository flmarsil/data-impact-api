package api

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) Login(c *gin.Context) {
	status := http.StatusOK
	message := ""

	var user *models.User

	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	token, err := m.authService.LogIn(user.UserId, user.Password)
	if err != nil {
		status = http.StatusUnauthorized
		message = err.Error()
	}

	c.JSON(status, gin.H{
		"method": "/login",
		"error":  message,
		"token":  token,
	})
}

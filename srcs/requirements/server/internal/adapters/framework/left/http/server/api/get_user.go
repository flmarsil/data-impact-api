package api

import (
	"data_impact/srcs/requirements/server/internal/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// method GET /user/:id
func (m *MicroserviceServer) GetUser(c *gin.Context) {
	status := http.StatusUnauthorized
	message := "Not authorized"
	user := &models.User{}

	// verify token from request header
	token := c.Request.Header["Token"]
	if token != nil {
		user_id, err := m.tokenManagerService.Parse(token[0])
		if user_id != nil && err == nil {
			// token is valid
			status = http.StatusOK
			message = ""

			user_id := c.Param("id")

			user, err = m.userService.GetUser(user_id)
			if err != nil {
				status = http.StatusInternalServerError
				message = "Get user has been failed"
				user = nil
			}
		} else {
			status = http.StatusForbidden
			message = "Invalid token"
		}
	}

	c.JSON(status, gin.H{
		"method": "/user/:id",
		"error":  message,
		"user":   user,
	})
}

package server

import (
	"data_impact/srcs/requirements/server/internal/adapters/framework/left/http/server/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// is used to create a GIN engine instance where all controller and routes will be placed
func NewRouter(ms *api.MicroserviceServer) *gin.Engine {
	router := gin.New()

	// middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.MaxMultipartMemory = 50000 // change here for update uploading file size

	// routes definition
	router.POST("/add/users", ms.CreateUsers)
	router.POST("/login", ms.Login)
	router.GET("/user/:id", ms.GetUser)
	router.GET("/users/list", ms.GetUsersList)
	router.PUT("/user/:id", ms.UpdateUser)
	router.DELETE("/delete/user/:id", ms.DeleteUser)

	return router
}

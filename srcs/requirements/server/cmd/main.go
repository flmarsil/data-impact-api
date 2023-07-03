package main

import (
	"data_impact/srcs/requirements/server/internal/adapters/framework/left/http/server"
	"data_impact/srcs/requirements/server/internal/adapters/framework/left/http/server/api"
	"data_impact/srcs/requirements/server/internal/adapters/framework/right/mongodb"
	"data_impact/srcs/requirements/server/internal/domain/services/authentification"
	tokenManager "data_impact/srcs/requirements/server/internal/domain/services/token_manager"
	"data_impact/srcs/requirements/server/internal/domain/services/user"
)

const (
	// do not do this in production, use var env or config file
	dbURL = "mongodb://flmarsil:845581EF64F3D5F1C0F17EE7D22524B165A7160FB28828B8420232D8F11D2A40@db:27017/"
	// do not do this in production, use var env or config file
	secretPhrase = "845581EF64F3D5F1C0F17EE7D22524B165A7160FB28828B8420232D8F11D2A40"
)

func main() {
	// TODO : change dbURL by getting env var
	db, err := mongodb.NewDB(dbURL)
	if err != nil {
		return
	}
	defer mongodb.CloseDB()

	// init db
	repo := mongodb.NewRepository(db)
	// init token manager
	tokenManager := tokenManager.NewTokenManager(secretPhrase)

	// register all services
	authService := authentification.NewAuthService(repo, tokenManager)
	userService := user.NewServiceUser(repo)

	server.HttpLauncher(api.NewMicroService(userService, tokenManager, authService))
}

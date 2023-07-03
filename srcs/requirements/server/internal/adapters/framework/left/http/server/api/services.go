package api

import (
	as "data_impact/srcs/requirements/server/internal/domain/services/authentification"
	tm "data_impact/srcs/requirements/server/internal/domain/services/token_manager"
	us "data_impact/srcs/requirements/server/internal/domain/services/user"
)

type MicroserviceServer struct {
	userService         us.User
	authService         as.Auth
	tokenManagerService tm.TokenManager
}

func NewMicroService(userService us.User, tokenManagerService tm.TokenManager, authService as.Auth) *MicroserviceServer {
	return &MicroserviceServer{
		userService:         userService,
		tokenManagerService: tokenManagerService,
		authService:         authService,
	}
}

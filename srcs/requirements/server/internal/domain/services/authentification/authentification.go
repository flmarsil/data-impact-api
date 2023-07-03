package authentification

import (
	"data_impact/srcs/requirements/server/internal/adapters/framework/right/mongodb"
	tokenManager "data_impact/srcs/requirements/server/internal/domain/services/token_manager"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	LogIn(user_id, password string) (*string, error)
}

type auth struct {
	repo        mongodb.Repository
	tokenManger tokenManager.TokenManager
}

func NewAuthService(repo mongodb.Repository, tokenManager tokenManager.TokenManager) Auth {
	return &auth{repo: repo, tokenManger: tokenManager}
}

func (a *auth) LogIn(user_id, user_password string) (*string, error) {
	// get hash password from database by user id
	password, err := a.repo.UserQuery().GetUserHashedPassword(user_id)
	if err != nil {
		return nil, err
	}

	// compare hash from db and user password plain text used for login
	err = bcrypt.CompareHashAndPassword([]byte(*password), []byte(user_password))
	if err != nil {
		return nil, fmt.Errorf("passwords don't match : %v", err)
	} else {
		user, err := a.repo.UserQuery().GetUser(user_id)
		if err != nil {
			return nil, err
		}
		// generate new token
		jwt, err := a.tokenManger.NewJWT(user.UserId)
		if err != nil {
			return nil, err
		}
		return &jwt, nil
	}
}

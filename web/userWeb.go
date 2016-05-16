package web

import (
	"net/http"
	"github.com/tcw/go-graph/repository"
)

type UserWeb struct {
	userDb repository.UserRepository
}

func NewUserRepository(userRepo repository.UserRepository) UserWeb {
	return UserWeb{userRepo}
}

func (uw UserWeb) AddUserHandler(username string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uw.userDb.SaveUser(username)
	}
	return http.HandlerFunc(fn)
}
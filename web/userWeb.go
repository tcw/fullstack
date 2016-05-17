package web

import (
	"net/http"
	"github.com/tcw/go-graph/repository"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
)

type UserWeb struct {
	userDb repository.UserRepository
	render *render.Render
}

func NewUserWeb(userRepo repository.UserRepository) UserWeb {
	return UserWeb{userRepo,render.New()}
}

func (uw UserWeb) AddUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uw.userDb.SaveUser(repository.User{Username:vars["username"],Lastname:vars["lastname"]})
	}
	return http.HandlerFunc(fn)
}

func (uw UserWeb) GetUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		user := uw.userDb.GetUser(username)
		uw.render.JSON(w, http.StatusOK, user)
	}
	return http.HandlerFunc(fn)
}
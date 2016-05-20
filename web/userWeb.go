package web

import (
	"net/http"
	"github.com/tcw/go-graph/repository"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)

type UserWeb struct {
	userDb repository.UserRepository
	render *render.Render
}

type ErrorResponse struct {
	httpStatus int32
	message    string
	context    map[string]string
}

func NewUserWeb(userRepo repository.UserRepository) UserWeb {
	return UserWeb{userRepo, render.New()}
}

func (uw UserWeb) AddUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user repository.User
		err := decoder.Decode(&user)
		if err != nil {
			errorMessage := fmt.Sprintf("Error decoding json %s", r.Body)
			uw.render.JSON(w, http.StatusInternalServerError, ErrorResponse{httpStatus:500, message:errorMessage})
		}else{
			uw.userDb.SaveUser(user)
			uw.render.JSON(w, http.StatusCreated,nil)
		}
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
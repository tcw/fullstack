package web

import (
	"net/http"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"github.com/tcw/fullstack/repository"
	"github.com/tcw/fullstack/domain"
)

var serializer *render.Render = render.New()

type UserWeb struct {
	userDb repository.UserRepository
}

type ErrorResponse struct {
	httpStatus int32
	message    string
	context    map[string]string
}

func NewUserWeb(userRepo repository.UserRepository) UserWeb {
	return UserWeb{userRepo}
}

func (uw UserWeb) AddUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user domain.User
		err := decoder.Decode(&user)
		if err != nil {
			errorMessage := fmt.Sprintf("Error decoding json %s", r.Body)
			serializer.JSON(w, http.StatusInternalServerError, ErrorResponse{httpStatus:500, message:errorMessage})
		}else{
			uw.userDb.SaveUser(user)
			serializer.JSON(w, http.StatusCreated,nil)
		}
	}
	return http.HandlerFunc(fn)
}

func (uw UserWeb) GetUserHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		users := uw.userDb.GetUser(username)
		if len(users.Users) > 0 {
			serializer.JSON(w, http.StatusOK, users)
		}else {
			serializer.JSON(w, http.StatusOK, nil)
		}
	}
	return http.HandlerFunc(fn)
}
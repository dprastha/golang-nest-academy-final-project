package service

import (
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"log"
	"time"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) CreateUser(request *params.User) *view.Response {
	user, err := request.ParseToModel()
	if err != nil {
		log.Printf("Error when casting request to user model %v\n", err)
		return view.ErrorRequest("CREATED_USER_FAIL", "BAD_REQUEST")
	}

	// TODO: change id value to ...
	user.Id = "string random"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// TODO: change to dynamic
	user.Role = "user"

	return view.UserCreated("CREATED_USER_SUCCESS")
}

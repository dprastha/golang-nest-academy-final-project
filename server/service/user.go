package service

import (
	"final-project/helper"
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserServices(repo repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Register(req *params.UserRegister) *view.Response {
	user := req.ParseToModel()

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashedPassword, err := helper.GeneratePassword(user.Password)
	if err != nil {
		return view.ErrorResponse("Failed to hash password", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	user.Password = hashedPassword

	err = u.repo.Register(user)
	if err != nil {
		return view.ErrorResponse("Failed to register", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("Success Register", user, http.StatusCreated)
}

func (u *UserService) CreateUser(request *params.User) *view.Response {
	user, err := request.ParseToModelUser()
	if err != nil {
		log.Printf("Error when casting request to user model %v\n", err)
		return view.ErrorResponse("CREATED_USER_FAIL", "BAD_REQUEST", http.StatusBadRequest)
	}

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// TODO: change to dynamic
	user.Role = "user"

	err = u.repo.Register(user)
	if err != nil {
		log.Printf("Error when create in repository to user model %v\n", err)
		return view.ErrorResponse("CREATED_USER_FAIL", "BAD_REQUEST", http.StatusBadRequest)
	}

	return view.SuccessResponse("CREATED_USER_SUCCESS", user, http.StatusCreated)
}

package service

import (
	"final-project/helper"
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepo
}

func NewService(repo repository.UserRepo) *UserService {
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

	return view.SuccessResponse("Success Register", http.StatusCreated)
}

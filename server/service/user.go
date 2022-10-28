package service

import (
	"database/sql"
	"final-project/helper"
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (u *UserService) Login(req *params.UserLogin) *view.Response {
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return view.ErrorResponse("User not found", "NOT_FOUND", http.StatusNotFound)
		}

		return view.ErrorResponse("Failed to login", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	err = helper.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return view.ErrorResponse("Failed to login", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("Succes login", user, http.StatusOK)
}

func (u *UserService) FindUserByEmail(email string) *view.Response {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrorResponse("User not found", "NOT_FOUND", http.StatusNotFound)
		}

		return view.ErrorResponse("Failed to find user", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("Succes find user", user, http.StatusOK)
}

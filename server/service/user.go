package service

import (
	"database/sql"
	"final-project/helper"
	"final-project/server/params"
	"final-project/server/repository"
	"final-project/server/view"
	"log"
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

	registeredEmail, _ := u.repo.FindUserByEmail(req.Email)
	if registeredEmail != nil {
		return view.ErrorResponse("Email already registered", "UNPROCESSABLE_ENTITY", http.StatusUnprocessableEntity)
	}

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
		return view.ErrorResponse("Credential invalid", "UNAUTHORIZED", http.StatusUnauthorized)
	}

	tokenPayload := helper.Token{
		UserId: user.Id,
		Email:  user.Email,
	}

	tokenString, err := helper.GenerateToken(&tokenPayload)
	if err != nil {
		return view.ErrorResponse("Failed to login", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	user.AccessToken = tokenString

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

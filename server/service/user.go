package service

import (
	"database/sql"
	"final-project/adaptor"
	"final-project/helper"
	"final-project/server/model"
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
	repo       repository.UserRepo
	rajaOngkir adaptor.RajaOngkirAdaptor
}

func NewUserServices(repo repository.UserRepo, rajaOngkir adaptor.RajaOngkirAdaptor) *UserService {
	return &UserService{
		repo:       repo,
		rajaOngkir: rajaOngkir,
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

	return view.SuccessUserResponse("Success Register", http.StatusCreated)
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

	tokenResponse := map[string]string{
		"token": tokenString,
	}

	return view.SuccessResponse("Succes login", tokenResponse, http.StatusOK)
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
	user.Role = "cashier"

	hashedPassword, err := helper.GeneratePassword("default")
	if err != nil {
		return view.ErrorResponse("Failed to hash password", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	user.Password = hashedPassword

	err = u.repo.Register(user)
	if err != nil {
		log.Printf("Error when create in repository to user model %v\n", err)
		return view.ErrorResponse("CREATED_USER_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessUserResponse("CREATED_USER_SUCCESS", http.StatusCreated)
}

func (u *UserService) GetUsers(page int, limit int) *view.Response {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 25
	}

	users, err := u.repo.GetUsers(page, limit)
	if err != nil {
		log.Printf("Error when find all users in service get users %v\n", err)
		return view.ErrorResponse("GET_ALL_USERS_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	payload, err := view.NewAllUsers(users, &u.rajaOngkir)
	if err != nil {
		log.Printf("Error when try to hit raja ongkir in service get users %v\n", err)
		return view.ErrorResponse("GET_ALL_USERS_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	query := &model.Pagination{
		Limit: limit,
		Page:  page,
		Total: len(*users),
	}

	return view.SuccessAllUsersResponse(query, payload)
}

func (u *UserService) ShowUserByEmail(email string) *view.Response {
	user, err := u.repo.FindUserByEmail(email)
	if err == gorm.ErrRecordNotFound {
		log.Printf("Error when try to get user by email %v\n", err)
		return view.ErrorResponse("GET_USER_BY_EMAIL_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	payload, err := view.NewUsers(user, &u.rajaOngkir)
	if err != nil {
		log.Printf("Error when try to hit raja ongkir in service show user by email %v\n", err)
		return view.ErrorResponse("GET_USER_BY_EMAIL_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("GET_USER_BY_EMAIL_SUCCESS", payload, http.StatusOK)
}

func (u *UserService) ShowUserById(id string) *view.Response {
	user, err := u.repo.DetailUserById(id)
	if err == gorm.ErrRecordNotFound {
		log.Printf("Error when try to get user by id %v\n", err)
		return view.ErrorResponse("GET_USER_PROFILE_FAIL", "NOT_FOUND", http.StatusNotFound)
	}

	payload, err := view.NewUsers(user, &u.rajaOngkir)
	if err != nil {
		log.Printf("Error when try to hit raja ongkir in service show user by id %v\n", err)
		return view.ErrorResponse("GET_USER_PROFILE_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessResponse("GET_USER_PROFILE_SUCCESS", payload, http.StatusOK)
}

func (u *UserService) UpdateUser(request *params.UpdateUser, id string) *view.Response {
	user, err := request.ParseUpdateToModelUser()
	if err != nil {
		log.Printf("Error when casting request to update user model %v\n", err)
		return view.ErrorResponse("UPDATE_USER_FAIL", "BAD_REQUEST", http.StatusBadRequest)
	}

	user.UpdatedAt = time.Now()
	err = u.repo.EditUser(user, id)
	if err != nil {
		log.Printf("Error when update user in repository to user model %v\n", err)
		return view.ErrorResponse("UPDATE_USER_FAIL", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	}

	return view.SuccessUserResponse("UPDATE_USER_SUCCESS", http.StatusOK)
}

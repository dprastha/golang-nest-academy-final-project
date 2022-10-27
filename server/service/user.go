package service

import (
	"final-project/helper"
	"final-project/server/params"
	"final-project/server/repository"
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

func (u *UserService) Register(req *params.UserRegister) {
	user := req.ParseToModel()

	user.Id = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashedPassword, err := helper.GeneratePassword(user.Password)
	if err != nil {
		// TODO add error response
	}

	user.Password = hashedPassword

	err = u.repo.Register(user)
	if err != nil {
		// TODO add error response
	}

	// TODO add success response
	return
}

package repository

import "final-project/server/model"

type UserRepo interface {
	Register(user *model.User) error
}

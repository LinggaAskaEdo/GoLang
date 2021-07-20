package repository

import "clean-arch-one/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

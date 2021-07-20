package presenter

import "clean-arch-one/domain/model"

type UserPresenter interface {
	ResponseUser(u []*model.User) []*model.User
}

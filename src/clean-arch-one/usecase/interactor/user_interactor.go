package interactor

import (
	"clean-arch-one/domain/model"
	"clean-arch-one/usecase/presenter"
	"clean-arch-one/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)

	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUser(u), nil
}

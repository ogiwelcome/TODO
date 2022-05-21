package usecase

import (
	"todo/domain"
	"todo/infra"
)

func AddUser(user *domain.User) (*domain.User, error) {
	user, err := infra.AddUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

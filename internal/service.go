package internal

import (
	"USERLOGIN/model"
	store "USERLOGIN/persistent"
)

type Service struct {
	store store.Store
}

func NewService(store store.Store) *Service {
	return &Service{store: store}
}

func (s *Service) UserLogin(username, pass string) (model.User, error) {

	user, err := s.store.UserLogin(username, pass)
	if err != nil {
		return model.User{}, err
	}

	return user, err
}

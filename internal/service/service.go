package service

import (
	"context"

	"gthub.com/Vladroon22/TestTask/internal/entity"
)

type Servicer interface {
	CreateUser(c context.Context, user entity.User) error
	UpdateUser(c context.Context, user entity.User) (int, error)
	GetUser(c context.Context, id int) (entity.User, error)
}

type Service struct {
	repo Servicer
}

func NewService(repo Servicer) Servicer {
	return &Service{repo: repo}
}
func (s *Service) CreateUser(c context.Context, user entity.User) error {
	return s.repo.CreateUser(c, user)
}

func (s *Service) UpdateUser(c context.Context, user entity.User) (int, error) {
	return s.repo.UpdateUser(c, user)
}
func (s *Service) GetUser(c context.Context, id int) (entity.User, error) {
	return s.repo.GetUser(c, id)
}

package service

import (
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/mauricioww/goauth/app/models"
	"github.com/mauricioww/goauth/app/repository"
	"github.com/mauricioww/goauth/app/utils"
)

type Service struct {
	repository repository.Repository
	logger     log.Logger
}

type Servicer interface {
	CreateUser(name string, lastName string, password string, admin bool) (bool, error)
}

func NewSercive(r repository.Repository, l log.Logger) *Service {
	return &Service{
		repository: r,
		logger:     l,
	}
}

func (s *Service) CreateUser(name string, lastName string, password string, admin bool) (bool, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	hash, err := utils.HashPassword(password)

	if err != nil {
		level.Error(logger).Log("Hash error: ", err)
	}

	res, err := s.repository.CreateUser(name, lastName, hash, admin)

	if err != nil {
		level.Error(logger).Log("Transaction error: ", err)
		return res, err
	}

	logger.Log("action", "success")
	return res, nil
}

func (s *Service) GetUser(email string) (models.User, error) {
	logger := log.With(s.logger, "method", "GetUser")

	res, err := s.repository.GetUser(email)

	if err != nil {
		level.Error(logger).Log("Transaction error: ", err)
		return models.User{}, err
	}

	logger.Log("action", "success")
	return res, nil
}

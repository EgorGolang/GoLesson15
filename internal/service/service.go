package service

import (
	"GoLessonFifteen/internal/contracts"
)

type Service struct {
	repository contracts.RepositoryI
}

func NewService(repository contracts.RepositoryI) *Service {
	return &Service{repository: repository}
}

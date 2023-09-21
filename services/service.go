package services

import "github.com/ezequiel-bugnon/brandmonitor/repository"

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository,
	}
}

func (s *service) PostFile() {

}

func (s *service) GetData() {

}

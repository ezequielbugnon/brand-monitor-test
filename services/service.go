package services

import (
	"github.com/ezequiel-bugnon/brandmonitor/dto"
	"github.com/ezequiel-bugnon/brandmonitor/entity"
	"github.com/ezequiel-bugnon/brandmonitor/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository,
	}
}

func (s *service) PostFile(data dto.FileDto) error {
	fileEntity := entity.FileMongoEntity{
		Indicador1: data.Indicador1,
		Indicador2: data.Indicador2,
		Indicador3: data.Indicador3,
	}
	return s.repository.Post(fileEntity)
}

func (s *service) GetData() ([]entity.FileMongoEntity, error) {

	return s.repository.Get()
}

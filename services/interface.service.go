package services

import (
	"github.com/ezequiel-bugnon/brandmonitor/dto"
	"github.com/ezequiel-bugnon/brandmonitor/entity"
)

type Service interface {
	PostFile(data dto.FileDto) error
	GetData() ([]entity.FileMongoEntity, error)
}

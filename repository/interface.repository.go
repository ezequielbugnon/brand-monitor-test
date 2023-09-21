package repository

import (
	"github.com/ezequiel-bugnon/brandmonitor/entity"
)

type Repository interface {
	Post(file entity.FileMongoEntity) error
	Get() ([]entity.FileMongoEntity, error)
}

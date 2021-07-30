package usecase

import (
	"app/models"
	"app/repository/database"
	"app/repository/redis"
)

type DatabaseUseCase struct {
	db  database.Database
	rdb redis.RedisDB
}

func NewDBService(db database.Database) *DatabaseUseCase {
	return &DatabaseUseCase{
		db: db,
	}
}

func (uc *DatabaseUseCase) CreateProxy(proxy *models.Nodes) error {
	return uc.db.Create(proxy)
}

func (uc *DatabaseUseCase) FindProxyByHostname(proxy *models.Nodes, hostname string) error {
	return uc.db.FindProxyByHostname(proxy, hostname)
}

func (uc *DatabaseUseCase) FindNodes(nodes *[]models.Nodes) error {
	return uc.db.FindNodes(nodes)
}

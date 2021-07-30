package database

import (
	"gorm.io/gorm"
)

type Database interface {
	Create(target interface{}) error
	FindProxyByHostname(target interface{}, hostname string) error
	FindNodes(target interface{}) error
}

type Postgres struct {
	db *gorm.DB
}

func NewPostges(db *gorm.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (psql *Postgres) Create(target interface{}) error {
	if err := psql.db.Create(target).Error; err != nil {
		return err
	}
	return nil
}

func (psql *Postgres) FindProxyByHostname(target interface{}, hostname string) error {
	if err := psql.db.Where("hostname = ?", hostname).First(target).Error; err != nil {
		return err
	}
	return nil
}

func (psql *Postgres) FindNodes(target interface{}) error {
	if err := psql.db.Find(target).Error; err != nil {
		return err
	}
	return nil
}

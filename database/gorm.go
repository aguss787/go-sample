package database

import (
	"api/config"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newConnection(connectionUrl string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connectionUrl))
}

func Register(container *dig.Container) (*dig.Container, error) {
	if err := container.Provide(func(config config.Config) (*gorm.DB, error) {
		return newConnection(config.Postgres.DatabaseUrl)
	}); err != nil {
		return nil, err
	}

	return container, nil
}

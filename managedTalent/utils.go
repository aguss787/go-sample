package managedTalent

import (
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func Register(container *dig.Container) (*dig.Container, error) {
	if err := container.Provide(func(hubberApi HubberAPI) (*resolver, error) {
		return &resolver{
			hubberApi,
		}, nil
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(db *gorm.DB) (HubberAPI, error) {
		return &HubberService{
			db,
		}, nil
	}); err != nil {
		return nil, err
	}

	return container, nil
}

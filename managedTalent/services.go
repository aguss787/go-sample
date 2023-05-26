package managedTalent

import (
	"api/managedTalent/models"
	"api/utils"
	"gorm.io/gorm"
)

type HubberAPI interface {
	GetAllHubbers() ([]Hubber, error)
}

type HubberService struct {
	db *gorm.DB
}

func (h *HubberService) GetAllHubbers() ([]Hubber, error) {
	var hubbers []models.Hubber
	h.db.Find(&hubbers)

	return utils.Fmap(hubbers, func(i models.Hubber) Hubber {
		return Hubber{
			Id:   i.ID,
			Code: i.Code,
			Name: i.Name,
		}
	}), nil
}

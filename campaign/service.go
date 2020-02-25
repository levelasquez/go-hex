package campaign

import (
	"time"

	"github.com/google/uuid"
)

// Service - interface
type Service interface {
	CreateCampaign(campaign Campaign) error
	FindCampaignByID(id string) (Campaign, error)
	FindAllCampaigns() ([]Campaign, error)
}

type service struct {
	repository Repository
}

// NewService - init
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CreateCampaign(campaign Campaign) error {
	campaign.ID = uuid.New().String()
	campaign.CreatedAt = time.Now()
	campaign.UpdatedAt = time.Now()

	return s.repository.Create(campaign)
}

func (s *service) FindCampaignByID(id string) (Campaign, error) {
	return s.repository.FindByID(id)
}

func (s *service) FindAllCampaigns() ([]Campaign, error) {
	return s.repository.FindAll()
}

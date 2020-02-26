package campaign

import (
	"time"

	"github.com/google/uuid"
)

// CampaignService - business logic
type CampaignService interface {
	CreateCampaign(campaign *Campaign) error
	FindCampaignByID(id string) (*Campaign, error)
	FindAllCampaigns() ([]*Campaign, error)
}

type campaignService struct {
	repo CampaignRepository
}

// NewCampaignService - initialize the service
func NewCampaignService(repo CampaignRepository) CampaignService {
	return &campaignService{
		repo,
	}
}

func (s *campaignService) CreateCampaign(campaign *Campaign) error {
	campaign.ID = uuid.New().String()
	campaign.CreatedAt = time.Now()
	campaign.UpdatedAt = time.Now()

	return s.repo.Create(campaign)
}

func (s *campaignService) FindCampaignByID(id string) (*Campaign, error) {
	return s.repo.FindByID(id)
}

func (s *campaignService) FindAllCampaigns() ([]*Campaign, error) {
	return s.repo.FindAll()
}

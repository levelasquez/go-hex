package campaign

// CampaignRepository - capacities of the domain
type CampaignRepository interface {
	Create(campaign *Campaign) error
	FindByID(id string) (*Campaign, error)
	FindAll() ([]*Campaign, error)
}

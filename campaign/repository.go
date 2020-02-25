package campaign

// Repository - interface for campaign
type Repository interface {
	Create(campaign Campaign) (err error)
	FindByID(id string) (campaign Campaign, err error)
	FindAll() (campaigns []Campaign, err error)
}

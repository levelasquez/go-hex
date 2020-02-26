package redis

import (
	"encoding/json"
	"gohex/campaign"

	"github.com/go-redis/redis"
)

const table = "campaigns"

type campaignRepository struct {
	connection *redis.Client
}

// NewRedisCampaignRepository - acts as a port
func NewRedisCampaignRepository(connection *redis.Client) campaign.CampaignRepository {
	return &campaignRepository{
		connection,
	}
}

func (r *campaignRepository) Create(campaign *campaign.Campaign) error {
	encoded, err := json.Marshal(campaign)

	if err != nil {
		return err
	}

	r.connection.HSet(table, campaign.ID, encoded)

	return nil
}

func (r *campaignRepository) FindByID(id string) (*campaign.Campaign, error) {
	bytes, err := r.connection.HGet(table, id).Bytes()

	if err != nil {
		return nil, err
	}

	campaign := new(campaign.Campaign)

	err = json.Unmarshal(bytes, campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (r *campaignRepository) FindAll() (campaigns []*campaign.Campaign, err error) {
	values := r.connection.HGetAll(table).Val()

	for key, value := range values {
		campaign := new(campaign.Campaign)
		err = json.Unmarshal([]byte(value), campaign)

		if err != nil {
			return nil, err
		}

		campaign.ID = key
		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}

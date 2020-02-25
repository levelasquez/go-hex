package redis

import (
	"encoding/json"
	"gohex/campaign"

	"github.com/go-redis/redis"
)

const table = "campaigns"

type repository struct {
	connection *redis.Client
}

// New - actua como puerto
func New(connection *redis.Client) campaign.Repository {
	return &repository{
		connection,
	}
}

func (r *repository) Create(campaign campaign.Campaign) (err error) {
	encoded, err := json.Marshal(campaign)

	if err != nil {
		return
	}

	r.connection.HSet(table, campaign.ID, encoded)

	return
}

func (r *repository) FindByID(id string) (campaign campaign.Campaign, err error) {
	bytes, err := r.connection.HGet(table, id).Bytes()

	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &campaign)

	if err != nil {
		return
	}

	return
}

func (r *repository) FindAll() (campaigns []campaign.Campaign, err error) {
	values := r.connection.HGetAll(table).Val()

	for key, value := range values {
		campaign := new(campaign.Campaign)
		err = json.Unmarshal([]byte(value), campaign)

		if err != nil {
			return
		}

		campaign.ID = key
		campaigns = append(campaigns, *campaign)
	}

	return
}

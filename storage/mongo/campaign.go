package mongo

import (
	"context"
	"gohex/campaign"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const table = "campaigns"

type campaignRepository struct {
	db *mongo.Database
}

// NewMongoCampaignRepository - acts as a port
func NewMongoCampaignRepository(db *mongo.Database) campaign.CampaignRepository {
	return &campaignRepository{
		db,
	}
}

func (r *campaignRepository) Create(campaign *campaign.Campaign) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	r.db.Collection(table).InsertOne(ctx, campaign)

	return nil
}

func (r *campaignRepository) FindByID(id string) (*campaign.Campaign, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	campaign := new(campaign.Campaign)

	filter := bson.M{"id": id}
	err := r.db.Collection(table).FindOne(ctx, filter).Decode(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (r *campaignRepository) FindAll() (campaigns []*campaign.Campaign, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.db.Collection(table).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &campaigns)

	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

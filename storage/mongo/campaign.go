package mongo

import (
	"context"
	"gohex/campaign"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const table = "campaigns"

type repository struct {
	db *mongo.Database
}

// New - actua como puerto
func New(db *mongo.Database) campaign.Repository {
	return &repository{
		db,
	}
}

func (r *repository) Create(campaign campaign.Campaign) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = r.db.Collection(table).InsertOne(ctx, campaign)

	return
}

func (r *repository) FindByID(id string) (campaign campaign.Campaign, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return
	}

	filter := bson.M{"_id": oid}
	err = r.db.Collection(table).FindOne(ctx, filter).Decode(&campaign)

	return
}

func (r *repository) FindAll() (campaigns []campaign.Campaign, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.db.Collection(table).Find(ctx, bson.M{})

	if err != nil {
		return
	}

	err = cursor.All(ctx, &campaigns)

	return
}

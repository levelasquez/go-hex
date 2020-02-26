package main

import (
	"context"
	"flag"
	"fmt"
	"gohex/campaign"
	mongodb "gohex/storage/mongo"
	redisdb "gohex/storage/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	dbType := flag.String("database", "redis", "database type [redis, mongo]")
	flag.Parse()

	var campaignRepo campaign.CampaignRepository

	switch *dbType {
	case "redis":
		redisConnection := connectRedis("localhost:6379")
		defer redisConnection.Close()
		campaignRepo = redisdb.NewRedisCampaignRepository(redisConnection)
	case "mongo":
		mongoConnection := connectMongo("mongodb://localhost:27017")
		campaignRepo = mongodb.NewMongoCampaignRepository(mongoConnection)
	default:
		panic("Unknown database")
	}

	service := campaign.NewCampaignService(campaignRepo)
	handler := campaign.NewCampaignHandler(service)

	router := gin.New()

	router.Use(gin.Recovery())
	router.POST("/campaigns", handler.Create)
	router.GET("/campaings/:id", handler.GetByID)
	router.GET("/campaings", handler.Get)

	fmt.Println("Listening on http://localhost:3000")
	router.Run(":3000")
}

func connectRedis(url string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	err := client.Ping().Err()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Redis DB")

	return client
}

func connectMongo(url string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(url).SetMaxPoolSize(10))

	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		panic(err)
	}

	db := client.Database("go-hex")

	fmt.Println("Connected to Mongo DB")

	return db
}

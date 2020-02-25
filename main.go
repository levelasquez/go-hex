package main

import (
	"context"
	"flag"
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

	var repository campaign.Repository

	switch *dbType {
	case "redis":
		conn := redisConnection("localhost:6379")
		repository = redisdb.New(conn)
	case "mongo":
		conn := mongoConnection("mongodb://localhost:27017")
		repository = mongodb.New(conn)
	default:
		panic("Unknown database")
	}

	service := campaign.Service(repository)
	handler := campaign.Handler(service)

	router := gin.New()

	router.Use(gin.Recovery())
	router.POST("/campaigns", handler.Create)
	router.GET("/campaings/:id", handler.GetByID)
	router.GET("/campaings", handler.GetAll)

	router.Run(":3000")
}

func redisConnection(url string) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	err := client.Ping().Err()

	if err != nil {
		panic(err)
	}

	return
}

func mongoConnection(url string) (db *mongo.Database) {
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

	db = client.Database("go-hex")

	return
}

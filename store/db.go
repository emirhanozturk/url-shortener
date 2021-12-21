package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Url struct {
	OriginalUrl string `bson:"originalurl,omitempty"`
	ShortUrl    string `bson:"shorturl,omitempty"`
}

const CacheDuration = 5 * time.Hour

func ConnectDb() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Print("Error! Database connection failed.")
	}
	return client, err
}

func SaveUrlMapping(shortUrl string, originalUrl string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := ConnectDb()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("urlshortener").Collection("urls")

	_, insertErr := collection.InsertOne(ctx, bson.D{{Key: "originalurl", Value: originalUrl}, {Key: "shorturl", Value: shortUrl}})
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	fmt.Println("Url save is successful:")

}

func RetrieveInitialUrl(shortUrl string) string {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := ConnectDb()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("urlshortener").Collection("urls")

	var result bson.M
	err = collection.FindOne(ctx, bson.M{"shorturl": shortUrl}).Decode(&result)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

	return result["originalurl"].(string)
}

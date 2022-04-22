package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Init(dbURL string) {
	mongoClient = mongoConn(dbURL)
	creatCollection("users")
}

func mongoConn(dbURL string) (client *mongo.Client) {
	log.Println("mongo DB Connection Start")
	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	return client
}

func creatCollection(name string) {
	client := GetMongoClient()
	opt := options.CreateCollection()
	err := client.Database("test").CreateCollection(context.TODO(), name, opt)
	if err != nil {
		log.Println(err.Error())
		log.Println("CANNOT CREATE COLLECTION NAME: " + name)
	} else {
		log.Println("CREATE COLLECTION SUCCESS!")
	}
	_, err = client.Database("test").
		Collection(name).
		Indexes().
		CreateOne(context.Background(), mongo.IndexModel{Keys: bson.D{{Key: "name", Value: 1}}})
	if err != nil {
		log.Println("Create Index Error : ", err.Error())
	}
}

func GetMongoClient() (client *mongo.Client) {
	return mongoClient
}

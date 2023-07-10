package database

import (
	"context"
	"log"
	"strings"

	env "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	s "github.com/mongo_sample_training/structs"
)

func NewClient() (*mongo.Client, error) {

	// URI
	envs, err := env.Read(".env")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(envs["MONGO_URI"]).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	return mongo.Connect(context.TODO(), opts)
}

func Disconnect(client *mongo.Client) error {
	return client.Disconnect(context.TODO())
}

func IsConnected(client *mongo.Client) bool {
	if err := client.Ping(context.TODO(), nil); err != nil {
		return false
	}
	return true
}

func GetCollection(client *mongo.Client, name string) *mongo.Collection {
	return client.Database("sample_training").Collection(name)
}

func GetObjs(client *mongo.Client, name string, filter string ,id string) ([]s.Zip, error) {
	collection := GetCollection(client, name)

	cursor, err := collection.Find(context.TODO(), bson.D{{ filter,bson.D{{ "$eq", strings.ToUpper(id) }} }} )
	if err != nil {
		return nil, err
	}

	var result []s.Zip

	cursor.All(context.TODO(),&result)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func GetAll(client *mongo.Client, name string) ([]s.Zip, error) {
	cursor, err := client.Database("sample_training").Collection(name).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var data []s.Zip

	cursor.All(context.TODO(),&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

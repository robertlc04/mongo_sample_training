package database

import (
	"context"
	"fmt"
	"log"

	env "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
) 

func init() {
	
	err := env.Load(".env")

	err != nil {
		log.Fatalf("Something happends %v\n", err)
	}
}


func NewClient() *mongo.Client, err {
	
	// URI

	uri := os.Getenv("MONGO_URI")

  // Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

  // Create a new client and connect to the server
  returb mongo.Connect(context.TODO(), opts)
}



func main() {
    // Create a new client and connect to the server
  client, err := NewClient()
  if err != nil {
    panic(err)
  }

  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()

  // Send a ping to confirm a successful connection
  if err := client.Database("AnkiDB").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
    panic(err)
  }
	data, err := client.Database("sample_training").Collection("zips").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	defer data.Close(context.TODO())

	for data.Next(context.TODO()) {
		var raw bson.M
		if err := data.Decode(&raw); err != nil {
			panic(err)
		}
		fmt.Println(raw)
	}

}


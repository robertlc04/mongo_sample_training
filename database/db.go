package database

import (
  "context"
  "fmt"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
) 

func main() {
  // Use the SetServerAPIOptions() method to set the Stable API version to 1
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://dazai:Roberto04@ankidb.wvernzb.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

  // Create a new client and connect to the server
  client, err := mongo.Connect(context.TODO(), opts)
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


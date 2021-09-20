package dao

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mwneko:iQcU8rt6EbvNB5EI@cluster0.x8keg.mongodb.net/test_data?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

}

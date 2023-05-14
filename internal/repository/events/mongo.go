package events

import (
	"context"
	"time"

	"github.com/ibra-bybuy/wsports-parser/internal/repository/mongodb"
	"github.com/ibra-bybuy/wsports-parser/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	client *mongodb.Client
}

func NewMongo(client *mongodb.Client) *MongoRepository {
	return &MongoRepository{client}
}

func (m *MongoRepository) Add(ctx context.Context, events *[]model.Event) bool {
	collection := m.client.Database("mongo").Collection("events")

	for _, event := range *events {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var foundEvent model.Event
		findFilter := bson.D{{Key: "id", Value: event.ID}}
		err := collection.FindOne(ctx, findFilter).Decode(&foundEvent)
		if err == mongo.ErrNoDocuments {
			collection.InsertOne(ctx, event)
			continue
		}

		update := bson.D{{Key: "$set", Value: event}}
		collection.UpdateOne(ctx, findFilter, update)
	}

	return false
}

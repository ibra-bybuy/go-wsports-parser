package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/ibra-bybuy/wsports-parser/internal/repository/dotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Client
	Ctx context.Context
}

func New() *Client {
	username := dotenv.Get("MONGO_USERNAME")
	pwd := dotenv.Get("MONGO_PWD")
	host := dotenv.Get("MONGO_HOST")
	port := dotenv.Get("MONGO_PORT")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, pwd, host, port)
	fmt.Println(url)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	return &Client{client, ctx}
}

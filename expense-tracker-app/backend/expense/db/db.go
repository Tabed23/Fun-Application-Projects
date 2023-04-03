package db
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connStr string = "mongodb://localhost:27017"


func Connect() *mongo.Client {
	clientOptions := options.Client()
	clientOptions.ApplyURI(connStr)
	cli, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = cli.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	cli.Database("test").CreateCollection(ctx, "expense")
	return cli

}

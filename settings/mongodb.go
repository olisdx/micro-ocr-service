package settings

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongo db connection
func ConnectMongoDb(collectionName string) *mongo.Collection {

	db, cancel, err := mongoDatabaseConnection()
	if err != nil {
		log.Fatal("Mongo Database Connection Error $s", err)
	}
	fmt.Println("MongoDB connection successfully!")

	bookCollection := db.Collection(collectionName)

	defer cancel()

	return bookCollection
}

func mongoDatabaseConnection() (*mongo.Database, context.CancelFunc, error) {

	mongoUrl := os.Getenv("DB_MONGO_URL")
	mongoDb := os.Getenv("DB_MONGO_DATABASE")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl).SetServerSelectionTimeout(5*time.Second))

	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database(mongoDb)

	return db, cancel, nil

}

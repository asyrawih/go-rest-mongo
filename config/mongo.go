package config

import (
	"context"
	"github.com/hananloser/rest-fiber-mongo-2/exception"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Generate Configuration For Mongo
// params Config
func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx , cancel := NewMongoContext()
	defer cancel()

	client , err := mongo.NewClient(options.Client().ApplyURI(configuration.Get("MONGO_URI")))
	exception.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception.PanicIfNeeded(err)

	database := client.Database(configuration.Get("MONGO_DATABASE"))

	return database
}

// Generate Context Mongo
func NewMongoContext() (context.Context , context.CancelFunc) {
	return context.WithTimeout(context.Background() , 10 * time.Second)
}

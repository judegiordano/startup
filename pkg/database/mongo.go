package database

import (
	"context"
	"fmt"
	"time"

	"github.com/judegiordano/startup/internal"
	"github.com/judegiordano/startup/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(internal.Settings.MongoUri)
	var e error
	Client, e = mongo.Connect(ctx, options)
	if e != nil {
		logger.Fatal(fmt.Sprintf("error connecting to mongo %v", e))
	}
}

func Collection(collection string) *mongo.Collection {
	c := Client.Database(internal.Settings.Database).Collection(collection)
	return c
}

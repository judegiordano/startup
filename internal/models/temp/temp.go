package temp

import (
	"context"
	"time"

	"github.com/judegiordano/startup/pkg/database"
	"github.com/judegiordano/startup/pkg/tools"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
	Id      string `bson:"_id,omitempty" json:"id,omitempty"`
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}

func Collection() *mongo.Collection {
	return database.Collection("temp")
}

func (t Document) Save() (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cuid, err := tools.Cuid()
	if err != nil {
		return nil, err
	}
	t.Id = cuid
	return Collection().InsertOne(ctx, t)
}

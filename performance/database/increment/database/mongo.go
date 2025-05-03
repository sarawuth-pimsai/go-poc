package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	ctx    context.Context
	client *mongo.Client
}

func NewMongo(ctx context.Context) *Mongo {
	uri := "mongodb://root:31Kp9T3Z!Hs+R=ma@localhost:27017"
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Println(err)
	}
	mongo := new(Mongo)
	mongo.ctx = ctx
	mongo.client = client
	return mongo
}
func (m *Mongo) Increment(value uint16) (uint16, error) {
	coll := m.client.Database("performance").Collection("increment")
	_, err := coll.UpdateOne(m.ctx, bson.M{"product_id": 1}, bson.D{{Key: "$inc", Value: bson.D{{Key: "quantity", Value: float64(value)}}}}, options.UpdateOne().SetUpsert(true))
	if err != nil {
		log.Println(err)
	}
	return uint16(value), nil
}

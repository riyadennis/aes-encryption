package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoD struct {
	Collection *mongo.Collection
}

type Store interface {
	Insert(ctx context.Context, d *Data) (*Result, error)
	RetrieveContent(ctx context.Context, id, key string) (string, error)
}

// Blog holds the db structure for the blog
type Data struct {
	Key     string `bson:"key, omitempty"`
	Title   string `bson:"title,omitempty"`
	Content string `bson:"content,omitempty"`
}

type Result struct {
	EncryptionKey string
	EncryptionId  string
}

func (db *MongoD) Insert(ctx context.Context, d *Data) (*Result, error) {
	re, err := db.Collection.InsertOne(ctx, d)
	if err != nil {
		return nil, err
	}
	id, ok := re.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("cant find the id")
	}
	return &Result{
		EncryptionKey: d.Key,
		EncryptionId:  id.Hex(),
	}, nil
}

func (db *MongoD) RetrieveContent(ctx context.Context, id string) (*Data, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": docID}
	data := &Data{}
	err = db.Collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

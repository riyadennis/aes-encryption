package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Insert(ctx context.Context, collection *mongo.Collection, content string) (*Result, error)
}

// Blog holds the db structure for the blog
type Data struct {
	ID      primitive.ObjectID `bson:"id, omitempty"`
	Key     string             `bson:"key, omitempty"`
	Title   string             `bson:"title,omitempty"`
	Content string             `bson:"content,omitempty"`
}

type Result struct {
	EncryptionKey string
	EncryptionId  string
}

func (d *Data) Insert(ctx context.Context, collection *mongo.Collection) (*Result, error) {
	re, err := collection.InsertOne(ctx, d)
	if err != nil {
		return nil, err
	}
	return &Result{
		EncryptionKey: d.Key,
		EncryptionId:  fmt.Sprintf("%s", re.InsertedID),
	}, nil
}

package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB interface {
	Create(ctx context.Context, collection string, document interface{}) (interface{}, error)
	ReadOne(ctx context.Context, collection string, id interface{}) (interface{}, error)
	ReadAll(ctx context.Context, collection string) ([]interface{}, error)
	Update(ctx context.Context, collection string, id interface{}, document interface{}) error
	Delete(ctx context.Context, collection string, id interface{}) error
}

type service struct {
	db *mongo.Database
}

func NewMongoDB(ctx context.Context, opt *options.ClientOptions, databaseName string) (MongoDB, error) {
	// connect
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, errors.New("can not connect: " + err.Error())
	}

	// test connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.New("can not ping: " + err.Error())
	}

	return &service{
		db: client.Database(databaseName),
	}, err
}

func (s *service) Create(ctx context.Context, collection string, document interface{}) (interface{}, error) {
	res, err := s.db.Collection(collection).InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (s *service) ReadOne(ctx context.Context, collection string, id interface{}) (interface{}, error) {
	var res interface{}
	_id, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, err
	}
	err = s.db.Collection(collection).FindOne(ctx, bson.M{"_id": _id}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *service) ReadAll(ctx context.Context, collection string) ([]interface{}, error) {
	var res []interface{}
	cur, err := s.db.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result interface{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		res = append(res, result)
	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *service) Update(ctx context.Context, collection string, id interface{}, document interface{}) error {
	_, err := s.db.Collection(collection).UpdateOne(ctx, bson.M{"_id": id}, document)
	return err
}

func (s *service) Delete(ctx context.Context, collection string, id interface{}) error {
	_, err := s.db.Collection(collection).DeleteOne(ctx, bson.M{"_id": id})
	return err
}

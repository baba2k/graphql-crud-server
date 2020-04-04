package storage

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
	Update(ctx context.Context, collection string, id interface{}, document interface{}) (interface{}, error)
	Delete(ctx context.Context, collection string, id interface{}) (interface{}, error)
}

type mongodbService struct {
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

	return &mongodbService{
		db: client.Database(databaseName),
	}, err
}

func (s *mongodbService) Create(ctx context.Context, collection string, document interface{}) (interface{}, error) {
	doc := map[string]interface{}{}
	for k, v := range document.(map[string]interface{})[collection].(map[string]interface{}) {
		doc[k] = v
	}
	res, err := s.db.Collection(collection).InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return s.ReadOne(ctx, collection, res.InsertedID.(primitive.ObjectID).Hex())
}

func (s *mongodbService) ReadOne(ctx context.Context, collection string, id interface{}) (interface{}, error) {
	var res primitive.D
	_id, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, err
	}
	err = s.db.Collection(collection).FindOne(ctx, bson.M{"_id": _id}).Decode(&res)
	if err != nil {
		return nil, err
	}
	object := res.Map()
	object["id"] = object["_id"].(primitive.ObjectID).Hex()
	delete(object, "_id")
	return convertMongoDocument(object), err
}

func (s *mongodbService) ReadAll(ctx context.Context, collection string) ([]interface{}, error) {
	var res []interface{}
	cur, err := s.db.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result primitive.D
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		object := result.Map()
		object["id"] = object["_id"].(primitive.ObjectID).Hex()
		delete(object, "_id")
		res = append(res, object)
	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *mongodbService) Update(ctx context.Context, collection string, id interface{}, document interface{}) (interface{}, error) {
	doc := map[string]interface{}{}
	for k, v := range document.(map[string]interface{})[collection].(map[string]interface{}) {
		doc[k] = v
	}
	_id, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, err
	}
	_, err = s.db.Collection(collection).UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": doc})
	if err != nil {
		return nil, err
	}
	return s.ReadOne(ctx, collection, id)
}

func (s *mongodbService) Delete(ctx context.Context, collection string, id interface{}) (interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, err
	}

	res, err := s.ReadOne(ctx, collection, id)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Collection(collection).DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, err
	}
	return res, err
}

func convertMongoDocument(object interface{}) interface{} {
	switch val := object.(type) {
	case primitive.M:
		for k, v := range val {
			val[k] = convertMongoDocument(v)
		}
	case primitive.A:
		for i, elem := range val {
			val[i] = convertMongoDocument(elem)
		}
	case primitive.D:
		return convertMongoDocument(val.Map())
	}

	return object
}
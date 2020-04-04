package storage

import (
	"context"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MapDB interface {
	Create(ctx context.Context, collection string, document interface{}) (interface{}, error)
	ReadOne(ctx context.Context, collection string, id interface{}) (interface{}, error)
	ReadAll(ctx context.Context, collection string) ([]interface{}, error)
	Update(ctx context.Context, collection string, id interface{}, document interface{}) (interface{}, error)
	Delete(ctx context.Context, collection string, id interface{}) (interface{}, error)
}

type mapService struct {
	db map[string][]interface{}
}

func NewMapDB() (MapDB, error) {
	return &mapService{
		db: map[string][]interface{}{},
	}, nil
}

func (s *mapService) Create(ctx context.Context, collection string, document interface{}) (interface{}, error) {
	doc := map[string]interface{}{}
	for k, v := range document.(map[string]interface{})[collection].(map[string]interface{}) {
		doc[k] = v
	}
	doc["id"] = primitive.NewObjectID().Hex()
	s.db[collection] = append(s.db[collection], doc)
	return s.ReadOne(ctx, collection, doc["id"])
}

func (s *mapService) ReadOne(ctx context.Context, collection string, id interface{}) (interface{}, error) {
	_id, ok := id.(string)
	if !ok {
		return nil, errors.New("no valid id found")
	}
	if s.db[collection] == nil {
		return nil, errors.New("collection not found")
	}
	for _, d := range s.db[collection] {
		doc, ok := d.(map[string]interface{})
		if !ok {
			return nil, errors.New("document has invalid type: " + reflect.TypeOf(d).String())
		}
		if doc["id"] == _id {
			ctx.Done()
			return doc, nil
		}
	}
	return nil, errors.New("document not found")
}

func (s *mapService) ReadAll(ctx context.Context, collection string) ([]interface{}, error) {
	if s.db[collection] == nil {
		return nil, errors.New("no document found")
	}
	ctx.Done()
	return s.db[collection], nil
}

func (s *mapService) Update(ctx context.Context, collection string, id interface{}, document interface{}) (interface{}, error) {
	_id, ok := id.(string)
	if !ok {
		return nil, errors.New("no valid id found")
	}
	if s.db[collection] == nil {
		return nil, errors.New("collection not found")
	}
	for i, d := range s.db[collection] {
		doc, ok := d.(map[string]interface{})
		if !ok {
			return nil, errors.New("document has invalid type2: " + reflect.TypeOf(d).String())
		}
		if doc["id"] == _id {
			doc, ok := document.(map[string]interface{})
			if !ok {
				return nil, errors.New("document has invalid type: " + reflect.TypeOf(document).String())
			}
			for k,v := range doc[collection].(map[string]interface{}) {
				doc[k] = v
			}
			s.db[collection][i] = doc
			return s.ReadOne(ctx, collection, id)
		}
	}
	return nil, errors.New("document not found")
}

func (s *mapService) Delete(ctx context.Context, collection string, id interface{}) (interface{}, error) {
	_id, ok := id.(string)
	if !ok {
		return nil, errors.New("no valid id found")
	}
	if s.db[collection] == nil {
		return nil, errors.New("collection not found")
	}
	for i, d := range s.db[collection] {
		doc, ok := d.(map[string]interface{})
		if !ok {
			return nil, errors.New("document has invalid type")
		}
		if doc["id"] == _id {
			s.db[collection] = append(s.db[collection][:i], s.db[collection][i+1:]...)
			ctx.Done()
			return doc, nil
		}
	}
	return nil, errors.New("document not found")
}

package storage

import (
	"context"
)

type Database interface {
	Create(ctx context.Context, collection string, document interface{}) (interface{}, error)
	ReadOne(ctx context.Context, collection string, id interface{}) (interface{}, error)
	ReadAll(ctx context.Context, collection string) ([]interface{}, error)
	Update(ctx context.Context, collection string, id interface{}, document interface{}) (interface{}, error)
	Delete(ctx context.Context, collection string, id interface{}) (interface{}, error)
}

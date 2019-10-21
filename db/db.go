package db

import (
	"context"

	"github.com/google/wire"
)

type (
	client struct {
		host DBHost
		name DBName
	}
	db struct {
		client *client
	}
	DB interface{}

	DBHost string
	DBName string
)

func NewClient(context context.Context, host DBHost, name DBName) *client {
	return &client{
		host: host,
		name: name,
	}

}
func NewDB(client *client) DB {

	return &db{
		client: client,
	}
}

var DBSet = wire.NewSet(NewDB, NewClient)

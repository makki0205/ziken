package api

import (
	"github.com/google/wire"
	"github.com/makki0205/ziken/db"
)

type (
	API interface {
	}
	api struct {
	}
)

func NewAPI(db db.DB) API {
	return &api{}
}

var APISet = wire.NewSet(db.DBSet, NewAPI)

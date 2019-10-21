// +build wireinject

package cmd

import (
	"context"

	"github.com/google/wire"
	"github.com/makki0205/ziken/api"
	"github.com/makki0205/ziken/db"
)

type app struct {
	api api.API
}

func NewApp(api api.API) app {
	return app{
		api: api,
	}
}
func initializeApp(ctx context.Context) app {
	wire.Build(
		NewApp,
		api.APISet,
		wire.Value(db.DBHost("host")),
		wire.Value(db.DBName("name")),
	)
	return app{}
}

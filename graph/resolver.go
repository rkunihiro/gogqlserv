package graph

import (
	"github.com/rkunihiro/gogqlserv/internal"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db internal.DBClient
}

func NewResolver(db internal.DBClient) (*Resolver, error) {
	return &Resolver{
		db: db,
	}, nil
}

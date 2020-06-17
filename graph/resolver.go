package graph

import (
	"github.com/rkunihiro/gogqlserv/repos"
)

// Resolver
type Resolver struct {
	userRepo repos.UserRepo
}

// NewResolver
func NewResolver(
	userRepo repos.UserRepo,
) (*Resolver, error) {
	return &Resolver{
		userRepo: userRepo,
	}, nil
}

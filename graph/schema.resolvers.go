package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rkunihiro/gogqlserv/entity"
	"github.com/rkunihiro/gogqlserv/graph/generated"
)

func (r *mutationResolver) AddUser(_ context.Context, name string) (*entity.User, error) {
	return r.userRepo.Register(name)
}

func (r *queryResolver) Users(_ context.Context) ([]*entity.User, error) {
	return r.userRepo.FindAll()
}

func (r *queryResolver) FindUser(_ context.Context, id int) (*entity.User, error) {
	return r.userRepo.FindByID(entity.UserID(id))
}

func (r *userResolver) ID(_ context.Context, obj *entity.User) (int, error) {
	return int(obj.ID), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

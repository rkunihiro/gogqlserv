package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rkunihiro/gogqlserv/entity"
	"github.com/rkunihiro/gogqlserv/graph/generated"
)

func (r *mutationResolver) AddUser(ctx context.Context, name string) (*entity.User, error) {
	user := &entity.User{
		Name: name,
	}
	err := r.db.Save(user)
	return user, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	err := r.db.Find(&users)
	return users, err
}

func (r *queryResolver) FindUser(ctx context.Context, id int) (*entity.User, error) {
	users := make([]*entity.User, 0)
	err := r.db.Find(&users, map[string]interface{}{
		"id": id,
	})
	if len(users) != 1 {
		return nil, nil
	}
	return users[0], err
}

func (r *userResolver) ID(ctx context.Context, obj *entity.User) (int, error) {
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

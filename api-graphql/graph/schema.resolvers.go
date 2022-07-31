package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-graphql/graph/generated"
	"api-graphql/graph/model"
	"api-graphql/users"
	"context"
)

// InsertUser is the resolver for the insertUser field.
func (r *mutationResolver) InsertUser(ctx context.Context, input model.UserInsertInput) (*string, error) {
	user := users.User{
		Email:     *input.Email,
		Name:      *input.Name,
		BirthDate: *input.BirthDate,
		Location: users.Location{
			Latitude:  *input.Location.Latitude,
			Longitude: *input.Location.Longitude,
		},
	}
	result, err := r.Resolver.Service.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, email *string) (*model.User, error) {
	user, err := r.Resolver.Service.GetByEmail(*email)
	if err != nil {
		return nil, err
	}
	result := model.User{
		Name:      &user.Name,
		Location:  &model.Location{Latitude: &user.Location.Latitude, Longitude: &user.Location.Longitude},
		Email:     &user.Email,
		ID:        user.ID,
		BirthDate: user.BirthDate,
	}
	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

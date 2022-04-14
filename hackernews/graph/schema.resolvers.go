package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/khofesh/hackernews/graph/generated"
	"github.com/khofesh/hackernews/graph/model"
	"github.com/khofesh/hackernews/internal/auth"
	"github.com/khofesh/hackernews/internal/links"
	"github.com/khofesh/hackernews/internal/users"
	"github.com/khofesh/hackernews/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	var link links.Link

	link.Address = input.Address
	link.Title = input.Title
	link.User = user

	linkID := link.Save()

	graphqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}

	return &model.Link{
		ID:      strconv.FormatInt(linkID, 10),
		Title:   link.Title,
		Address: link.Address, User: graphqlUser,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	// panic(fmt.Errorf("not implemented"))

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password

	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}

	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks = links.GetAll()

	for _, link := range dbLinks {
		graphqlUser := &model.User{
			ID:   link.User.ID,
			Name: link.User.Username,
		}
		resultLinks = append(resultLinks, &model.Link{
			ID:      link.ID,
			Title:   link.Title,
			Address: link.Address,
			User:    graphqlUser,
		})
	}

	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

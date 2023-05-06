package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/zzihyeon/golang-server.git/graph/generated"
	"github.com/zzihyeon/golang-server.git/graph/model"
	"github.com/zzihyeon/golang-server.git/service"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{{
		UID:    1,
		Email:  "test@sk.com",
		Name:   "zzihyeon",
		Phone:  "010-0000-0000",
		Gender: model.GenderMale,
	}}, nil
}

func (r *queryResolver) Allusers(ctx context.Context) ([]model.UserInterface, error) {
	return service.GetAllUser()
}

func (r *queryResolver) AdminUsers(ctx context.Context) ([]*model.AdminUser, error) {
	return service.GetAdminUser()
}

func (r *queryResolver) SellerUsers(ctx context.Context) ([]*model.SellerUser, error) {
	return service.GetSellerUser()
}

func (r *queryResolver) BuyerUsers(ctx context.Context) ([]*model.BuyerUser, error) {
	return service.GetBuyerUsers()
}

func (r *queryResolver) User(ctx context.Context, input *model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllStuff(ctx context.Context) ([]model.StuffResult, error) {
	return service.GetAllStuff()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

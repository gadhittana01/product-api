package services

import (
	"context"

	"github.com/gadhittana01/product-api/pkg/domain"
)

type (
	ProductResource interface {
		GetOrderHistories(ctx context.Context, req domain.GetOrderHistoriesReq) (domain.GetOrderHistoriesRes, error)
		CreateOrderHistories(ctx context.Context, req domain.CreateOrderHistoriesReq) (domain.OrderHistories, error)
		CreateUser(ctx context.Context, req domain.CreateUserReq) (domain.CreateUserRes, error)
		UpdateUser(ctx context.Context, req domain.UpdateUserReq) (domain.UpdateUserRes, error)
		DeleteUser(ctx context.Context, req domain.DeleteUserReq) (domain.DeleteUserRes, error)
		GetUsers(ctx context.Context, req domain.GetUsersReq) (domain.GetUsersRes, error)
		GetUser(ctx context.Context, req domain.GetUserReq) (domain.Users, error)
		CreateItem(ctx context.Context, req domain.CreateItemReq) (domain.CreateItemRes, error)
		UpdateItem(ctx context.Context, req domain.UpdateItemReq) (domain.UpdateItemRes, error)
		GetItem(ctx context.Context, req domain.GetItemReq) (domain.OrderItems, error)
		DeleteItem(ctx context.Context, req domain.DeleteItemReq) (domain.DeleteItemRes, error)
		GetItems(ctx context.Context, req domain.GetItemsReq) (domain.GetItemsRes, error)
		GetUserCache(ctx context.Context, cacheKey string) (domain.Users, error)
		SetUserCache(ctx context.Context, cacheKey string, data domain.Users) error
		GetItemCache(ctx context.Context, cacheKey string) (domain.OrderItems, error)
		SetItemCache(ctx context.Context, cacheKey string, data domain.OrderItems) error
	}
)

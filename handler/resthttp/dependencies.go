package resthttp

import (
	"context"

	"github.com/gadhittana01/product-api/services"
)

type (
	ProductService interface {
		GetOrderHistories(ctx context.Context, req services.GetOrderHistoriesReq) (services.GetOrderHistoriesRes, error)
		CreateOrderHistories(ctx context.Context, req services.CreateOrderHistoriesReq) (services.OrderHistories, error)
		CreateUser(ctx context.Context, req services.CreateUserReq) (services.CreateUserRes, error)
		UpdateUser(ctx context.Context, req services.UpdateUserReq) (services.UpdateUserRes, error)
		DeleteUser(ctx context.Context, req services.DeleteUserReq) (services.DeleteUserRes, error)
		GetUser(ctx context.Context, req services.GetUserReq) (services.Users, error)
		GetUsers(ctx context.Context, req services.GetUsersReq) (services.GetUsersRes, error)
		CreateItem(ctx context.Context, req services.CreateItemReq) (services.CreateItemRes, error)
		UpdateItem(ctx context.Context, req services.UpdateItemReq) (services.UpdateItemRes, error)
		GetItem(ctx context.Context, req services.GetItemReq) (services.OrderItems, error)
		DeleteItem(ctx context.Context, req services.DeleteItemReq) (services.DeleteItemRes, error)
		GetItems(ctx context.Context, req services.GetItemsReq) (services.GetItemsRes, error)
	}
)

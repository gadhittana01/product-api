package product

import (
	"context"

	"github.com/gadhittana01/product-api/config"
	"github.com/gadhittana01/product-api/pkg/domain"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type IResource interface {
	GetOrderHistories(ctx context.Context, req domain.GetOrderHistoriesReq) (domain.GetOrderHistoriesRes, error)
	CreateOrderHistories(ctx context.Context, req domain.CreateOrderHistoriesReq) (domain.OrderHistories, error)

	CreateUser(ctx context.Context, req domain.CreateUserReq) (domain.CreateUserRes, error)
	GetUsers(ctx context.Context, req domain.GetUsersReq) (domain.GetUsersRes, error)
	UpdateUser(ctx context.Context, req domain.UpdateUserReq) (domain.UpdateUserRes, error)
	GetUser(ctx context.Context, req domain.GetUserReq) (domain.Users, error)
	DeleteUser(ctx context.Context, req domain.DeleteUserReq) (domain.DeleteUserRes, error)

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

type module struct {
	persistent persistent
	secondary  secondary
}

func New(db *gorm.DB, redis *redis.Client, cfg *config.GlobalConfig) IResource {
	return module{
		persistent: newPersistent(db),
		secondary:  newCache(redis, cfg),
	}
}

func (m module) GetOrderHistories(ctx context.Context, req domain.GetOrderHistoriesReq) (domain.GetOrderHistoriesRes, error) {
	return m.persistent.getOrderHistories(ctx, req)
}

func (m module) CreateOrderHistories(ctx context.Context, req domain.CreateOrderHistoriesReq) (domain.OrderHistories, error) {
	return m.persistent.createOrderHistories(ctx, req)
}

func (m module) CreateUser(ctx context.Context, req domain.CreateUserReq) (domain.CreateUserRes, error) {
	return m.persistent.createUser(ctx, req)
}

func (m module) GetUsers(ctx context.Context, req domain.GetUsersReq) (domain.GetUsersRes, error) {
	return m.persistent.getUsers(ctx, req)
}

func (m module) UpdateUser(ctx context.Context, req domain.UpdateUserReq) (domain.UpdateUserRes, error) {
	return m.persistent.updateUser(ctx, req)
}

func (m module) DeleteUser(ctx context.Context, req domain.DeleteUserReq) (domain.DeleteUserRes, error) {
	return m.persistent.deleteUser(ctx, req)
}

func (m module) GetUser(ctx context.Context, req domain.GetUserReq) (domain.Users, error) {
	return m.persistent.getUser(ctx, req)
}

func (m module) CreateItem(ctx context.Context, req domain.CreateItemReq) (domain.CreateItemRes, error) {
	return m.persistent.createItem(ctx, req)
}

func (m module) UpdateItem(ctx context.Context, req domain.UpdateItemReq) (domain.UpdateItemRes, error) {
	return m.persistent.updateItem(ctx, req)
}

func (m module) GetItem(ctx context.Context, req domain.GetItemReq) (domain.OrderItems, error) {
	return m.persistent.getItem(ctx, req)
}

func (m module) DeleteItem(ctx context.Context, req domain.DeleteItemReq) (domain.DeleteItemRes, error) {
	return m.persistent.deleteItem(ctx, req)
}

func (m module) GetItems(ctx context.Context, req domain.GetItemsReq) (domain.GetItemsRes, error) {
	return m.persistent.getItems(ctx, req)
}

func (m module) GetUserCache(ctx context.Context, cacheKey string) (domain.Users, error) {
	return m.secondary.getUserCache(ctx, cacheKey)
}

func (m module) SetUserCache(ctx context.Context, cacheKey string, data domain.Users) error {
	return m.secondary.setUserCache(ctx, cacheKey, data)
}

func (m module) GetItemCache(ctx context.Context, cacheKey string) (domain.OrderItems, error) {
	return m.secondary.getItemCache(ctx, cacheKey)
}

func (m module) SetItemCache(ctx context.Context, cacheKey string, data domain.OrderItems) error {
	return m.secondary.setItemCache(ctx, cacheKey, data)
}

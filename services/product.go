package services

import (
	"context"
	"fmt"

	"github.com/gadhittana01/product-api/pkg/domain"
)

type ProductService interface {
	GetOrderHistories(ctx context.Context, req GetOrderHistoriesReq) (GetOrderHistoriesRes, error)
	CreateOrderHistories(ctx context.Context, req CreateOrderHistoriesReq) (OrderHistories, error)
	CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error)
	UpdateUser(ctx context.Context, req UpdateUserReq) (UpdateUserRes, error)
	DeleteUser(ctx context.Context, req DeleteUserReq) (DeleteUserRes, error)
	GetUsers(ctx context.Context, req GetUsersReq) (GetUsersRes, error)
	GetUser(ctx context.Context, req GetUserReq) (Users, error)
	CreateItem(ctx context.Context, req CreateItemReq) (CreateItemRes, error)
	UpdateItem(ctx context.Context, req UpdateItemReq) (UpdateItemRes, error)
	GetItem(ctx context.Context, req GetItemReq) (OrderItems, error)
	DeleteItem(ctx context.Context, req DeleteItemReq) (DeleteItemRes, error)
	GetItems(ctx context.Context, req GetItemsReq) (GetItemsRes, error)
}

type productService struct {
	pr ProductResource
}

func NewProductService(dep ProductDependencies) (ProductService, error) {
	return &productService{
		pr: dep.PR,
	}, nil
}

func (p productService) GetOrderHistories(ctx context.Context, req GetOrderHistoriesReq) (GetOrderHistoriesRes, error) {
	var result GetOrderHistoriesRes = GetOrderHistoriesRes{}

	res, err := p.pr.GetOrderHistories(ctx, domain.GetOrderHistoriesReq{
		Limit: req.Limit,
		Page:  req.Page,
	})
	if err != nil {
		return result, err
	}

	result = GetOrderHistoriesRes{
		Limit:      res.Limit,
		Page:       res.Page,
		TotalRows:  res.TotalRows,
		TotalPages: res.TotalPages,
	}

	for _, item := range res.OrderHistories {
		result.OrderHistories = append(result.OrderHistories, OrderHistories{
			ID:           item.ID,
			UserID:       item.UserID,
			OrderItemID:  item.OrderItemID,
			Descriptions: item.Descriptions,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}

	return result, nil
}

func (p productService) CreateOrderHistories(ctx context.Context, req CreateOrderHistoriesReq) (OrderHistories, error) {
	var result OrderHistories = OrderHistories{}

	res, err := p.pr.CreateOrderHistories(ctx, domain.CreateOrderHistoriesReq{
		UserID:       req.UserID,
		OrderItemID:  req.OrderItemID,
		Descriptions: req.Descriptions,
	})
	if err != nil {
		return result, err
	}

	result = OrderHistories{
		ID:           res.ID,
		UserID:       res.UserID,
		OrderItemID:  res.OrderItemID,
		Descriptions: res.Descriptions,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}

	return result, nil
}

func (p productService) CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error) {
	var result CreateUserRes = CreateUserRes{}

	res, err := p.pr.CreateUser(ctx, domain.CreateUserReq{
		FullName: req.FullName,
	})
	if err != nil {
		return result, err
	}

	result = CreateUserRes{
		FullName: res.FullName,
	}

	return result, nil
}

func (p productService) GetUsers(ctx context.Context, req GetUsersReq) (GetUsersRes, error) {
	var result GetUsersRes = GetUsersRes{}

	res, err := p.pr.GetUsers(ctx, domain.GetUsersReq{
		Limit: req.Limit,
		Page:  req.Page,
	})
	if err != nil {
		return result, err
	}

	result = GetUsersRes{
		Limit:      res.Limit,
		Page:       res.Page,
		TotalRows:  res.TotalPages,
		TotalPages: res.TotalPages,
	}

	for _, item := range res.Users {
		result.Users = append(result.Users, Users{
			Model:      item.Model,
			FullName:   item.FullName,
			FirstOrder: item.FirstOrder,
		})
	}

	return result, nil
}

func (p productService) UpdateUser(ctx context.Context, req UpdateUserReq) (UpdateUserRes, error) {
	var result UpdateUserRes = UpdateUserRes{}

	res, err := p.pr.UpdateUser(ctx, domain.UpdateUserReq{
		UserID:   req.UserID,
		FullName: req.FullName,
	})
	if err != nil {
		return result, err
	}

	result = UpdateUserRes{
		UserID:   res.UserID,
		FullName: res.FullName,
	}

	return result, nil
}

func (p productService) DeleteUser(ctx context.Context, req DeleteUserReq) (DeleteUserRes, error) {
	var result DeleteUserRes = DeleteUserRes{}

	res, err := p.pr.DeleteUser(ctx, domain.DeleteUserReq{
		UserID: req.UserID,
	})
	if err != nil {
		return result, err
	}

	result = DeleteUserRes{
		UserID:   res.UserID,
		FullName: res.FullName,
	}

	return result, nil
}

func (p productService) GetUser(ctx context.Context, req GetUserReq) (Users, error) {
	var result Users = Users{}
	key := fmt.Sprintf("%s-%d", "user", req.UserID)
	resCache, err := p.pr.GetUserCache(ctx, key)
	if err != nil {
		res, err := p.pr.GetUser(ctx, domain.GetUserReq{
			UserID: req.UserID,
		})
		if err != nil {
			return result, err
		}
		result = Users{
			Model:      res.Model,
			FullName:   res.FullName,
			FirstOrder: res.FirstOrder,
		}
		if err := p.pr.SetUserCache(ctx, key, res); err != nil {
			return result, err
		}

		return result, nil
	}

	result = Users{
		Model:      resCache.Model,
		FullName:   resCache.FullName,
		FirstOrder: resCache.FirstOrder,
	}

	return result, nil
}

func (p productService) CreateItem(ctx context.Context, req CreateItemReq) (CreateItemRes, error) {
	var result CreateItemRes = CreateItemRes{}

	res, err := p.pr.CreateItem(ctx, domain.CreateItemReq{
		Name:      req.Name,
		Price:     req.Price,
		ExpiredAt: req.ExpiredAt,
	})
	if err != nil {
		return result, err
	}

	result = CreateItemRes{
		Name:      res.Name,
		Price:     res.Price,
		ExpiredAt: res.ExpiredAt,
	}

	return result, nil
}

func (p productService) UpdateItem(ctx context.Context, req UpdateItemReq) (UpdateItemRes, error) {
	var result UpdateItemRes = UpdateItemRes{}

	res, err := p.pr.UpdateItem(ctx, domain.UpdateItemReq{
		ID:        req.ID,
		Name:      req.Name,
		Price:     req.Price,
		ExpiredAt: req.ExpiredAt,
	})
	if err != nil {
		return result, err
	}

	result = UpdateItemRes{
		ID:        res.ID,
		Name:      res.Name,
		Price:     res.Price,
		ExpiredAt: res.ExpiredAt,
	}

	return result, nil
}

func (p productService) GetItem(ctx context.Context, req GetItemReq) (OrderItems, error) {
	var result OrderItems = OrderItems{}
	key := fmt.Sprintf("%s-%d", "item", req.ID)
	resCache, err := p.pr.GetItemCache(ctx, key)
	if err != nil {
		res, err := p.pr.GetItem(ctx, domain.GetItemReq{
			ID: req.ID,
		})
		if err != nil {
			return result, err
		}

		result = OrderItems{
			Model:     res.Model,
			Name:      res.Name,
			Price:     res.Price,
			ExpiredAt: res.ExpiredAt,
		}
		if err := p.pr.SetItemCache(ctx, key, res); err != nil {
			return result, err
		}

		return result, nil
	}

	result = OrderItems{
		Model:     resCache.Model,
		Name:      resCache.Name,
		Price:     resCache.Price,
		ExpiredAt: resCache.ExpiredAt,
	}

	return result, nil
}

func (p productService) DeleteItem(ctx context.Context, req DeleteItemReq) (DeleteItemRes, error) {
	var result DeleteItemRes = DeleteItemRes{}

	res, err := p.pr.DeleteItem(ctx, domain.DeleteItemReq{
		ID: req.ID,
	})
	if err != nil {
		return result, err
	}

	result = DeleteItemRes{
		ID:        res.ID,
		Name:      res.Name,
		Price:     res.Price,
		ExpiredAt: res.ExpiredAt,
	}

	return result, nil
}

func (p productService) GetItems(ctx context.Context, req GetItemsReq) (GetItemsRes, error) {
	var result GetItemsRes = GetItemsRes{}

	res, err := p.pr.GetItems(ctx, domain.GetItemsReq{
		Limit: req.Limit,
		Page:  req.Page,
	})
	if err != nil {
		return result, err
	}

	result = GetItemsRes{
		Limit:      res.Limit,
		Page:       res.Page,
		TotalRows:  res.TotalPages,
		TotalPages: res.TotalPages,
	}

	for _, item := range res.OrderItems {
		result.OrderItems = append(result.OrderItems, OrderItems{
			Model:     item.Model,
			Name:      item.Name,
			Price:     item.Price,
			ExpiredAt: item.ExpiredAt,
		})
	}

	return result, nil
}

package product

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/gadhittana01/product-api/pkg/domain"
	"github.com/jinzhu/gorm"
	"golang.org/x/sync/errgroup"
)

type persistent interface {
	// order histories
	getOrderHistories(ctx context.Context, req domain.GetOrderHistoriesReq) (domain.GetOrderHistoriesRes, error)
	createOrderHistories(ctx context.Context, req domain.CreateOrderHistoriesReq) (domain.OrderHistories, error)

	// users
	createUser(ctx context.Context, req domain.CreateUserReq) (domain.CreateUserRes, error)
	updateUser(ctx context.Context, req domain.UpdateUserReq) (domain.UpdateUserRes, error)
	getUser(ctx context.Context, req domain.GetUserReq) (domain.Users, error)
	deleteUser(ctx context.Context, req domain.DeleteUserReq) (domain.DeleteUserRes, error)
	getUsers(ctx context.Context, req domain.GetUsersReq) (domain.GetUsersRes, error)

	// items
	createItem(ctx context.Context, req domain.CreateItemReq) (domain.CreateItemRes, error)
	updateItem(ctx context.Context, req domain.UpdateItemReq) (domain.UpdateItemRes, error)
	getItem(ctx context.Context, req domain.GetItemReq) (domain.OrderItems, error)
	deleteItem(ctx context.Context, req domain.DeleteItemReq) (domain.DeleteItemRes, error)
	getItems(ctx context.Context, req domain.GetItemsReq) (domain.GetItemsRes, error)
}

type psql struct {
	db *gorm.DB
}

func newPersistent(db *gorm.DB) persistent {
	return psql{
		db: db,
	}
}

func (p psql) getOrderHistories(ctx context.Context, req domain.GetOrderHistoriesReq) (domain.GetOrderHistoriesRes, error) {
	var res = domain.GetOrderHistoriesRes{}
	var resQuery = []domain.OrderHistories{}
	err := new(errgroup.Group)

	var offset = (req.Page - 1) * req.Limit
	err.Go(func() error {
		return p.db.Offset(offset).Limit(req.Limit).Find(&resQuery).Error
	})

	totalRows := 0
	err.Go(func() error {
		return p.db.Model(resQuery).Count(&totalRows).Error
	})

	if err := err.Wait(); err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	res = domain.GetOrderHistoriesRes{
		Limit:          req.Limit,
		Page:           req.Page,
		TotalRows:      totalRows,
		TotalPages:     totalPages,
		OrderHistories: resQuery,
	}

	return res, nil
}

func (p psql) createOrderHistories(ctx context.Context, req domain.CreateOrderHistoriesReq) (domain.OrderHistories, error) {
	var res = domain.OrderHistories{}
	var user = domain.Users{}
	err := new(errgroup.Group)
	reqQuery := domain.OrderHistories{
		UserID:       req.UserID,
		OrderItemID:  req.OrderItemID,
		Descriptions: req.Descriptions,
		CreatedAt:    time.Now().Format("02-01-2006"),
		UpdatedAt:    time.Now().Format("02-01-2006"),
	}

	err.Go(func() error {
		var errs error
		if errs = p.db.Where("id = ?", req.UserID).First(&user).Error; !errors.Is(errs, nil) {
			return errs
		}

		if user.FirstOrder == "" {
			user.FirstOrder = time.Now().Format("02-01-2006")
		}

		if errs := p.db.Save(&user).Error; !errors.Is(errs, nil) {
			return errs
		}
		return nil
	})

	err.Go(func() error {
		return p.db.Create(&reqQuery).Scan(&res).Error
	})
	if err := err.Wait(); err != nil {
		return res, err
	}

	return res, nil
}

func (p psql) createUser(ctx context.Context, req domain.CreateUserReq) (domain.CreateUserRes, error) {
	var res = domain.CreateUserRes{}
	err := new(errgroup.Group)

	reqQuery := domain.Users{
		FullName: req.FullName,
	}

	err.Go(func() error {
		return p.db.Select("full_name").Create(&reqQuery).Scan(&res).Error
	})
	if err := err.Wait(); err != nil {
		return res, err
	}

	return res, nil
}

func (p psql) updateUser(ctx context.Context, req domain.UpdateUserReq) (domain.UpdateUserRes, error) {
	var res = domain.UpdateUserRes{}
	err := new(errgroup.Group)
	var user = domain.Users{}

	err.Go(func() error {
		var errs error
		if errs = p.db.Where("id = ?", req.UserID).First(&user).Error; !errors.Is(errs, nil) {
			return errs
		}

		user.FullName = req.FullName

		if errs := p.db.Save(&user).Error; !errors.Is(errs, nil) {
			return errs
		}
		return nil
	})

	if err := err.Wait(); err != nil {
		return res, err
	}

	res = domain.UpdateUserRes{
		UserID:   int(user.ID),
		FullName: user.FullName,
	}

	return res, nil
}

func (p psql) deleteUser(ctx context.Context, req domain.DeleteUserReq) (domain.DeleteUserRes, error) {
	var res = domain.DeleteUserRes{}
	err := new(errgroup.Group)
	var user = domain.Users{}

	if err := p.db.Where("id = ?", req.UserID).First(&user).Error; !errors.Is(err, nil) {
		return res, err
	}

	err.Go(func() error {
		var errs error
		if errs = p.db.Delete(user, req.UserID).Error; !errors.Is(errs, nil) {
			return errs
		}

		return nil
	})

	if err := err.Wait(); err != nil {
		return res, err
	}

	res = domain.DeleteUserRes{
		UserID:   int(user.ID),
		FullName: user.FullName,
	}

	return res, nil
}

func (p psql) getUsers(ctx context.Context, req domain.GetUsersReq) (domain.GetUsersRes, error) {
	var res = domain.GetUsersRes{}
	var resQuery = []domain.Users{}
	err := new(errgroup.Group)

	var offset = (req.Page - 1) * req.Limit
	err.Go(func() error {
		return p.db.Offset(offset).Limit(req.Limit).Find(&resQuery).Error
	})

	totalRows := 0
	err.Go(func() error {
		return p.db.Model(resQuery).Count(&totalRows).Error
	})

	if err := err.Wait(); err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	res = domain.GetUsersRes{
		Limit:      req.Limit,
		Page:       req.Page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Users:      resQuery,
	}

	return res, nil
}

func (p psql) getUser(ctx context.Context, req domain.GetUserReq) (domain.Users, error) {
	var user domain.Users

	if err := p.db.Where("id = ?", req.UserID).First(&user).Error; !errors.Is(err, nil) {
		return user, err
	}

	return user, nil
}

func (p psql) createItem(ctx context.Context, req domain.CreateItemReq) (domain.CreateItemRes, error) {
	var res = domain.CreateItemRes{}
	err := new(errgroup.Group)

	reqQuery := domain.OrderItems{
		Name:      req.Name,
		Price:     req.Price,
		ExpiredAt: req.ExpiredAt,
	}

	err.Go(func() error {
		return p.db.Create(&reqQuery).Scan(&res).Error
	})
	if err := err.Wait(); err != nil {
		return res, err
	}

	return res, nil
}

func (p psql) updateItem(ctx context.Context, req domain.UpdateItemReq) (domain.UpdateItemRes, error) {
	var res = domain.UpdateItemRes{}
	err := new(errgroup.Group)
	var item = domain.OrderItems{}

	err.Go(func() error {
		var errs error
		if errs = p.db.Where("id = ?", req.ID).First(&item).Error; !errors.Is(errs, nil) {
			return errs
		}

		item.Name = req.Name
		item.Price = req.Price
		item.ExpiredAt = req.ExpiredAt

		if errs := p.db.Save(&item).Error; !errors.Is(errs, nil) {
			return errs
		}
		return nil
	})

	if err := err.Wait(); err != nil {
		return res, err
	}

	res = domain.UpdateItemRes{
		ID:        int(item.ID),
		Name:      item.Name,
		Price:     item.Price,
		ExpiredAt: item.ExpiredAt,
	}

	return res, nil
}

func (p psql) getItem(ctx context.Context, req domain.GetItemReq) (domain.OrderItems, error) {
	var item domain.OrderItems

	if err := p.db.Where("id = ?", req.ID).First(&item).Error; !errors.Is(err, nil) {
		return item, err
	}

	return item, nil
}

func (p psql) deleteItem(ctx context.Context, req domain.DeleteItemReq) (domain.DeleteItemRes, error) {
	var res = domain.DeleteItemRes{}
	err := new(errgroup.Group)
	var item = domain.OrderItems{}

	if err := p.db.Where("id = ?", req.ID).First(&item).Error; !errors.Is(err, nil) {
		return res, err
	}

	err.Go(func() error {
		var errs error
		if errs = p.db.Delete(item, req.ID).Error; !errors.Is(errs, nil) {
			return errs
		}

		return nil
	})

	if err := err.Wait(); err != nil {
		return res, err
	}

	res = domain.DeleteItemRes{
		ID:        int(item.ID),
		Name:      item.Name,
		Price:     item.Price,
		ExpiredAt: item.ExpiredAt,
	}

	return res, nil
}

func (p psql) getItems(ctx context.Context, req domain.GetItemsReq) (domain.GetItemsRes, error) {
	var res = domain.GetItemsRes{}
	var resQuery = []domain.OrderItems{}
	err := new(errgroup.Group)

	var offset = (req.Page - 1) * req.Limit
	err.Go(func() error {
		return p.db.Offset(offset).Limit(req.Limit).Find(&resQuery).Error
	})

	totalRows := 0
	err.Go(func() error {
		return p.db.Model(resQuery).Count(&totalRows).Error
	})

	if err := err.Wait(); err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	res = domain.GetItemsRes{
		Limit:      req.Limit,
		Page:       req.Page,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		OrderItems: resQuery,
	}

	return res, nil
}

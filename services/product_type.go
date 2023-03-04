package services

import "github.com/jinzhu/gorm"

type ProductDependencies struct {
	PR ProductResource
}

type GetOrderHistoriesReq struct {
	Limit int
	Page  int
}

type GetOrderHistoriesRes struct {
	Limit          int              `json:"limit"`
	Page           int              `json:"page"`
	TotalRows      int              `json:"total_rows"`
	TotalPages     int              `json:"total_pages"`
	OrderHistories []OrderHistories `json:"order_histories"`
}

type OrderHistories struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	OrderItemID  int    `json:"order_item_id"`
	Descriptions string `json:"descriptions"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Users struct {
	gorm.Model
	FullName   string `json:"full_name"`
	FirstOrder string `json:"first_order"`
}

type OrderItems struct {
	gorm.Model
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ExpiredAt string  `json:"expired_at"`
}

type CreateOrderHistoriesReq struct {
	UserID       int    `json:"user_id"`
	OrderItemID  int    `json:"order_item_id"`
	Descriptions string `json:"descriptions"`
}

type CreateUserReq struct {
	FullName string `json:"full_name"`
}

type CreateUserRes struct {
	FullName string `json:"full_name"`
}

type GetUsersReq struct {
	Limit int
	Page  int
}

type GetUsersRes struct {
	Limit      int     `json:"limit"`
	Page       int     `json:"page"`
	TotalRows  int     `json:"total_rows"`
	TotalPages int     `json:"total_pages"`
	Users      []Users `json:"users"`
}

type UpdateUserReq struct {
	UserID   int    `json:"user_id"`
	FullName string `json:"full_name"`
}

type UpdateUserRes struct {
	UserID   int    `json:"user_id"`
	FullName string `json:"full_name"`
}

type DeleteUserReq struct {
	UserID int `json:"user_id"`
}

type DeleteUserRes struct {
	UserID   int    `json:"user_id"`
	FullName string `json:"full_name"`
}

type GetUserReq struct {
	UserID int `json:"user_id"`
}

type CreateItemReq struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ExpiredAt string  `json:"expired_at"`
}

type CreateItemRes struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ExpiredAt string  `json:"expired_at"`
}

type UpdateItemReq struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ExpiredAt string  `json:"expired_at"`
}

type UpdateItemRes struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ExpiredAt string  `json:"expired_at"`
}

type GetItemReq struct {
	ID int `json:"id"`
}

type DeleteItemReq struct {
	ID int `json:"id"`
}

type DeleteItemRes struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ExpiredAt string  `json:"expired_at"`
}

type GetItemsReq struct {
	Limit int
	Page  int
}

type GetItemsRes struct {
	Limit      int          `json:"limit"`
	Page       int          `json:"page"`
	TotalRows  int          `json:"total_rows"`
	TotalPages int          `json:"total_pages"`
	OrderItems []OrderItems `json:"order_items"`
}

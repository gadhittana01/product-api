package resthttp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gadhittana01/product-api/services"
	"github.com/labstack/echo"
)

type productHandler struct {
	service ProductService
}

func newProductHandler(service ProductService) *productHandler {
	return &productHandler{
		service: service,
	}
}

const (
	invalidReq        = "Invalid Request Parameter"
	internalServerErr = "Internal Server Error"
)

func (p productHandler) GetOrderHistories(c echo.Context) error {
	logPath := "handler.resthttp.product.GetOrderHistories"
	qLimit := c.QueryParam("limit")
	if qLimit == "" {
		qLimit = "10"
	}
	qPage := c.QueryParam("page")
	if qPage == "" {
		qPage = "1"
	}

	limit, err := strconv.Atoi(qLimit)
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if limit < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "limit must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	page, err := strconv.Atoi(qPage)
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if page < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "page must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.GetOrderHistories(context.Background(), services.GetOrderHistoriesReq{
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) CreateOrderHistories(c echo.Context) error {
	logPath := "handler.resthttp.product.CreateOrderHistories"

	type OrderHistories struct {
		UserID       int    `json:"user_id"`
		OrderItemID  int    `json:"order_item_id"`
		Descriptions string `json:"descriptions"`
	}

	req := OrderHistories{}
	if err := c.Bind(&req); !errors.Is(err, nil) {
		return err
	}

	res, err := p.service.CreateOrderHistories(context.Background(), services.CreateOrderHistoriesReq{
		UserID:       req.UserID,
		OrderItemID:  req.OrderItemID,
		Descriptions: req.Descriptions,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) CreateUser(c echo.Context) error {
	logPath := "handler.resthttp.product.CreateUser"

	type User struct {
		FullName string `json:"full_name"`
	}

	req := User{}
	if err := c.Bind(&req); !errors.Is(err, nil) {
		return err
	}

	if req.FullName == "" {
		log.Println(fmt.Sprintf("%s : %s", logPath, "FullName must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.CreateUser(context.Background(), services.CreateUserReq{
		FullName: req.FullName,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) UpdateUser(c echo.Context) error {
	logPath := "handler.resthttp.product.UpdateUser"

	type User struct {
		FullName string `json:"full_name"`
	}

	req := User{}
	if err := c.Bind(&req); !errors.Is(err, nil) {
		return err
	}

	if req.FullName == "" {
		log.Println(fmt.Sprintf("%s : %s", logPath, "FullName must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if id < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "id must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.UpdateUser(context.Background(), services.UpdateUserReq{
		UserID:   id,
		FullName: req.FullName,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) DeleteUser(c echo.Context) error {
	logPath := "handler.resthttp.product.DeleteUser"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if id < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "id must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.DeleteUser(context.Background(), services.DeleteUserReq{
		UserID: id,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) GetUser(c echo.Context) error {
	logPath := "handler.resthttp.product.GetUser"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if id < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "id must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.GetUser(context.Background(), services.GetUserReq{
		UserID: id,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) GetUsers(c echo.Context) error {
	logPath := "handler.resthttp.product.GetUsers"
	qLimit := c.QueryParam("limit")
	if qLimit == "" {
		qLimit = "10"
	}
	qPage := c.QueryParam("page")
	if qPage == "" {
		qPage = "1"
	}

	limit, err := strconv.Atoi(qLimit)
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if limit < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "limit must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	page, err := strconv.Atoi(qPage)
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if page < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "page must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.GetUsers(context.Background(), services.GetUsersReq{
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) CreateItem(c echo.Context) error {
	logPath := "handler.resthttp.product.CreateItem"

	type Item struct {
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		ExpiredAt string  `json:"expired_at"`
	}

	req := Item{}
	if err := c.Bind(&req); !errors.Is(err, nil) {
		return err
	}

	if req.Name == "" {
		log.Println(fmt.Sprintf("%s : %s", logPath, "Name must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	if req.ExpiredAt == "" {
		log.Println(fmt.Sprintf("%s : %s", logPath, "Expired At must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	if req.Price == 0 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "Price must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.CreateItem(context.Background(), services.CreateItemReq{
		Name:      req.Name,
		Price:     req.Price,
		ExpiredAt: req.ExpiredAt,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) UpdateItem(c echo.Context) error {
	logPath := "handler.resthttp.product.UpdateItem"

	type Item struct {
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		ExpiredAt string  `json:"expired_at"`
	}

	req := Item{}
	if err := c.Bind(&req); !errors.Is(err, nil) {
		return err
	}

	if req.Name == "" {
		log.Println(fmt.Sprintf("%s : %s", logPath, "Name must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	if req.Price == 0 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "Price must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	if req.ExpiredAt == "" {
		log.Println(fmt.Sprintf("%s : %s", logPath, "Expired at must be fill"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if id < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "id must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.UpdateItem(context.Background(), services.UpdateItemReq{
		ID:        id,
		Name:      req.Name,
		Price:     req.Price,
		ExpiredAt: req.ExpiredAt,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) GetItem(c echo.Context) error {
	logPath := "handler.resthttp.product.GetItem"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if id < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "id must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.GetItem(context.Background(), services.GetItemReq{
		ID: id,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) GetItems(c echo.Context) error {
	logPath := "handler.resthttp.product.GetItems"
	qLimit := c.QueryParam("limit")
	if qLimit == "" {
		qLimit = "10"
	}
	qPage := c.QueryParam("page")
	if qPage == "" {
		qPage = "1"
	}

	limit, err := strconv.Atoi(qLimit)
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if limit < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "limit must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	page, err := strconv.Atoi(qPage)
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if page < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "page must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.GetItems(context.Background(), services.GetItemsReq{
		Limit: limit,
		Page:  page,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (p productHandler) DeleteItem(c echo.Context) error {
	logPath := "handler.resthttp.product.DeleteItem"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}
	if id < 1 {
		log.Println(fmt.Sprintf("%s : %s", logPath, "id must greater than 0"))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": invalidReq,
		})
	}

	res, err := p.service.DeleteItem(context.Background(), services.DeleteItemReq{
		ID: id,
	})
	if err != nil {
		log.Println(fmt.Sprintf("%s : %s", logPath, err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": internalServerErr,
		})
	}

	return c.JSON(http.StatusOK, res)
}

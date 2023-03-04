package resthttp

import (
	"github.com/labstack/echo"
)

type RouterDependencies struct {
	PS ProductService
}

func NewRoutes(rd RouterDependencies) *echo.Echo {
	router := echo.New()

	ph := newProductHandler(rd.PS)

	// order histories
	router.GET("/get-order-histories", ph.GetOrderHistories)
	router.POST("/create-order-histories", ph.CreateOrderHistories)

	// user
	router.POST("/create-user", ph.CreateUser)
	router.PUT("/update-user/:id", ph.UpdateUser)
	router.DELETE("/delete-user/:id", ph.DeleteUser)
	router.GET("/get-users", ph.GetUsers)
	router.GET("/get-user/:id", ph.GetUser)

	// items
	router.POST("/create-item", ph.CreateItem)
	router.PUT("/update-item/:id", ph.UpdateItem)
	router.DELETE("/delete-item/:id", ph.DeleteItem)
	router.GET("/get-items", ph.GetItems)
	router.GET("/get-item/:id", ph.GetItem)

	return router
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gadhittana01/product-api/config"
	"github.com/gadhittana01/product-api/db"
	"github.com/gadhittana01/product-api/handler/resthttp"
	"github.com/gadhittana01/product-api/pkg/product"
	"github.com/gadhittana01/product-api/services"
	"github.com/go-redis/redis/v8"
)

func initApp(c *config.GlobalConfig) error {
	db := db.InitDB()
	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
	})
	if err := redis.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redis connection failed, err : %v", err)
		return err
	}
	productPkg := product.New(db, redis, c)

	ps, err := services.NewProductService(services.ProductDependencies{
		PR: productPkg,
	})
	if err != nil {
		return err
	}

	return startHTTPServer(resthttp.NewRoutes(resthttp.RouterDependencies{
		PS: ps,
	}), c)
}

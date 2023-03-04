package product

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gadhittana01/product-api/config"
	"github.com/gadhittana01/product-api/pkg/domain"
	"github.com/go-redis/redis/v8"
)

type cache struct {
	redis  *redis.Client
	config *config.GlobalConfig
}

type secondary interface {
	getUserCache(ctx context.Context, cacheKey string) (domain.Users, error)
	setUserCache(ctx context.Context, cacheKey string, data domain.Users) error
	getItemCache(ctx context.Context, cacheKey string) (domain.OrderItems, error)
	setItemCache(ctx context.Context, cacheKey string, data domain.OrderItems) error
}

func newCache(rd *redis.Client, cfg *config.GlobalConfig) secondary {
	return &cache{
		redis:  rd,
		config: cfg,
	}
}

func (ch *cache) getUserCache(ctx context.Context, cacheKey string) (domain.Users, error) {
	var users domain.Users
	res := ch.redis.Get(context.Background(), cacheKey)
	if err := res.Err(); err != nil {
		log.Printf("unable to GET data. error: %v", err)
		return users, err
	}
	resString, err := res.Result()
	if err != nil {
		log.Printf("unable to GET data. error: %v", err)
		return users, err
	}
	if err := json.Unmarshal([]byte(resString), &users); err != nil {
		return users, err
	}
	return users, nil
}

func (ch *cache) setUserCache(ctx context.Context, cacheKey string, data domain.Users) error {
	req, err := json.Marshal(data)
	if err != nil {
		log.Printf("unable to Unmarshal data. error: %v", err)
		return err
	}
	res := ch.redis.Set(context.Background(), cacheKey, string(req), time.Duration(ch.config.Redis.TTL)*time.Second)
	if err := res.Err(); err != nil {
		log.Printf("unable to SET data. error: %v", err)
		return err
	}
	return nil
}

func (ch *cache) getItemCache(ctx context.Context, cacheKey string) (domain.OrderItems, error) {
	var item domain.OrderItems
	res := ch.redis.Get(context.Background(), cacheKey)
	if err := res.Err(); err != nil {
		log.Printf("unable to GET data. error: %v", err)
		return item, err
	}
	resString, err := res.Result()
	if err != nil {
		log.Printf("unable to GET data. error: %v", err)
		return item, err
	}
	if err := json.Unmarshal([]byte(resString), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (ch *cache) setItemCache(ctx context.Context, cacheKey string, data domain.OrderItems) error {
	req, err := json.Marshal(data)
	if err != nil {
		log.Printf("unable to Unmarshal data. error: %v", err)
		return err
	}
	res := ch.redis.Set(context.Background(), cacheKey, string(req), time.Duration(ch.config.Redis.TTL)*time.Second)
	if err := res.Err(); err != nil {
		log.Printf("unable to SET data. error: %v", err)
		return err
	}
	return nil
}

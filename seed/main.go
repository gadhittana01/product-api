package main

import (
	"errors"
	"log"

	"github.com/gadhittana01/product-api/db"
	"github.com/gadhittana01/product-api/pkg/domain"
	"github.com/jaswdr/faker"
)

func seed(n int) {
	fake := faker.New()
	db := db.InitDB()

	defer db.Close()

	for i := 0; i < n; i++ {
		user := domain.Users{
			FullName:   fake.Person().Name(),
			FirstOrder: "2022-01-01",
		}
		if err := db.Create(&user).Error; !errors.Is(err, nil) {
			log.Fatal(err)
		}

		ordersItem := domain.OrderItems{
			Name:      fake.Beer().Name(),
			Price:     fake.Payment().Faker.Float64(2, 1, 200),
			ExpiredAt: "2022-01-01",
		}
		if err := db.Create(&ordersItem).Error; !errors.Is(err, nil) {
			log.Fatal(err)
		}

		orderHistories := domain.OrderHistories{
			UserID:       1,
			OrderItemID:  1,
			Descriptions: "User with id 1 order item with id 1",
			CreatedAt:    "2022-01-01",
			UpdatedAt:    "2022-01-01",
		}
		if err := db.Create(&orderHistories).Error; !errors.Is(err, nil) {
			log.Fatal(err)
		}
	}
}

func main() {
	n := 10
	seed(n)
}

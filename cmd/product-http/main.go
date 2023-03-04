package main

import (
	"log"

	"github.com/gadhittana01/product-api/config"
	"github.com/gadhittana01/product-api/helper"
)

func main() {
	f := helper.InitLog()
	defer f.Close()

	config := &config.GlobalConfig{}
	helper.LoadConfig(config)
	err := initApp(config)
	if err != nil {
		log.Println(err)
	}
}

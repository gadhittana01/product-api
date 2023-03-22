package helper

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/gadhittana01/product-api/config"
	"gopkg.in/yaml.v2"
)

func LoadConfig(c *config.GlobalConfig) {
	path := "config/product-http.yaml"
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	replaceWithENVValue(c)

}

func replaceWithENVValue(c *config.GlobalConfig) {
	// DB
	DBHost := os.Getenv("DB_HOST")
	if DBHost != "" {
		c.DB.Host = DBHost
	}

	DBPort := os.Getenv("DB_PORT")
	if DBPort != "" {
		port, err := strconv.Atoi(DBPort)
		if err != nil {
			log.Fatalf("Error Parse : %v", err)
		}
		c.DB.Port = int32(port)
	}

	DBUser := os.Getenv("DB_USER")
	if DBUser != "" {
		c.DB.User = DBUser
	}

	DBPassword := os.Getenv("DB_PASSWORD")
	if DBPassword != "" {
		c.DB.Password = DBPassword
	}

	DBName := os.Getenv("DB_NAME")
	if DBName != "" {
		c.DB.Name = DBName
	}

	// REDIS
	RedisHost := os.Getenv("REDIS_HOST")
	if RedisHost != "" {
		c.Redis.Host = RedisHost
	}

	RedisPort := os.Getenv("REDIS_PORT")
	if RedisPort != "" {
		port, err := strconv.Atoi(RedisPort)
		if err != nil {
			log.Fatalf("Error Parse : %v", err)
		}
		c.Redis.Port = int32(port)
	}

	RedisPassword := os.Getenv("REDIS_PASSWORD")
	if RedisPassword != "" {
		c.Redis.Password = RedisPassword
	}
}

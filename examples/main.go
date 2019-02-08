package main

import (
	"webservice-library/examples/route"
	libroute "webservice-library/route"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

func main() {
	router := gin.Default()

	sr := &libroute.Resource{
		DB:          &gorm.DB{},
		Validator:   &validator.Validate{},
		RedisClient: &redis.Client{},
	}

	for _, v := range route.InitRoute() {
		router.GET("/"+v.Name, v.Register("GetAll", sr))
	}

	router.Run(":9090")
}

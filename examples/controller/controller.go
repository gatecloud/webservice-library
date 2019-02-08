package controller

import (
	"webservice-library/controller"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

// Controller is a base controller for new projects
type Controller struct {
	controller.BaseControl
}

// Init inits the Control data
func (ctrl *Controller) Init(db *gorm.DB, validate *validator.Validate, model interface{}, redis *redis.Client) {
	ctrl.DB = db
	ctrl.Validate = validate
	ctrl.Model = model
	ctrl.RedisClient = redis
}

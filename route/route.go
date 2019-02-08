package route

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

// Router is a common interface to proceed the router's operation
type Router interface {
	Register(action string, sr *Resource) func(ctx *gin.Context)
}

// Route stores a single route entity's information
type Route struct {
	Name       string
	Controller reflect.Type
	Model      interface{}
}

// Resource representing the shared resource
type Resource struct {
	DB          *gorm.DB
	Validator   *validator.Validate
	RedisClient *redis.Client
}

// Register registers the API to route
func (r *Route) Register(action string, sr *Resource) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ptr := reflect.New(r.Controller)
		methodInit := ptr.MethodByName("Init")
		if methodInit.IsValid() {
			args := make([]reflect.Value, 4)
			args[0] = reflect.ValueOf(sr.DB)
			args[1] = reflect.ValueOf(sr.Validator)
			args[2] = reflect.ValueOf(r.Model)
			args[3] = reflect.ValueOf(sr.RedisClient)
			methodInit.Call(args)
		}

		m := ptr.MethodByName(action)
		if m.IsValid() {
			args := make([]reflect.Value, 1)
			args[0] = reflect.ValueOf(ctx)
			m.Call(args)
		}
	}
}

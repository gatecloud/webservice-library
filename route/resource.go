package route

import (
	"errors"
	"reflect"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

// Resourcer is a common interface to allow to add custom resources
type Resourcer interface {
	Resources() ([]reflect.Value, error)
}

// Resource representing the shared resource
type Resource struct {
	DB          *gorm.DB            `gatecloud:"initparam"`
	Validator   *validator.Validate `gatecloud:"initparam"`
	RedisClient *redis.Client       `gatecloud:"initparam"`
}

// Resources returns all resources that tag as initparam
func (r *Resource) Resources() ([]reflect.Value, error) {
	rf := reflect.ValueOf(r).Elem()
	if !rf.IsValid() {
		return []reflect.Value{}, errors.New("reflect error")
	}

	args := make([]reflect.Value, rf.NumField())
	for i := 0; i < rf.NumField(); i++ {
		if rf.Type().Field(i).Tag.Get("gatecloud") == "initparam" {
			args[i] = rf.Field(i)
		}
	}
	return args, nil
}

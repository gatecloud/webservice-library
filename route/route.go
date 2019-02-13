package route

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
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

// Register registers the API to route
func (r Route) Register(action string, resourcer Resourcer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ptr := reflect.New(r.Controller)
		methodInit := ptr.MethodByName("Init")
		if methodInit.IsValid() {
			resources, err := resourcer.Resources()
			if err != nil {
				log.Fatal(err)
			}

			resCount := len(resources)
			args := make([]reflect.Value, resCount+1)
			for i, v := range resources {
				args[i] = v
			}
			args[resCount] = reflect.ValueOf(r.Model)
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

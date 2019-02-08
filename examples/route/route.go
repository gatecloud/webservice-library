package route

import (
	"reflect"
	"webservice-library/examples/controller"
	"webservice-library/examples/models"
	"webservice-library/route"
)

// InitRoute inits the route group
func InitRoute() []route.Route {
	r := []route.Route{
		route.Route{
			Name:       "Users",
			Controller: reflect.TypeOf(controller.UserController{}),
			Model:      &models.User{},
		},
	}
	return r
}

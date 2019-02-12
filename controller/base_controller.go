package controller

import (
	"net/http"

	"github.com/gatecloud/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

// Controller is an control interface based on go-gin
type Controller interface {
	Init(...interface{})
	Prepare()
	Post(ctx *gin.Context)
	Patch(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Head(ctx *gin.Context)
	Options(ctx *gin.Context)
}

// BaseControl representing the structure of a RESTful API handler
type BaseControl struct {
	DB          *gorm.DB
	Validator   *validator.Validate
	RedisClient *redis.Client
	Model       interface{}
}

// Init inits the Control data
func (ctrl *BaseControl) Init(db *gorm.DB, validate *validator.Validate, redis *redis.Client, model interface{}) {
	ctrl.DB = db
	ctrl.Validator = validate
	ctrl.RedisClient = redis
	ctrl.Model = model
}

// Prepare corresponds http Prepare method
func (ctrl *BaseControl) Prepare() {}

// Post corresponds http Post method
func (ctrl *BaseControl) Post(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Patch corresponds http Patch method
func (ctrl *BaseControl) Patch(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Delete corresponds http Delete method
func (ctrl *BaseControl) Delete(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// GetByID corresponds http Get :id method
func (ctrl *BaseControl) GetByID(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// GetAll corresponds http Get method
func (ctrl *BaseControl) GetAll(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Head corresponds http Head method
func (ctrl *BaseControl) Head(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Options corresponds http Options method
func (ctrl *BaseControl) Options(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
	ctx.Abort()
	return
}

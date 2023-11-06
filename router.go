package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"captcha/api"
	"captcha/middelware"
	"captcha/service"
	"captcha/util"
)

func Newrouter(database *mongo.Database, jwt *util.JWT) *gin.Engine {
	router := gin.Default()
	captcha := util.Newcaptcha()
	svc := service.Newservice(database)
	api := api.Newapi(svc, captcha.Captcha, jwt)
	registerRoutes(router, api, jwt)
	return router
}
func registerRoutes(router *gin.Engine, api *api.Api, jwt *util.JWT) {

	rg := router.Group("/jwt", middelware.HTTPMiddlewareJWT(jwt))
	rg.GET("/getcaptcha", api.GetCaptcha)
}

package router

import (
	"gin-framework-test/basic-api/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine

	healthController controllers.HealthController
}

func NewRouter() *Router {
	r := &Router{}
	r.engine = gin.Default()
	r.healthController = controllers.NewHealthController()

	return r
}

func (r *Router) Run() {
	r.engine.Run()
}
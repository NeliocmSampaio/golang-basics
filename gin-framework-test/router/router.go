package router

import (
	"gin-framework-test/basic-api/controllers"
	"gin-framework-test/basic-api/services"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine

	healthController controllers.HealthController
	bookController   controllers.BookController
}

func NewRouter() *Router {
	r := &Router{}
	r.engine = gin.Default()

	// TODO: Dependency Injection
	// services
	bookService := services.NewBookService()

	r.bookController = controllers.NewBookController(bookService)
	r.healthController = controllers.NewHealthController()

	return r
}

func (r *Router) Run() {
	r.engine.Run()
}

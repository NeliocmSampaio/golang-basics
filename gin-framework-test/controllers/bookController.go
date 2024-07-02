package controllers

import (
	"gin-framework-test/basic-api/domain"
	"gin-framework-test/basic-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) BookController {
	return BookController{
		bookService: bookService,
	}
}

func (c *BookController) HandlePostBook(ctx *gin.Context) {

	book := domain.Book{
		Name: "test",
	}

	err := c.bookService.Save(book)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusCreated)
}

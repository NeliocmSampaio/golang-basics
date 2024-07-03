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
		Name:   "name of the book",
		Author: "author name",
		Price:  5.99,
	}

	err := c.bookService.Save(book)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusCreated)
}

func (c *BookController) HandleGetBooks(ctx *gin.Context) {
	books, err := c.bookService.GetBooks()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, books)
}

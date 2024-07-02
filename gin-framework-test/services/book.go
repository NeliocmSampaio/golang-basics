package services

import (
	"fmt"
	"gin-framework-test/basic-api/domain"
)

type BookService interface {
	Save(book domain.Book) error
}

type bookService struct {
}

func NewBookService() BookService {
	return &bookService{}
}

func (s *bookService) Save(book domain.Book) error {
	fmt.Println("Book Created!")
	return nil
}

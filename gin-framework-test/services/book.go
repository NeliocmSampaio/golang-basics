package services

import (
	"database/sql"
	"fmt"
	"gin-framework-test/basic-api/domain"
)

type BookService interface {
	Save(book domain.Book) error
	GetBooks() (books []domain.Book, err error)
}

type bookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) BookService {
	return &bookService{
		db: db,
	}
}

func (s *bookService) Save(book domain.Book) error {
	query := `
		INSERT INTO tab_books (title, author, price)
		VALUES (?, ?, ?)
		;
	`

	result, err := s.db.Exec(query,
		book.Name,
		book.Author,
		book.Price)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println(id)

	fmt.Println("Book Created!")
	return nil
}

func (s *bookService) GetBooks() (books []domain.Book, err error) {

	query := `
		SELECT * FROM tab_books;
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Price); err != nil {
			return books, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return books, err
	}

	return books, nil
}

package service

import (
	"errors"

	"github.com/kevindiu/gotest2/example/model"
	"github.com/kevindiu/gotest2/example/repository"
	"github.com/kevindiu/gotest2/example/utils"
)

type BookService struct {
	repo repository.Repository[model.Book, string]
}

func NewBookService(repo repository.Repository[model.Book, string]) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(title, author, isbn string) (model.Book, error) {
	// Validate ISBN
	if !utils.ParseISBN(isbn) {
		return model.Book{}, errors.New("invalid ISBN")
	}

	formattedISBN := utils.FormatISBN(isbn)

	book := model.Book{
		ID:     formattedISBN, // Using ISBN as ID for simplicity
		Title:  title,
		Author: author,
		ISBN:   formattedISBN,
	}

	if err := s.repo.Create(book.ID, book); err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (s *BookService) GetBook(id string) (model.Book, error) {
	return s.repo.Get(id)
}

func (s *BookService) ListBooks() ([]model.Book, error) {
	return s.repo.List()
}

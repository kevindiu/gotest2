package service

import (
	"context"
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

// BatchCreate creates multiple books concurrently using channels.
func (s *BookService) BatchCreate(ctx context.Context, books []model.Book) <-chan error {
	errChan := make(chan error, len(books))
	go func() {
		defer close(errChan)
		for _, book := range books {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				if err := s.repo.Create(book.ID, book); err != nil {
					errChan <- err
				}
			}
		}
	}()
	return errChan
}

// validateBook is an unexported helper method.
func (s *BookService) validateBook(b model.Book) error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Author == "" {
		return errors.New("author is required")
	}
	return nil
}

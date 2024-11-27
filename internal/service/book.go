package service

import (
	"context"
	"time"

	"github.com/qooodr/CRUD-APP/internal/domain"
)

type BooksRepository interface {
	Create(ctx context.Context, book domain.Book) error
	GetById(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, imp domain.UpdateBookInput) error
}

type Books struct {
	reps BooksRepository
}

func NewBooks(reps BooksRepository) *Books {
	return &Books{
		reps: reps,
	}
}

func (b *Books) Create(ctx context.Context, book domain.Book) error {
	if book.PublishDate.IsZero() {
		book.PublishDate = time.Now()
	}

	return b.reps.Create(ctx, book)
}

func (b *Books) GetById(ctx context.Context, id int64) (domain.Book, error) {
	return b.reps.GetById(ctx, id)
}

func (s *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	return s.reps.GetAll(ctx)
}

func (b *Books) Delete(ctx context.Context, id int64) error {
	return b.reps.Delete(ctx, id)
}

func (b *Books) Update(ctx context.Context, id int64, inp domain.UpdateBookInput) error {
	return b.reps.Update(ctx, id, inp)
}

package domain

import (
	"context"
	"errors"
	"time"
)

type MenuID int64

type MenuItem struct {
	ID          MenuID
	Name        string
	Description string
	Price       float64
	CategoryID  int
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMenuItem(name, description string, price float64, categoryID int, imageURL string) (*MenuItem, error) {
	if name == "" {
		return nil, errors.New("название не может быть пустым")
	}
	if price < 0 {
		return nil, errors.New("цена не может быть отрицательной")
	}
	if categoryID <= 0 {
		return nil, errors.New("некорректный ID категории")
	}

	return &MenuItem{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// internal/domain/menu.go
type MenuRepository interface {
	Save(ctx context.Context, item *MenuItem) error
	FindAll(ctx context.Context) ([]*MenuItem, error)
	FindByCategory(ctx context.Context, categoryID int) ([]*MenuItem, error)
	FindByID(ctx context.Context, id MenuID) (*MenuItem, error)
}

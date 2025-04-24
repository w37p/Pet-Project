package menu

import (
	"context"
	"fmt"

	"github.com/bullockz21/pet_project21/internal/modules/menu/domain"
	"github.com/bullockz21/pet_project21/internal/modules/menu/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MenuRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{
		db:  db,
		log: logrus.New(),
	}
}

func (r *MenuRepository) domainToEntity(m *domain.MenuItem) *entity.MenuItem {
	return &entity.MenuItem{
		ID:          uint(m.ID),
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		CategoryID:  m.CategoryID,
		ImageURL:    m.ImageURL,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func (r *MenuRepository) entityToDomain(e *entity.MenuItem) *domain.MenuItem {
	return &domain.MenuItem{
		ID:          domain.MenuID(e.ID),
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		CategoryID:  e.CategoryID,
		ImageURL:    e.ImageURL,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func (r *MenuRepository) Save(ctx context.Context, item *domain.MenuItem) error {
	eItem := r.domainToEntity(item)
	result := r.db.WithContext(ctx).Save(eItem)
	if result.Error != nil {
		r.log.Errorf("Ошибка сохранения меню: %v", result.Error)
		return fmt.Errorf("ошибка сохранения меню: %w", result.Error)
	}
	return nil
}

func (r *MenuRepository) FindAll(ctx context.Context) ([]*domain.MenuItem, error) {
	var items []entity.MenuItem
	result := r.db.WithContext(ctx).Find(&items)
	if result.Error != nil {
		return nil, fmt.Errorf("ошибка получения меню: %w", result.Error)
	}

	domainItems := make([]*domain.MenuItem, len(items))
	for i, item := range items {
		domainItems[i] = r.entityToDomain(&item)
	}
	return domainItems, nil
}

// internal/modules/repository/menu/menu_repository.go

func (r *MenuRepository) FindByCategory(ctx context.Context, categoryID int) ([]*domain.MenuItem, error) {
	var dbItems []*entity.MenuItem
	result := r.db.WithContext(ctx).Where("category_id = ?", categoryID).Find(&dbItems)
	if result.Error != nil {
		return nil, result.Error
	}

	domainItems := make([]*domain.MenuItem, len(dbItems))
	for i, item := range dbItems {
		domainItems[i] = r.entityToDomain(item)
	}

	return domainItems, nil
}
func (r *MenuRepository) FindByID(ctx context.Context, id domain.MenuID) (*domain.MenuItem, error) {
	var item entity.MenuItem
	result := r.db.WithContext(ctx).First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.entityToDomain(&item), nil
}

// Реализуйте остальные методы аналогично

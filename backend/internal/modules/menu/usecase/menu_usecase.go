package usecase

import (
	"context"

	"github.com/bullockz21/pet_project21/internal/modules/menu/domain"
)

type MenuUseCase struct {
	repo domain.MenuRepository
}

func NewMenuUseCase(repo domain.MenuRepository) *MenuUseCase {
	return &MenuUseCase{repo: repo}
}

func (uc *MenuUseCase) GetFullMenu(ctx context.Context) ([]*domain.MenuItem, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *MenuUseCase) GetByCategory(ctx context.Context, categoryID int) ([]*domain.MenuItem, error) {
	return uc.repo.FindByCategory(ctx, categoryID)
}

// internal/modules/presenter/menu/menu_presenter.go
package menu

import "github.com/bullockz21/pet_project21/internal/modules/domain"

type MenuPresenter struct{}

func NewMenuPresenter() *MenuPresenter {
	return &MenuPresenter{}
}

type MenuItemResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}

func (p *MenuPresenter) Present(items []*domain.MenuItem) []MenuItemResponse {
	response := make([]MenuItemResponse, len(items))
	for i, item := range items {
		response[i] = MenuItemResponse{
			ID:          int64(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			ImageURL:    item.ImageURL,
		}
	}
	return response
}

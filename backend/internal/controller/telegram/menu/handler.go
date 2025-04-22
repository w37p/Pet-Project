package menu

import (
	"net/http"

	presenter "github.com/bullockz21/pet_project21/internal/modules/presenter/menu"
	usecase "github.com/bullockz21/pet_project21/internal/modules/usecase/menu"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	usecase   *usecase.MenuUseCase
	presenter *presenter.MenuPresenter
}

func NewMenuController(uc *usecase.MenuUseCase, p *presenter.MenuPresenter) *MenuController {
	return &MenuController{
		usecase:   uc,
		presenter: p,
	}
}

// GetMenu godoc
// @Summary Get full menu
// @Description Get complete restaurant menu
// @Tags Menu
// @Produce json
// @Success 200 {array} presenter.MenuItemResponse
// @Router /api/v1/menu [get]
func (c *MenuController) GetMenu(ctx *gin.Context) {
	items, err := c.usecase.GetFullMenu(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get menu"})
		return
	}

	response := c.presenter.PresentMenu(items)
	ctx.JSON(http.StatusOK, response)
}

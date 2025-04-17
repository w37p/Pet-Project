// internal/controller/telegram/start_handler.go
package bot

import (
	"context"
	"log"

	"github.com/bullockz21/pet_project21/configs"
	presenterUser "github.com/bullockz21/pet_project21/internal/modules/presenter/user"
	usecaseUser "github.com/bullockz21/pet_project21/internal/modules/usecase/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	userUC        *usecaseUser.UserUseCase
	userPresenter *presenterUser.UserPresenter
	config        *configs.Config
}

func NewStartHandler(userUC *usecaseUser.UserUseCase, userPresenter *presenterUser.UserPresenter, cfg *configs.Config) *StartHandler {
	return &StartHandler{
		userUC:        userUC,
		userPresenter: userPresenter,
		config:        cfg,
	}
}

// HandleStart обрабатывает команду /start.
func (h *StartHandler) HandleStart(ctx context.Context, update tgbotapi.Update) {
	telegramID := update.Message.From.ID
	username := update.Message.From.UserName
	firstName := update.Message.From.FirstName
	language := update.Message.From.LanguageCode

	// Вызываем usecase для создания пользователя
	if _, err := h.userUC.CreateUser(ctx, telegramID, username, firstName, language); err != nil {
		log.Printf("[ERROR] HandleStart: %v", err)
		h.userPresenter.PresentError(update.Message.Chat.ID, "Не удалось создать пользователя")
		return
	}

	// Используем переданную конфигурацию для получения URL мини‑аппа.
	miniAppURL := h.config.Telegram.WebhookURL

	// Отправляем приветственное сообщение с кнопкой для открытия мини‑аппа
	if err := h.userPresenter.PresentWelcomeMessage(update.Message.Chat.ID, firstName, miniAppURL); err != nil {
		log.Printf("Ошибка отправки приветственного сообщения: %v", err)
	}
}

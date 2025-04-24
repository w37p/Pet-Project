package bot

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/bullockz21/pet_project21/configs"
	presenterUser "github.com/bullockz21/pet_project21/internal/modules/user/presenter"
	usecaseUser "github.com/bullockz21/pet_project21/internal/modules/user/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	userUC        *usecaseUser.UserUseCase
	userPresenter *presenterUser.UserPresenter
	config        *configs.Config
	logger        *logrus.Logger
}

func NewStartHandler(userUC *usecaseUser.UserUseCase, userPresenter *presenterUser.UserPresenter, cfg *configs.Config) *StartHandler {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	return &StartHandler{
		userUC:        userUC,
		userPresenter: userPresenter,
		config:        cfg,
		logger:        logger,
	}
}

// HandleStart обрабатывает команду /start.
func (h *StartHandler) HandleStart(ctx context.Context, update tgbotapi.Update) {
	telegramID := update.Message.From.ID
	username := update.Message.From.UserName
	firstName := update.Message.From.FirstName
	language := update.Message.From.LanguageCode

	h.logger.WithFields(logrus.Fields{
		"telegram_id": telegramID,
		"username":    username,
		"first_name":  firstName,
		"language":    language,
	}).Info("Обработка команды /start")

	// Создание пользователя
	if _, err := h.userUC.CreateUser(ctx, telegramID, username, firstName, language); err != nil {
		h.logger.WithError(err).WithField("telegram_id", telegramID).Error("Не удалось создать пользователя")
		h.userPresenter.PresentError(update.Message.Chat.ID, "Не удалось создать пользователя")
		return
	}

	miniAppURL := h.config.Telegram.WebhookURL
	h.logger.WithField("mini_app_url", miniAppURL).Info("Отправка приветственного сообщения с кнопкой MiniApp")

	if err := h.userPresenter.PresentWelcomeMessage(update.Message.Chat.ID, firstName, miniAppURL); err != nil {
		h.logger.WithError(err).WithField("chat_id", update.Message.Chat.ID).Error("Ошибка отправки приветственного сообщения")
	}
}

package bot

import (
	"context"

	presenterUser "github.com/bullockz21/pet_project21/internal/modules/user/presenter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type CommandHandler struct {
	startHandler  *StartHandler
	userPresenter *presenterUser.UserPresenter
	logger        *logrus.Logger
}

func NewCommandHandler(startHandler *StartHandler, userPresenter *presenterUser.UserPresenter) *CommandHandler {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	return &CommandHandler{
		startHandler:  startHandler,
		userPresenter: userPresenter,
		logger:        logger,
	}
}

func (h *CommandHandler) HandleCommand(ctx context.Context, update tgbotapi.Update) {
	command := update.Message.Command()
	chatID := update.Message.Chat.ID

	h.logger.WithFields(logrus.Fields{
		"chat_id": chatID,
		"command": command,
	}).Info("Received command")

	switch command {
	case "start":
		h.logger.WithField("chat_id", chatID).Info("Handling /start command")
		h.startHandler.HandleStart(ctx, update)
	default:
		h.logger.WithFields(logrus.Fields{
			"chat_id": chatID,
			"command": command,
		}).Warn("Unknown command received")
		h.userPresenter.PresentError(chatID, "Неизвестная команда.")
	}
}

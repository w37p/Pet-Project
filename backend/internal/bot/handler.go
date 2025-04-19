package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	bot             *tgbotapi.BotAPI
	commandHandler  *CommandHandler
	callbackHandler *CallbackHandler
	logger          *logrus.Logger
}

func NewHandler(bot *tgbotapi.BotAPI, commandHandler *CommandHandler, callbackHandler *CallbackHandler) *Handler {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	return &Handler{
		bot:             bot,
		commandHandler:  commandHandler,
		callbackHandler: callbackHandler,
		logger:          logger,
	}
}

func (h *Handler) ProcessUpdate(ctx context.Context, update tgbotapi.Update) {
	switch {
	case update.Message != nil && update.Message.IsCommand():
		h.logger.WithFields(logrus.Fields{
			"type":    "command",
			"chat_id": update.Message.Chat.ID,
			"command": update.Message.Command(),
		}).Info("Processing command update")

		h.commandHandler.HandleCommand(ctx, update)

	case update.CallbackQuery != nil:
		h.logger.WithFields(logrus.Fields{
			"type":     "callback_query",
			"chat_id":  update.CallbackQuery.Message.Chat.ID,
			"callback": update.CallbackQuery.Data,
		}).Info("Processing callback query")

		h.callbackHandler.HandleCallback(ctx, update)

	default:
		h.logger.WithField("update_id", update.UpdateID).Warn("Received unsupported update type")
	}
}

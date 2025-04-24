package bot

import (
	"context"

	buttons "github.com/bullockz21/pet_project21/internal/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type CallbackHandler struct {
	bot *tgbotapi.BotAPI
}

func NewCallbackHandler(bot *tgbotapi.BotAPI) *CallbackHandler {
	return &CallbackHandler{bot: bot}
}

func (h *CallbackHandler) HandleCallback(ctx context.Context, update tgbotapi.Update) {
	// Настроим логгер logrus
	log := logrus.New()

	callback := update.CallbackQuery
	chatID := callback.Message.Chat.ID
	data := callback.Data

	// Логирование входящего обновления
	log.WithFields(logrus.Fields{
		"chat_id": chatID,
		"data":    data,
	}).Info("Handling callback")

	var text string
	switch data {
	case buttons.MenuButton.Data:
		newKeyboard := buttons.InlineKeyboardColumn(buttons.ShawarmaButton, buttons.DrinksButton, buttons.DessertsButton, buttons.BackButton)
		edit := tgbotapi.NewEditMessageReplyMarkup(chatID, callback.Message.MessageID, newKeyboard)
		if _, err := h.bot.Send(edit); err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Failed to update keyboard")
		} else {
			log.Info("Keyboard updated successfully")
		}

	case buttons.PromotionsButton.Data:
		text = "🔥 Актуальные акции:"
		log.Info("Showing promotions")
	case buttons.ReviewsButton.Data:
		text = "⭐ Отзывы наших клиентов:"
		log.Info("Showing reviews")
	default:
		text = "Неизвестная кнопка"
		log.Warn("Unknown button clicked")
	}

	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := h.bot.Send(msg); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to send message")
	} else {
		log.Info("Message sent successfully")
	}

	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := h.bot.Request(callbackConfig); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to send callback response")
	} else {
		log.Info("Callback response sent successfully")
	}
}

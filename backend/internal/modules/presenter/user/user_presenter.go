package user

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type UserPresenter struct {
	bot    *tgbotapi.BotAPI
	logger *logrus.Entry
}

func NewUserPresenter(bot *tgbotapi.BotAPI) *UserPresenter {
	return &UserPresenter{
		bot:    bot,
		logger: logrus.WithField("component", "UserPresenter"),
	}
}

// PresentError отправляет сообщение об ошибке.
func (p *UserPresenter) PresentError(chatID int64, errorMsg string) error {
	msgText := fmt.Sprintf("🚫 Ошибка: %s", errorMsg)
	msg := tgbotapi.NewMessage(chatID, msgText)

	p.logger.WithFields(logrus.Fields{
		"chatID":   chatID,
		"errorMsg": errorMsg,
	}).Info("Отправка сообщения об ошибке")

	_, err := p.bot.Send(msg)
	if err != nil {
		p.logger.WithError(err).Error("Не удалось отправить сообщение об ошибке")
	}
	return err
}

// PresentWelcomeMessage отправляет приветственное сообщение с Web App кнопкой.
func (p *UserPresenter) PresentWelcomeMessage(chatID int64, firstName, miniAppURL string) error {
	text := fmt.Sprintf("Привет, %s! На связи служба доставки \"Рыба и Рис\".\nНажмите кнопку ниже, чтобы сделать заказ.", firstName)
	webAppButton := tgbotapi.NewInlineKeyboardButtonWebApp("Сделать заказ", tgbotapi.WebAppInfo{
		URL: miniAppURL,
	})
	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(webAppButton))
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard

	p.logger.WithFields(logrus.Fields{
		"chatID":     chatID,
		"firstName":  firstName,
		"miniAppURL": miniAppURL,
	}).Info("Отправка приветственного сообщения")

	_, err := p.bot.Send(msg)
	if err != nil {
		p.logger.WithError(err).Error("Не удалось отправить приветственное сообщение")
	}
	return err
}

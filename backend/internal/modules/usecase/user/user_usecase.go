package user

import (
	"context"
	"fmt"

	"github.com/bullockz21/pet_project21/internal/modules/domain"
	"github.com/sirupsen/logrus"
)

type UserUseCase struct {
	repo domain.Repository
	log  *logrus.Logger
}

func NewUserUseCase(repo domain.Repository) *UserUseCase {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	return &UserUseCase{repo: repo, log: log}
}

// CreateUser создает пользователя, используя доменную логику, и сохраняет его через репозиторий.
func (uc *UserUseCase) CreateUser(ctx context.Context, telegramID int64, username, firstName, language string) (*domain.User, error) {
	uc.log.WithFields(logrus.Fields{
		"telegram_id": telegramID,
		"username":    username,
		"first_name":  firstName,
		"language":    language,
	}).Info("Создание пользователя")

	// Создаем доменного пользователя с 4 аргументами.
	user, err := domain.NewUser(telegramID, username, firstName, language)
	if err != nil {
		uc.log.WithError(err).Error("Ошибка при создании пользователя")
		return nil, err
	}

	// Сохраняем пользователя через репозиторий.
	if err := uc.repo.Save(ctx, user); err != nil {
		uc.log.WithError(err).Error("Ошибка при сохранении пользователя")
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	uc.log.WithField("telegram_id", user.TelegramID).Info("Пользователь успешно создан и сохранен")
	return user, nil
}

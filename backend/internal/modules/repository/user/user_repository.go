package repository

import (
	"context"

	"github.com/bullockz21/pet_project21/internal/modules/domain"
	entityUser "github.com/bullockz21/pet_project21/internal/modules/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	return &UserRepository{
		db:  db,
		log: log,
	}
}

// Преобразование доменной модели в entity
func domainToEntity(u *domain.User) *entityUser.User {
	return &entityUser.User{
		TelegramID: u.TelegramID,
		Username:   u.Username,
		FirstName:  u.FirstName,
		Language:   u.Language,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

// Преобразование entity в доменную модель
func entityToDomain(e *entityUser.User) (*domain.User, error) {
	u, err := domain.NewUser(e.TelegramID, e.Username, e.FirstName, e.Language)
	if err != nil {
		return nil, err
	}
	u.CreatedAt = e.CreatedAt
	u.UpdatedAt = e.UpdatedAt
	return u, nil
}

func (r *UserRepository) Save(ctx context.Context, u *domain.User) error {
	r.log.WithField("telegram_id", u.TelegramID).Info("Сохраняем пользователя в базу данных")

	eUser := domainToEntity(u)
	err := r.db.WithContext(ctx).
		Where(&entityUser.User{TelegramID: eUser.TelegramID}).
		Assign(eUser).
		FirstOrCreate(eUser).Error

	if err != nil {
		r.log.WithError(err).Error("Ошибка при сохранении пользователя")
		return err
	}

	r.log.WithField("telegram_id", eUser.TelegramID).Info("Пользователь успешно сохранён")
	return nil
}

func (r *UserRepository) FindByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	r.log.WithField("telegram_id", telegramID).Info("Поиск пользователя по TelegramID")

	var eUser entityUser.User
	if err := r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).First(&eUser).Error; err != nil {
		r.log.WithError(err).Error("Ошибка при поиске пользователя")
		return nil, err
	}

	user, err := entityToDomain(&eUser)
	if err != nil {
		r.log.WithError(err).Error("Ошибка при преобразовании entity в доменную модель")
		return nil, err
	}

	r.log.WithField("telegram_id", telegramID).Info("Пользователь найден")
	return user, nil
}

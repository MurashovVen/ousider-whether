package service

import (
	"context"
	"fmt"
	"strconv"

	sdkentities "github.com/MurashovVen/outsider-sdk/entities"
	"github.com/MurashovVen/outsider-sdk/tg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"outsider-whether/internal/entities"
	"outsider-whether/internal/models"
	"outsider-whether/internal/repository"
)

type WhetherService struct {
	tg *tg.Client

	repo *repository.Repository
}

func New(tg *tg.Client, repo *repository.Repository) *WhetherService {
	return &WhetherService{
		tg:   tg,
		repo: repo,
	}
}

func (s *WhetherService) ActionProcess(ctx context.Context, fromChat int64, action string) error {
	actionType := sdkentities.ActionTypeParseString(action)

	switch actionType {
	case sdkentities.ActionWhetherConfigureType:
		return s.actionWhetherConfigureProcess(fromChat)

	case sdkentities.ActionWhetherTemperatureConfigureType:
		cfg, err := entities.WhetherTemperatureConfigurationParseString(action)
		if err != nil {
			return err
		}

		return s.actionWhetherConfigureTemperatureProcess(ctx, cfg.Temperature, fromChat)

	default:
		return ErrUnknownAction
	}
}

func (s *WhetherService) actionWhetherConfigureProcess(fromChat int64) error {
	_, err := s.tg.Send(
		&tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: fromChat,
				ReplyMarkup: tgbotapi.InlineKeyboardMarkup{
					InlineKeyboard: callbackDataConfigureWhetherTemperatureCreateButtons(),
				},
			},
			Text: `Выберете критическое значение температуры`,
		},
	)
	if err != nil {
		return fmt.Errorf("sending message: %w", err)
	}

	return err
}

func callbackDataConfigureWhetherTemperatureCreateButtons() [][]tgbotapi.InlineKeyboardButton {
	var (
		rowLen = 5

		from, to = -40, 40

		res = make([][]tgbotapi.InlineKeyboardButton, 0, (to-from+1)/rowLen+1)

		currRow = make([]tgbotapi.InlineKeyboardButton, 0, rowLen)
	)

	for ; from <= to; from++ {
		cbData := entities.NewWhetherTemperatureConfiguration(from).String()
		currRow = append(currRow,
			tgbotapi.InlineKeyboardButton{
				Text:         strconv.Itoa(from),
				CallbackData: &cbData,
			},
		)

		if len(currRow) == rowLen {
			res = append(res, currRow)

			currRow = make([]tgbotapi.InlineKeyboardButton, 0, rowLen)
		}
	}

	return res
}

func (s *WhetherService) actionWhetherConfigureTemperatureProcess(
	ctx context.Context, temperature int, fromChat int64) error {
	_, err := s.repo.WhetherConfigurationSave(
		ctx,
		&models.WhetherConfigurationMDBv1{
			ChatID:      fromChat,
			Temperature: temperature,
		},
	)
	if err != nil {
		return fmt.Errorf("saving config: %w", err)
	}

	_, err = s.tg.Send(
		&tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: fromChat,
			},
			Text: `Вы подписались на обновления и сконфигурировали критическую температуру. Спасибо)`,
		},
	)
	if err != nil {
		return fmt.Errorf("sending message: %w", err)
	}

	return nil
}

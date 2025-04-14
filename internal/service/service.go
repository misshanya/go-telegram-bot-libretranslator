package service

import (
	"context"
	"fmt"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/repository"
	"log"
)

type Service interface {
	RegisterUser(ctx context.Context, id int64) bool

	Translate(ctx context.Context, textToTranslate string, userID int64) (string, error)

	IsAutoDetect(ctx context.Context, uid int64) bool

	GetSourceLang(ctx context.Context, id int64) (string, error)
	GetTargetLang(ctx context.Context, id int64) (string, error)

	SwitchSourceLang(ctx context.Context, id int64) error
	SwitchTargetLang(ctx context.Context, id int64) error

	ChangeAutoDetect(ctx context.Context, id int64) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) RegisterUser(ctx context.Context, id int64) bool {
	return s.repo.RegisterUser(ctx, id)
}

func (s *service) Translate(ctx context.Context, textToTranslate string, userID int64) (string, error) {
	var sourceLang, targetLang string
	var err error

	if s.repo.IsAutoDetect(ctx, userID) {
		sourceLang, err = detectLanguage(textToTranslate)
		if err != nil {
			log.Println(err)
			return "", err
		}
		if sourceLang == "ru" {
			targetLang = "en"
		} else {
			targetLang = "ru"
		}
	} else {
		sourceLang, err = s.repo.GetSourceLang(ctx, userID)
		if err != nil {
			log.Println(fmt.Errorf("[translate] failed to get original language: %w", err))
			return "", err
		}
		targetLang, err = s.repo.GetTargetLang(ctx, userID)
		if err != nil {
			log.Println(fmt.Errorf("[translate] failed to get target language: %w", err))
			return "", err
		}
	}

	translatedText, err := translate(textToTranslate, sourceLang, targetLang)
	if err != nil {
		log.Println(fmt.Errorf("[translate] failed to translate: %w", err))
		return "", err
	}
	return translatedText, nil
}

func (s *service) SwitchSourceLang(ctx context.Context, id int64) error {
	currentSourceLang, err := s.repo.GetSourceLang(ctx, id)
	if err != nil {
		log.Println(err)
	}
	newSourceLang := getOppositeLang(currentSourceLang)
	return s.repo.SetSourceLang(ctx, id, newSourceLang)
}

func (s *service) SwitchTargetLang(ctx context.Context, id int64) error {
	currentTargetLang, err := s.repo.GetTargetLang(ctx, id)
	if err != nil {
		log.Println(err)
	}
	newTargetLang := getOppositeLang(currentTargetLang)
	return s.repo.SetTargetLang(ctx, id, newTargetLang)
}

func (s *service) ChangeAutoDetect(ctx context.Context, id int64) error {
	return s.repo.ChangeAutoDetect(ctx, id)
}

func (s *service) GetSourceLang(ctx context.Context, id int64) (string, error) {
	return s.repo.GetSourceLang(ctx, id)
}

func (s *service) GetTargetLang(ctx context.Context, id int64) (string, error) {
	return s.repo.GetTargetLang(ctx, id)
}

func (s *service) IsAutoDetect(ctx context.Context, uid int64) bool {
	return s.repo.IsAutoDetect(ctx, uid)
}

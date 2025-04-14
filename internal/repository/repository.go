package repository

import (
	"context"
	"github.com/misshanya/go-telegram-bot-libretranslator/internal/db/users"
	"log"
)

type Repository interface {
	RegisterUser(ctx context.Context, id int64) bool

	IsAutoDetect(ctx context.Context, id int64) bool
	ChangeAutoDetect(ctx context.Context, id int64) error

	GetSourceLang(ctx context.Context, id int64) (string, error)
	GetTargetLang(ctx context.Context, id int64) (string, error)

	SetSourceLang(ctx context.Context, id int64, newSourceLang string) error
	SetTargetLang(ctx context.Context, id int64, newTargetLang string) error
}

type repository struct {
	queries *users.Queries
}

func NewRepository(queries *users.Queries) Repository {
	return &repository{queries: queries}
}

// RegisterUser attempts to register a new user with the given UID.
// It returns true if the user was successfully registered, or false if the user already exists.
func (r *repository) RegisterUser(ctx context.Context, id int64) bool {
	_, err := r.queries.GetUser(ctx, id)
	if err == nil {
		return false
	}
	_, err = r.queries.CreateUser(ctx, id)
	return err != nil
}

// IsAutoDetect checks if the user's language autodetect feature is enabled.
// It returns true if enabled, false otherwise.
func (r *repository) IsAutoDetect(ctx context.Context, id int64) bool {
	user, err := r.queries.GetUser(ctx, id)
	if err != nil {
		return false
	}
	return user.LangAutodetect
}

// ChangeAutoDetect toggles the language autodetect feature for the user with the given UID.
// It logs an error if the operation fails.
func (r *repository) ChangeAutoDetect(ctx context.Context, id int64) error {
	_, err := r.queries.ChangeLangAutodetect(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetSourceLang returns the user's source language setting
func (r *repository) GetSourceLang(ctx context.Context, id int64) (string, error) {
	sourceLang, err := r.queries.GetSourceLang(ctx, id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return sourceLang, nil
}

// GetTargetLang returns the user's target language setting
func (r *repository) GetTargetLang(ctx context.Context, id int64) (string, error) {
	sourceLang, err := r.queries.GetTargetLang(ctx, id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return sourceLang, nil
}

// SetSourceLang updates the user's source language setting
func (r *repository) SetSourceLang(ctx context.Context, id int64, newSourceLang string) error {
	return r.queries.SetSourceLang(ctx, users.SetSourceLangParams{
		SourceLang: newSourceLang,
		TgID:       id,
	})
}

// SetTargetLang updates the user's target language setting
func (r *repository) SetTargetLang(ctx context.Context, id int64, newTargetLang string) error {
	return r.queries.SetTargetLang(ctx, users.SetTargetLangParams{
		TargetLang: newTargetLang,
		TgID:       id,
	})
}

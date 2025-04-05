package repositories

import (
	"context"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
)

type SessionRepository interface {
	UpsertSession(ctx context.Context, session models.Session) error
	GetSessionToken(ctx context.Context, token string) (*models.Session, error)
	DeleteSession(ctx context.Context, ID string) error
}

type sessionRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewSessionRepository(di *pkgs.Di) (SessionRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}
	return &sessionRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (s *sessionRepository) UpsertSession(ctx context.Context, session models.Session) error {
	existingSession := &models.Session{}
	err := s.repo.FindOne(ctx, existingSession, persistence.WithConditions("id = ?", session.ID))
	if err != nil && err != persistence.ErrRecordNotFound {
		return err
	}

	if err == persistence.ErrRecordNotFound {
		return s.repo.Create(ctx, &session)
	}
	return s.repo.Update(ctx, &session)
}

func (s *sessionRepository) GetSessionToken(ctx context.Context, token string) (*models.Session, error) {
	session := &models.Session{}
	err := s.repo.FindOne(ctx, session, persistence.WithConditions("token = ?", token))
	if err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}
	return session, nil
}

func (s *sessionRepository) DeleteSession(ctx context.Context, ID string) error {
	if err := s.repo.Delete(ctx, ID, &models.Session{}); err != nil {
		return err
	}

	return nil
}

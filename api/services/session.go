package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/g-villarinho/xp-life-api/repositories"
	"github.com/google/uuid"
)

type SessionService interface {
	CreateSession(ctx context.Context, user *models.User, ssi models.SessionSecurityInfo) (*models.Session, error)
	ValidSession(ctx context.Context, token string) (*models.Session, error)
}

type sessionService struct {
	di *pkgs.Di
	ts TokenService
	sr repositories.SessionRepository
}

func NewSessionService(di *pkgs.Di) (SessionService, error) {
	tokenService, err := pkgs.Invoke[TokenService](di)
	if err != nil {
		return nil, err
	}

	sessionRepository, err := pkgs.Invoke[repositories.SessionRepository](di)
	if err != nil {
		return nil, err
	}
	return &sessionService{
		di: di,
		ts: tokenService,
		sr: sessionRepository,
	}, nil
}

func (s *sessionService) CreateSession(ctx context.Context, user *models.User, ssi models.SessionSecurityInfo) (*models.Session, error) {
	now := time.Now()
	session := models.Session{
		ID:        uuid.New(),
		Email:     user.Email,
		IP:        ssi.IP,
		UserID:    user.ID,
		UserAgent: ssi.UserAgent,
		ExpiresAt: now.Add(time.Hour * 1),
		CreatedAt: now,
	}

	// Quando a sessão é validada, o userID e o sessionID é adicionado ao token
	// Isso é feito para garantir que o token só possa ser usado pelo usuário correto
	token, err := s.ts.CreateToken(ctx, "", "", session.Email, session.CreatedAt, session.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("create token: %w", err)
	}

	session.Token = token
	if err := s.sr.UpsertSession(ctx, session); err != nil {
		return nil, fmt.Errorf("upsert session: %w", err)
	}

	return &session, nil
}

func (s *sessionService) ValidSession(ctx context.Context, authToken string) (*models.Session, error) {
	now := time.Now()

	session, err := s.sr.GetSessionToken(ctx, authToken)
	if err != nil {
		return nil, fmt.Errorf("get session by token: %w", err)
	}

	if session == nil {
		return nil, models.ErrSessionNotFoundOrExpired
	}

	if session.IsExpired() {
		return nil, models.ErrSessionNotFoundOrExpired
	}

	session.VerifiedAt = sql.NullTime{Time: now, Valid: true}
	session.ExpiresAt = now.Add(time.Hour * 24 * 7)

	authToken, err = s.ts.CreateToken(ctx, session.UserID.String(), session.ID.String(), session.Email, session.CreatedAt, session.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("create token: %w", err)
	}

	session.Token = authToken
	if err := s.sr.UpsertSession(ctx, *session); err != nil {
		return nil, fmt.Errorf("upsert session: %w", err)
	}

	return session, nil
}

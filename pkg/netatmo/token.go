package netatmo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/url"

	"github.com/ViBiOh/absto/pkg/model"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

func (s *Service) refreshAccessToken(ctx context.Context) error {
	payload := url.Values{}
	payload.Add("grant_type", "refresh_token")
	payload.Add("refresh_token", s.token.RefreshToken)
	payload.Add("client_id", s.clientID)
	payload.Add("client_secret", s.clientSecret)

	resp, err := request.Post(netatmoRefreshTokenURL).Form(ctx, payload)
	if err != nil {
		return fmt.Errorf("post: %w", err)
	}

	var token Token
	if err := httpjson.Read(resp, &token); err != nil {
		return fmt.Errorf("read: %w", err)
	}

	if err := s.saveToken(ctx, token); err != nil {
		return fmt.Errorf("save: %w", err)
	}

	return nil
}

func (s *Service) loadToken(ctx context.Context) error {
	reader, err := s.storage.ReadFrom(ctx, "netatmo.json")
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	encrypted, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	payload, err := s.e2e.Decrypt(encrypted)
	if err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "decrypt token", slog.Any("error", err))
		payload = encrypted
	}

	if err := json.Unmarshal(payload, &s.token); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

func (s *Service) saveToken(ctx context.Context, token Token) error {
	s.token = token

	payload, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	encrypted, err := s.e2e.Encrypt(payload)
	if err != nil {
		return fmt.Errorf("encrypt: %w", err)
	}

	err = s.storage.WriteTo(ctx, "netatmo.json", bytes.NewReader(encrypted), model.WriteOpts{Size: int64(len(encrypted))})
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

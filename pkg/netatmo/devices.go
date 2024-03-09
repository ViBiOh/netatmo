package netatmo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ViBiOh/absto/pkg/model"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

const (
	netatmoGetStationsDataURL   = "https://api.netatmo.com/api/getstationsdata?access_token="
	netatmoGetHomeCoachsDataURL = "https://api.netatmo.com/api/gethomecoachsdata?access_token="
	netatmoRefreshTokenURL      = "https://api.netatmo.com/oauth2/token"
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

func (s *Service) getData(ctx context.Context, url string) (StationsData, error) {
	resp, err := request.Get(fmt.Sprintf("%s%s", url, s.token.AccessToken)).Send(ctx, nil)
	if err != nil && resp != nil && resp.StatusCode == http.StatusForbidden {
		if err := s.refreshAccessToken(ctx); err != nil {
			return StationsData{}, fmt.Errorf("refresh: %w", err)
		}

		resp, err = request.Get(fmt.Sprintf("%s%s", url, s.token.AccessToken)).Send(ctx, nil)
	}

	if err != nil {
		return StationsData{}, fmt.Errorf("fetch: %w", err)
	}

	var infos StationsData
	if err := httpjson.Read(resp, &infos); err != nil {
		return StationsData{}, fmt.Errorf("read: %w", err)
	}

	return infos, nil
}

func (s *Service) getDevices(ctx context.Context) ([]Device, error) {
	devices := make([]Device, 0)

	if s.HasScope("read_station") {
		stationsData, err := s.getData(ctx, netatmoGetStationsDataURL)
		if err != nil {
			return nil, fmt.Errorf("get station: %w", err)
		}

		devices = append(devices, stationsData.Body.Devices...)
	}

	if s.HasScope("read_homecoach") {
		homeCoachsData, err := s.getData(ctx, netatmoGetHomeCoachsDataURL)
		if err != nil {
			return nil, fmt.Errorf("get homecoach: %w", err)
		}

		devices = append(devices, homeCoachsData.Body.Devices...)
	}

	return devices, nil
}

func (s *Service) loadToken(ctx context.Context) error {
	reader, err := s.storage.ReadFrom(ctx, "netatmo.json")
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	payload, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("read: %w", err)
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

	err = s.storage.WriteTo(ctx, "netatmo.json", bytes.NewReader(payload), model.WriteOpts{Size: int64(len(payload))})
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

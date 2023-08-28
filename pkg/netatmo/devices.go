package netatmo

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

const (
	netatmoGetStationsDataURL   = "https://api.netatmo.com/api/getstationsdata?access_token="
	netatmoGetHomeCoachsDataURL = "https://api.netatmo.com/api/gethomecoachsdata?access_token="
	netatmoRefreshTokenURL      = "https://api.netatmo.com/oauth2/token"
)

func (s *Service) refreshAccessToken(ctx context.Context) error {
	slog.Info("Refreshing token")

	payload := url.Values{
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{s.refreshToken},
		"client_id":     []string{s.clientID},
		"client_secret": []string{s.clientSecret},
	}

	resp, err := request.Post(netatmoRefreshTokenURL).Form(ctx, payload)
	if err != nil {
		return err
	}

	var token Token
	if err := httpjson.Read(resp, &token); err != nil {
		return fmt.Errorf("read token: %w", err)
	}

	s.accessToken = token.AccessToken

	return nil
}

func (s *Service) getData(ctx context.Context, url string) (StationsData, error) {
	if !s.Enabled() {
		return noneStationsData, fmt.Errorf("app not enabled")
	}

	resp, err := request.Get(fmt.Sprintf("%s%s", url, s.accessToken)).Send(ctx, nil)
	if err != nil && resp != nil && resp.StatusCode == http.StatusForbidden {
		if err := s.refreshAccessToken(ctx); err != nil {
			return noneStationsData, err
		}

		resp, err = request.Get(fmt.Sprintf("%s%s", url, s.accessToken)).Send(ctx, nil)
	}

	if err != nil {
		return noneStationsData, err
	}

	var infos StationsData
	if err := httpjson.Read(resp, &infos); err != nil {
		return noneStationsData, fmt.Errorf("read data: %w", err)
	}

	return infos, nil
}

func (s *Service) getDevices(ctx context.Context) ([]Device, error) {
	devices := make([]Device, 0)

	if s.HasScope("read_station") {
		stationsData, err := s.getData(ctx, netatmoGetStationsDataURL)
		if err != nil {
			return nil, fmt.Errorf("read station: %w", err)
		}
		devices = append(devices, stationsData.Body.Devices...)
	}

	if s.HasScope("read_homecoach") {
		homeCoachsData, err := s.getData(ctx, netatmoGetHomeCoachsDataURL)
		if err != nil {
			return nil, fmt.Errorf("read homecoach: %w", err)
		}
		devices = append(devices, homeCoachsData.Body.Devices...)
	}

	return devices, nil
}

package netatmo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

const (
	netatmoEndpoint             = "https://api.netatmo.com/"
	netatmoGetStationsDataURL   = netatmoEndpoint + "/api/getstationsdata"
	netatmoGetHomeCoachsDataURL = netatmoEndpoint + "/api/gethomecoachsdata"
	netatmoRefreshTokenURL      = netatmoEndpoint + "/oauth2/token"
)

func (s *Service) getData(ctx context.Context, url string) (StationsData, error) {
	if s.token.IsExpired() {
		if err := s.refreshAccessToken(ctx); err != nil {
			return StationsData{}, fmt.Errorf("refresh: %w", err)
		}
	}

	resp, err := request.New().Header("Authorization", "Bearer "+s.token.AccessToken).Get(url).Send(ctx, nil)
	if err != nil && resp != nil && resp.StatusCode == http.StatusForbidden {
		if err := s.refreshAccessToken(ctx); err != nil {
			return StationsData{}, fmt.Errorf("refresh: %w", err)
		}

		resp, err = request.New().Header("Authorization", "Bearer "+s.token.AccessToken).Get(url).Send(ctx, nil)
	}

	if err != nil {
		return StationsData{}, fmt.Errorf("fetch: %w", err)
	}

	infos, err := httpjson.Read[StationsData](resp)
	if err != nil {
		return StationsData{}, fmt.Errorf("read: %w", err)
	}

	return infos, nil
}

func (s *Service) getDevices(ctx context.Context) ([]Device, error) {
	var devices []Device

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

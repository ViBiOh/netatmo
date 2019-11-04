package netatmo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ViBiOh/httputils/v3/pkg/logger"
	"github.com/ViBiOh/httputils/v3/pkg/request"
)

const (
	netatmoGetStationsDataURL   = "https://api.netatmo.com/api/getstationsdata?access_token="
	netatmoGetHomeCoachsDataURL = "https://api.netatmo.com/api/gethomecoachsdata?access_token="
	netatmoRefreshTokenURL      = "https://api.netatmo.com/oauth2/token"
)

func (a *app) refreshAccessToken(ctx context.Context) error {
	logger.Info("Refreshing token")

	payload := url.Values{
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{a.refreshToken},
		"client_id":     []string{a.clientID},
		"client_secret": []string{a.clientSecret},
	}

	resp, err := request.New().Post(netatmoRefreshTokenURL).Form(ctx, payload)
	if err != nil {
		return err
	}

	rawData, err := request.ReadBodyResponse(resp)
	if err != nil {
		return err
	}

	var token Token
	if err := json.Unmarshal(rawData, &token); err != nil {
		return err
	}

	a.accessToken = token.AccessToken

	return nil
}

func (a *app) getData(ctx context.Context, url string) (*StationsData, error) {
	if !a.Enabled() {
		return nil, fmt.Errorf("app not enabled")
	}

	resp, err := request.New().Get(fmt.Sprintf("%s%s", url, a.accessToken)).Send(ctx, nil)
	if err != nil && resp != nil && resp.StatusCode == http.StatusForbidden {
		if err := a.refreshAccessToken(ctx); err != nil {
			return nil, err
		}

		resp, err = request.New().Get(fmt.Sprintf("%s%s", url, a.accessToken)).Send(ctx, nil)
	}

	rawData, err := request.ReadBodyResponse(resp)
	if err != nil {
		return nil, err
	}

	var infos StationsData
	if err := json.Unmarshal(rawData, &infos); err != nil {
		return nil, err
	}

	return &infos, nil
}

func (a *app) GetDevices(ctx context.Context) ([]Device, error) {
	devices := make([]Device, 0)

	if a.HasScope("read_station") {
		stationsData, err := a.getData(ctx, netatmoGetStationsDataURL)
		if err != nil {
			return nil, err
		}
		devices = append(devices, stationsData.Body.Devices...)
	}

	if a.HasScope("read_homecoach") {
		homeCoachsData, err := a.getData(ctx, netatmoGetHomeCoachsDataURL)
		if err != nil {
			return nil, err
		}
		devices = append(devices, homeCoachsData.Body.Devices...)
	}

	return devices, nil
}

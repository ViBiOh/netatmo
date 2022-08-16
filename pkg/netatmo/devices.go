package netatmo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

const (
	netatmoGetStationsDataURL   = "https://api.netatmo.com/api/getstationsdata?access_token="
	netatmoGetHomeCoachsDataURL = "https://api.netatmo.com/api/gethomecoachsdata?access_token="
	netatmoRefreshTokenURL      = "https://api.netatmo.com/oauth2/token"
)

func (a *App) refreshAccessToken(ctx context.Context) error {
	logger.Info("Refreshing token")

	payload := url.Values{
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{a.refreshToken},
		"client_id":     []string{a.clientID},
		"client_secret": []string{a.clientSecret},
	}

	resp, err := request.Post(netatmoRefreshTokenURL).Form(ctx, payload)
	if err != nil {
		return err
	}

	var token Token
	if err := httpjson.Read(resp, &token); err != nil {
		return fmt.Errorf("read token: %w", err)
	}

	a.accessToken = token.AccessToken

	return nil
}

func (a *App) getData(ctx context.Context, url string) (StationsData, error) {
	if !a.Enabled() {
		return noneStationsData, fmt.Errorf("app not enabled")
	}

	resp, err := request.Get(fmt.Sprintf("%s%s", url, a.accessToken)).Send(ctx, nil)
	if err != nil && resp != nil && resp.StatusCode == http.StatusForbidden {
		if err := a.refreshAccessToken(ctx); err != nil {
			return noneStationsData, err
		}

		resp, err = request.Get(fmt.Sprintf("%s%s", url, a.accessToken)).Send(ctx, nil)
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

func (a *App) getDevices(ctx context.Context) ([]Device, error) {
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

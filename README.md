# netatmo

[![Build Status](https://travis-ci.com/ViBiOh/goweb.svg?branch=master)](https://travis-ci.com/ViBiOh/goweb)
[![codecov](https://codecov.io/gh/ViBiOh/goweb/branch/master/graph/badge.svg)](https://codecov.io/gh/ViBiOh/goweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/ViBiOh/goweb)](https://goreportcard.com/report/github.com/ViBiOh/goweb)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_netatmo&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_netatmo)

## CI

Following variables are required for CI:

| Name | Purpose |
|:--:|:--:|
| **DOCKER_USER** | for publishing Docker image |
| **DOCKER_PASS** | for publishing Docker image |
| **CODECOV_TOKEN** | for publishing coverage to codecov |

## Usage

```bash
Usage of netatmo:
  -accessToken string
        [netatmo] Access Token {NETATMO_ACCESS_TOKEN}
  -apiAddress string
        [api] Listen address {NETATMO_API_ADDRESS}
  -apiCert string
        [api] Certificate file {NETATMO_API_CERT}
  -apiGraceDuration string
        [api] Grace duration when SIGTERM received {NETATMO_API_GRACE_DURATION} (default "15s")
  -apiKey string
        [api] Key file {NETATMO_API_KEY}
  -apiOkStatus int
        [api] Healthy HTTP Status code {NETATMO_API_OK_STATUS} (default 204)
  -apiPort uint
        [api] Listen port {NETATMO_API_PORT} (default 1080)
  -clientID string
        [netatmo] Client ID {NETATMO_CLIENT_ID}
  -clientSecret string
        [netatmo] Client Secret {NETATMO_CLIENT_SECRET}
  -corsCredentials
        [cors] Access-Control-Allow-Credentials {NETATMO_CORS_CREDENTIALS}
  -corsExpose string
        [cors] Access-Control-Expose-Headers {NETATMO_CORS_EXPOSE}
  -corsHeaders string
        [cors] Access-Control-Allow-Headers {NETATMO_CORS_HEADERS} (default "Content-Type")
  -corsMethods string
        [cors] Access-Control-Allow-Methods {NETATMO_CORS_METHODS} (default "GET")
  -corsOrigin string
        [cors] Access-Control-Allow-Origin {NETATMO_CORS_ORIGIN} (default "*")
  -csp string
        [owasp] Content-Security-Policy {NETATMO_CSP} (default "default-src 'self'; base-uri 'self'")
  -frameOptions string
        [owasp] X-Frame-Options {NETATMO_FRAME_OPTIONS} (default "deny")
  -hsts
        [owasp] Indicate Strict Transport Security {NETATMO_HSTS} (default true)
  -prometheusPath string
        [prometheus] Path for exposing metrics {NETATMO_PROMETHEUS_PATH} (default "/metrics")
  -refreshToken string
        [netatmo] Refresh Token {NETATMO_REFRESH_TOKEN}
  -scopes string
        [netatmo] Scopes, comma separated {NETATMO_SCOPES}
  -swaggerTitle string
        [swagger] API Title {NETATMO_SWAGGER_TITLE} (default "API")
  -swaggerVersion string
        [swagger] API Version {NETATMO_SWAGGER_VERSION} (default "1.0.0")
  -url string
        [alcotest] URL to check {NETATMO_URL}
  -userAgent string
        [alcotest] User-Agent for check {NETATMO_USER_AGENT} (default "Alcotest")
```

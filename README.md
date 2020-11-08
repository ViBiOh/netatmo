# netatmo

[![Build](https://github.com/ViBiOh/netatmo/workflows/Build/badge.svg)](https://github.com/ViBiOh/netatmo/actions)
[![codecov](https://codecov.io/gh/ViBiOh/goweb/branch/master/graph/badge.svg)](https://codecov.io/gh/ViBiOh/goweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/ViBiOh/goweb)](https://goreportcard.com/report/github.com/ViBiOh/goweb)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_netatmo&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_netatmo)

## CI

Following variables are required for CI:

|       Name        |              Purpose               |
| :---------------: | :--------------------------------: |
|  **DOCKER_USER**  |    for publishing Docker image     |
|  **DOCKER_PASS**  |    for publishing Docker image     |
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
        [api] Grace duration when SIGTERM received {NETATMO_API_GRACE_DURATION} (default "30s")
  -apiIdleTimeout string
        [api] Idle Timeout {NETATMO_API_IDLE_TIMEOUT} (default "2m")
  -apiKey string
        [api] Key file {NETATMO_API_KEY}
  -apiOkStatus int
        [api] Healthy HTTP Status code {NETATMO_API_OK_STATUS} (default 204)
  -apiPort uint
        [api] Listen port {NETATMO_API_PORT} (default 1080)
  -apiReadTimeout string
        [api] Read Timeout {NETATMO_API_READ_TIMEOUT} (default "5s")
  -apiShutdownTimeout string
        [api] Shutdown Timeout {NETATMO_API_SHUTDOWN_TIMEOUT} (default "10s")
  -apiWriteTimeout string
        [api] Write Timeout {NETATMO_API_WRITE_TIMEOUT} (default "10s")
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
  -loggerJson
        [logger] Log format as JSON {NETATMO_LOGGER_JSON}
  -loggerLevel string
        [logger] Logger level {NETATMO_LOGGER_LEVEL} (default "INFO")
  -loggerLevelKey string
        [logger] Key for level in JSON {NETATMO_LOGGER_LEVEL_KEY} (default "level")
  -loggerMessageKey string
        [logger] Key for message in JSON {NETATMO_LOGGER_MESSAGE_KEY} (default "message")
  -loggerTimeKey string
        [logger] Key for timestamp in JSON {NETATMO_LOGGER_TIME_KEY} (default "time")
  -prometheusIgnore string
        [prometheus] Ignored path prefixes for metrics, comma separated {NETATMO_PROMETHEUS_IGNORE}
  -prometheusPath string
        [prometheus] Path for exposing metrics {NETATMO_PROMETHEUS_PATH} (default "/metrics")
  -refreshToken string
        [netatmo] Refresh Token {NETATMO_REFRESH_TOKEN}
  -scopes string
        [netatmo] Scopes, comma separated {NETATMO_SCOPES}
  -url string
        [alcotest] URL to check {NETATMO_URL}
  -userAgent string
        [alcotest] User-Agent for check {NETATMO_USER_AGENT} (default "Alcotest")
```

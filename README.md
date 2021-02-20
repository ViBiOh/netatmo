# netatmo

[![Build](https://github.com/ViBiOh/netatmo/workflows/Build/badge.svg)](https://github.com/ViBiOh/netatmo/actions)
[![codecov](https://codecov.io/gh/ViBiOh/netatmo/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/netatmo)
[![Go Report Card](https://goreportcard.com/badge/github.com/ViBiOh/netatmo)](https://goreportcard.com/report/github.com/ViBiOh/netatmo)
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
  -address string
        [server] Listen address {NETATMO_ADDRESS}
  -cert string
        [server] Certificate file {NETATMO_CERT}
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
  -graceDuration string
        [http] Grace duration when SIGTERM received {NETATMO_GRACE_DURATION} (default "30s")
  -hsts
        [owasp] Indicate Strict Transport Security {NETATMO_HSTS} (default true)
  -idleTimeout string
        [server] Idle Timeout {NETATMO_IDLE_TIMEOUT} (default "2m")
  -key string
        [server] Key file {NETATMO_KEY}
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
  -okStatus int
        [http] Healthy HTTP Status code {NETATMO_OK_STATUS} (default 204)
  -port uint
        [server] Listen port {NETATMO_PORT} (default 1080)
  -prometheusAddress string
        [prometheus] Listen address {NETATMO_PROMETHEUS_ADDRESS}
  -prometheusCert string
        [prometheus] Certificate file {NETATMO_PROMETHEUS_CERT}
  -prometheusIdleTimeout string
        [prometheus] Idle Timeout {NETATMO_PROMETHEUS_IDLE_TIMEOUT} (default "10s")
  -prometheusIgnore string
        [prometheus] Ignored path prefixes for metrics, comma separated {NETATMO_PROMETHEUS_IGNORE}
  -prometheusKey string
        [prometheus] Key file {NETATMO_PROMETHEUS_KEY}
  -prometheusPort uint
        [prometheus] Listen port {NETATMO_PROMETHEUS_PORT} (default 9090)
  -prometheusReadTimeout string
        [prometheus] Read Timeout {NETATMO_PROMETHEUS_READ_TIMEOUT} (default "5s")
  -prometheusShutdownTimeout string
        [prometheus] Shutdown Timeout {NETATMO_PROMETHEUS_SHUTDOWN_TIMEOUT} (default "5s")
  -prometheusWriteTimeout string
        [prometheus] Write Timeout {NETATMO_PROMETHEUS_WRITE_TIMEOUT} (default "10s")
  -readTimeout string
        [server] Read Timeout {NETATMO_READ_TIMEOUT} (default "5s")
  -refreshToken string
        [netatmo] Refresh Token {NETATMO_REFRESH_TOKEN}
  -scopes string
        [netatmo] Scopes, comma separated {NETATMO_SCOPES}
  -shutdownTimeout string
        [server] Shutdown Timeout {NETATMO_SHUTDOWN_TIMEOUT} (default "10s")
  -url string
        [alcotest] URL to check {NETATMO_URL}
  -userAgent string
        [alcotest] User-Agent for check {NETATMO_USER_AGENT} (default "Alcotest")
  -writeTimeout string
        [server] Write Timeout {NETATMO_WRITE_TIMEOUT} (default "10s")
```

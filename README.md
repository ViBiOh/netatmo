# netatmo

[![Build Status](https://travis-ci.org/ViBiOh/goweb.svg?branch=master)](https://travis-ci.org/ViBiOh/goweb)
[![codecov](https://codecov.io/gh/ViBiOh/goweb/branch/master/graph/badge.svg)](https://codecov.io/gh/ViBiOh/goweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/ViBiOh/goweb)](https://goreportcard.com/report/github.com/ViBiOh/goweb)
[![Dependabot Status](https://api.dependabot.com/badges/status?host=github&repo=ViBiOh/goweb)](https://dependabot.com)

## CI

Following variables are required for CI:

| Name | Purpose |
|:--:|:--:|
| **GITHUB_OAUTH_TOKEN** | for creating release from  Github API |
| **DOMAIN** | for setting Traefik domain for app |
| **DEPLOY_CREDENTIALS** | for deploying app to server |
| **DOCKER_USER** | for publishing Docker image |
| **DOCKER_PASS** | for publishing Docker image |

## Usage

```bash
Usage of netatmo:
  -accessToken string
        [netatmo] Access Token {NETATMO_ACCESS_TOKEN}
  -apiAddress string
        [api] Listen address {NETATMO_API_ADDRESS}
  -apiCert string
        [api] Certificate file {NETATMO_API_CERT}
  -apiKey string
        [api] Key file {NETATMO_API_KEY}
  -apiPort int
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
  -url string
        [alcotest] URL to check {NETATMO_URL}
  -userAgent string
        [alcotest] User-Agent for check {NETATMO_USER_AGENT} (default "Golang alcotest")
```

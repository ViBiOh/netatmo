# netatmo

[![Build](https://github.com/ViBiOh/netatmo/workflows/Build/badge.svg)](https://github.com/ViBiOh/netatmo/actions)
[![codecov](https://codecov.io/gh/ViBiOh/netatmo/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/netatmo)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_netatmo&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_netatmo)

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of netatmo:
  --accessToken       string    [netatmo] Access Token ${NETATMO_ACCESS_TOKEN}
  --clientID          string    [netatmo] Client ID ${NETATMO_CLIENT_ID}
  --clientSecret      string    [netatmo] Client Secret ${NETATMO_CLIENT_SECRET}
  --graceDuration     duration  [http] Grace duration when SIGTERM received ${NETATMO_GRACE_DURATION} (default 30s)
  --loggerJson                  [logger] Log format as JSON ${NETATMO_LOGGER_JSON} (default false)
  --loggerLevel       string    [logger] Logger level ${NETATMO_LOGGER_LEVEL} (default "INFO")
  --loggerLevelKey    string    [logger] Key for level in JSON ${NETATMO_LOGGER_LEVEL_KEY} (default "level")
  --loggerMessageKey  string    [logger] Key for message in JSON ${NETATMO_LOGGER_MESSAGE_KEY} (default "msg")
  --loggerTimeKey     string    [logger] Key for timestamp in JSON ${NETATMO_LOGGER_TIME_KEY} (default "time")
  --okStatus          int       [http] Healthy HTTP Status code ${NETATMO_OK_STATUS} (default 204)
  --refreshToken      string    [netatmo] Refresh Token ${NETATMO_REFRESH_TOKEN}
  --scopes            string    [netatmo] Scopes, comma separated ${NETATMO_SCOPES}
  --telemetryRate     string    [telemetry] OpenTelemetry sample rate, 'always', 'never' or a float value ${NETATMO_TELEMETRY_RATE} (default "always")
  --telemetryURL      string    [telemetry] OpenTelemetry gRPC endpoint (e.g. otel-exporter:4317) ${NETATMO_TELEMETRY_URL}
```

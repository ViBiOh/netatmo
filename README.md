# netatmo

[![Build](https://github.com/ViBiOh/netatmo/workflows/Build/badge.svg)](https://github.com/ViBiOh/netatmo/actions)
[![codecov](https://codecov.io/gh/ViBiOh/netatmo/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/netatmo)

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of netatmo:
  --cipherSecret                string    [secret] Secret for ciphering token ${NETATMO_CIPHER_SECRET}
  --clientID                    string    [netatmo] Client ID ${NETATMO_CLIENT_ID}
  --clientSecret                string    [netatmo] Client Secret ${NETATMO_CLIENT_SECRET}
  --graceDuration               duration  [http] Grace duration when signal received ${NETATMO_GRACE_DURATION} (default 30s)
  --loggerJson                            [logger] Log format as JSON ${NETATMO_LOGGER_JSON} (default false)
  --loggerLevel                 string    [logger] Logger level ${NETATMO_LOGGER_LEVEL} (default "INFO")
  --loggerLevelKey              string    [logger] Key for level in JSON ${NETATMO_LOGGER_LEVEL_KEY} (default "level")
  --loggerMessageKey            string    [logger] Key for message in JSON ${NETATMO_LOGGER_MESSAGE_KEY} (default "msg")
  --loggerTimeKey               string    [logger] Key for timestamp in JSON ${NETATMO_LOGGER_TIME_KEY} (default "time")
  --okStatus                    int       [http] Healthy HTTP Status code ${NETATMO_OK_STATUS} (default 204)
  --scopes                      string    [netatmo] Scopes, comma separated ${NETATMO_SCOPES}
  --storageFileSystemDirectory  /data     [storage] Path to directory. Default is dynamic. /data on a server and Current Working Directory in a terminal. ${NETATMO_STORAGE_FILE_SYSTEM_DIRECTORY} (default /Users/macbook/code/netatmo)
  --storageObjectAccessKey      string    [storage] Storage Object Access Key ${NETATMO_STORAGE_OBJECT_ACCESS_KEY}
  --storageObjectBucket         string    [storage] Storage Object Bucket ${NETATMO_STORAGE_OBJECT_BUCKET}
  --storageObjectClass          string    [storage] Storage Object Class ${NETATMO_STORAGE_OBJECT_CLASS}
  --storageObjectEndpoint       string    [storage] Storage Object endpoint ${NETATMO_STORAGE_OBJECT_ENDPOINT}
  --storageObjectRegion         string    [storage] Storage Object Region ${NETATMO_STORAGE_OBJECT_REGION}
  --storageObjectSSL                      [storage] Use SSL ${NETATMO_STORAGE_OBJECT_SSL} (default true)
  --storageObjectSecretAccess   string    [storage] Storage Object Secret Access ${NETATMO_STORAGE_OBJECT_SECRET_ACCESS}
  --storagePartSize             uint      [storage] PartSize configuration ${NETATMO_STORAGE_PART_SIZE} (default 5242880)
  --telemetryRate               string    [telemetry] OpenTelemetry sample rate, 'always', 'never' or a float value ${NETATMO_TELEMETRY_RATE} (default "always")
  --telemetryURL                string    [telemetry] OpenTelemetry gRPC endpoint (e.g. otel-exporter:4317) ${NETATMO_TELEMETRY_URL}
  --telemetryUint64                       [telemetry] Change OpenTelemetry Trace ID format to an unsigned int 64 ${NETATMO_TELEMETRY_UINT64} (default true)
```

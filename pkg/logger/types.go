package logger

import (
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

// LogLabel creates a GCP Cloud Logging label field.
//
// Labels are special — in GCP they land in a separate "labels" field,
// not in the JSON payload. This makes them indexed and filterable
// in Cloud Logging console.
//
// Use LogLabel for metadata like app name, version, environment.
// Use regular key-value pairs for dynamic data like request_id, user_id.
//
// Example:
//
//	appLogger := logger.With(logger.LogLabel("app", "image-builder"), logger.LogLabel("version", "1.0.0"))
func LogLabel(key, value string) zap.Field {
	return zapdriver.Label(key, value)
}

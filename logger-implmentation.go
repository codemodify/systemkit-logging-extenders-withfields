package extenders

import (
	"time"

	logging "github.com/codemodify/systemkit-logging"
)

type withFieldsLogger struct {
	logger logging.Logger
}

// NewWithFieldsLogger -
func NewWithFieldsLogger(logger logging.Logger) WithFieldsLogger {
	return &withFieldsLogger{
		logger: logger,
	}
}

func (thisRef withFieldsLogger) TraceWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypeTrace,
		Message: fields.String(),
	})
}

func (thisRef withFieldsLogger) PanicWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypePanic,
		Message: fields.String(),
	})
}

func (thisRef withFieldsLogger) FatalWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypeFatal,
		Message: fields.String(),
	})
}

func (thisRef withFieldsLogger) ErrorWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypeError,
		Message: fields.String(),
	})
}

func (thisRef withFieldsLogger) WarningWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypeWarning,
		Message: fields.String(),
	})
}

func (thisRef withFieldsLogger) InfoWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypeInfo,
		Message: fields.String(),
	})
}

func (thisRef withFieldsLogger) SuccessWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypeSuccess,
		Message: fields.String(),
	})
}

func (thisRef withFieldsLogger) DebugWithFields(fields Fields) {
	thisRef.logger.Log(logging.LogEntry{
		Time:    time.Now(),
		Type:    logging.TypeDebug,
		Message: fields.String(),
	})
}

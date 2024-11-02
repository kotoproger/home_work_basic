package log

type Level string

const (
	LogLevelTrace   Level = "trace"
	LogLevelDebug   Level = "debug"
	LogLevelInfo    Level = "info"
	LogLevelWarning Level = "warn"
	LogLevelError   Level = "error"
	LogLevelFatal   Level = "fatal"
)

type Record struct {
	Level   Level  `json:"level,omitempty"`
	Message string `json:"message,omitempty"`
}

package log

type Level int

// log level types
const (
	DEBUG Level = iota
	INFO
	WARN
	TRACE
	ERROR
	FATAL
)

// StringToLevel convert string to Level
func StringToLevel(level string) Level {
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "TRACE":
		return TRACE
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return DEBUG
	}
}

// LevelToString convert level to string
func LevelToString(level Level) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case TRACE:
		return "TRACE"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

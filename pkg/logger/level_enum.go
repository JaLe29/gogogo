package logger

/*
Level is an enum of common levels as we know them.
  - debug 0
  - info 1
  - warn 2
  - error 3

Selected level and all "above it" (ascending order) is used to log messages.
*/
type Level int

const (
	LevelDebug Level = 0
	LevelInfo  Level = 1
	LevelWarn  Level = 2
	LevelError Level = 3
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARNING"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

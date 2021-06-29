package log

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"time"
)

const (
	timeFormat   = "2006-01-02 15:04:05"
	headerFormat = "[%l]%F:%L|%T"
)

// record
func record(level Level, format string, args ...interface{}) string {
	return formatHeader(level) + "|" + fmt.Sprintf(format, args...)
}

// headerFormatFunc
// ref https://github.com/ian-kent/go-log/blob/master/layout/pattern.go
func formatHeader(level Level) string {
	_, file, line, _ := runtime.Caller(3)
	re := regexp.MustCompile("%(\\w|%)(?:{([^}]+)})?")

	message := re.ReplaceAllStringFunc(headerFormat, func(m string) string {
		parts := re.FindStringSubmatch(m)
		switch parts[1] {
		case "l":
			return LevelToString(level)
		case "F":
			return filepath.Base(file)
		case "L":
			return strconv.Itoa(line)
		case "T":
			return time.Now().Format(timeFormat)
		case "%":
			return "%"
		}
		return m
	})

	return message
}

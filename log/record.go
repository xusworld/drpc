package log

import (
	"github.com/spf13/cast"
)

const (
	RecordDelimiter = "|"
)

// Record defines logger format
type Record struct {
	Timestamp  string
	LogLevel   int
	FileName   string
	LineNumber int
	Caller     string
}

// RecordToString
func RecordToString(record *Record) string {
	message := cast.ToString(record.Timestamp) + RecordDelimiter
	message += cast.ToString(loggerLevelToString(record.LogLevel)) + RecordDelimiter
	message += cast.ToString(record.FileName) + RecordDelimiter
	message += cast.ToString(record.LineNumber) + RecordDelimiter
	message += cast.ToString(record.Caller) + RecordDelimiter
	return message
}

// RecordToBytes
func RecordToBytes(record *Record) []byte {
	return []byte(RecordToString(record))
}

package log

import "fmt"

func Debug(args ...interface{}) {
	fmt.Println(record(DEBUG, "", args...))
}

func Debugf(format string, args ...interface{}) {
	fmt.Println(record(DEBUG, format, args...))
}

package log

import "testing"

func Test_formatHeader(t *testing.T) {

}

func Test_record(t *testing.T) {
	t.Log(record(DEBUG, "%s", "Hello, world!"))
}
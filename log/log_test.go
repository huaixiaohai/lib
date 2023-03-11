package log

import "testing"

func TestInitLog(t *testing.T) {
	Init(&Cfg{
		Level:    4,
		Output:   "file",
		FilePath: "./logs",
	})

	//for true {
	Info("info ahaha")
	Error("error ahaha")
	//}
}
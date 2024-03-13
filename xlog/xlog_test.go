package xlog

import "testing"

func TestLog(t *testing.T) {
	GoLogger.InfoP("hello")
	GoLogger.WarnP("hello")
	GoLogger.ErrorP("hello")
	//GoLogger.FatalP("hello")
	GoLogger.WarnF("hello")
	GoLogger.ErrorF("hello")
	GoLogger.FatalF("hello")
	GoLogger.InfoF("hello")

	Info("hello")
	Warn("hello")
	Error("hello")
	InfoF("hello %v", "flairamos")
	WarnF("hello %v", "flairamos")
	ErrorF("hello %v", "flairamos")
}

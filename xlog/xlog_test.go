package xlog

import "testing"

func TestLog(t *testing.T) {
	GoLogger.InfoP("asdsdaaaaaaaaaaaaaaaaaaaaaaa")
	GoLogger.WarnP("asdsdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	GoLogger.ErrorP("asdsdaaaaaaaaaaaaaaa")
	//GoLogger.FatalP("asdsdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	GoLogger.WarnF("asdsdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	GoLogger.ErrorF("asdsdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	GoLogger.FatalF("asdsdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	GoLogger.InfoF("asdsdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}

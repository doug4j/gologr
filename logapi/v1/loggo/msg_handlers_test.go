package loggo

import (
	"fmt"
	"testing"

	logapi "github.com/doug4j/gologr/logapi/v1"
)

// IMPORTANT: All of the tests with LineMessageHandler are line number dependent on this source code!
// Thus, these tests are at the top of this file.

func TestNewLineMessageHandlerDebug(t *testing.T) {
	logr := NewDefaultLogging(logapi.InfoLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewLineNumMessageHandler(LineNumMessageHandlerParms{})

	msg := "My Line Handler with Debug"
	level := logapi.DebugLogging
	debugMsg := LineMsgCallr(handlr, msg, level)
	// fmt.Printf("\n\n[%v]\n\n\n", debugMsg)
	// IMPORTANT: this test depends on the line number above for it's result
	expected := `DEBUG [github.com/doug4j/gologr/logapi/v1/loggo/msg_handlers_test.go:22]
My Line Handler with Debug`
	if debugMsg != expected {
		t.Fail()
	}
	logr.Info(debugMsg)
}

func TestNewLineMessageHandlerWarn(t *testing.T) {
	logr := NewDefaultLogging(logapi.InfoLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewLineNumMessageHandler(LineNumMessageHandlerParms{})

	msg := "My Line Handler with Warn"
	level := logapi.WarnLogging
	warnMsg := LineMsgCallr(handlr, msg, level)
	// fmt.Printf("\n%v\n\n\n", warnMsg)
	// IMPORTANT: this test depends on the line number above for it's result
	expected := `WARN [github.com/doug4j/gologr/logapi/v1/loggo/msg_handlers_test.go:42]
My Line Handler with Warn`
	if warnMsg != expected {
		t.Fail()
	}
	logr.Info(warnMsg)
}

func TestNewLineMessageHandlerError(t *testing.T) {
	logr := NewDefaultLogging(logapi.InfoLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewLineNumMessageHandler(LineNumMessageHandlerParms{})

	msg := "My Line Handler with Error"
	level := logapi.ErrorLogging
	errorMsg := LineMsgCallr(handlr, msg, level)
	// fmt.Printf("\n\n\n%v\n\n\n", errorMsg)
	// IMPORTANT: this test depends on the line number above for it's result
	expected := `ERROR [github.com/doug4j/gologr/logapi/v1/loggo/msg_handlers_test.go:62]
My Line Handler with Error`
	if errorMsg != expected {
		t.Fail()
	}
	logr.Info(errorMsg)
}

func LineMsgCallr(handlr MessageHandler, msg string, level logapi.Level) string {
	return handlr(msg, level)
}

func TestNewLineMessageHandlerInfo(t *testing.T) {
	logr := NewDefaultLogging(logapi.InfoLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewLineNumMessageHandler(LineNumMessageHandlerParms{})

	msg := "My Line Handler with Info"
	level := logapi.InfoLogging
	infoMsg := LineMsgCallr(handlr, msg, level)
	fmt.Printf("\n[%v]\n\n\n", infoMsg)
	// IMPORTANT: this test depends on the line number above for it's result
	expected := `INFO [github.com/doug4j/gologr/logapi/v1/loggo/msg_handlers_test.go:86]
My Line Handler with Info`
	if infoMsg != expected {
		t.Fail()
	}
	logr.Info(infoMsg)
}

// Note: End of LineMessageHandler and line number dependent testing from this line below.
func TestEmojiMessageHandlerDebug(t *testing.T) {
	logr := NewDefaultLogging(logapi.DebugLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewEmojiMessageHandler()

	msg := "My Debug"
	level := logapi.DebugLogging
	debugMsg := handlr(msg, level)
	if debugMsg != formatEmoji(defaultDebugEmoji, msg, level) {
		t.Fail()
	}
	logr.Info(debugMsg)
}

func TestEmojiMessageHandlerInfo(t *testing.T) {
	logr := NewDefaultLogging(logapi.DebugLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewEmojiMessageHandler()

	msg := "My Info"
	level := logapi.InfoLogging
	infoMsg := handlr(msg, level)
	if infoMsg != formatEmoji(defaultInfoEmoji, msg, level) {
		t.Fail()
	}
	logr.Info(infoMsg)
}

func TestEmojiMessageHandlerWarn(t *testing.T) {
	logr := NewDefaultLogging(logapi.DebugLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewEmojiMessageHandler()

	msg := "My Warn"
	level := logapi.WarnLogging
	warnMsg := handlr(msg, level)
	if warnMsg != formatEmoji(defaultWarnEmoji, msg, level) {
		t.Fail()
	}
	logr.Info(warnMsg)
}

func TestEmojiMessageHandlerError(t *testing.T) {
	logr := NewDefaultLogging(logapi.DebugLogging)
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)
	handlr := NewEmojiMessageHandler()

	msg := "My Error"
	level := logapi.ErrorLogging
	errorMsg := handlr(msg, level)
	if errorMsg != formatEmoji(defaultErrorEmoji, msg, level) {
		t.Fail()
	}
	logr.Info(errorMsg)
}

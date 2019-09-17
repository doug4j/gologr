package loggo

import (
	"fmt"
	"runtime"
	"strings"

	logapi "github.com/doug4j/OmniFocusTools/go/oft/pkg/logapi/v1"
)

// Adding options for not using Emoji (UTF-8/16) on windows due to
// https://stackoverflow.com/questions/44054983/how-to-output-emoji-to-console-in-node-js-on-windows

const windowsOSName = "windows"
const loggestLogPrefixPad = 6 // this is the longest log prefix plus a space

func logPrefixPad(msg string) string {
	// From https://github.com/git-time-metric/gtm/blob/master/util/string.go#L53-L88
	var overallLen = loggestLogPrefixPad
	var padStr = " "
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = msg + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// https://github.com/mrowa44/emojify/blob/master/emojify
// https://www.webfx.com/tools/emoji-cheat-sheet/

const defaultDebugEmoji string = "🔵"
const defaultInfoEmoji string = "ℹ️ "
const defaultWarnEmoji string = "🤔"
const defaultErrorEmoji string = "⛔"

// NewEmojiMessageHandler creates MessageHandler that prefixes output with an emoji on MacOS and linux
// (and text only on windows) along with the log level then the supplied message. For instance, the following are
// expected results:
//	Msg		Logging			Output (MacOS and Linux)
//	Hi		Debug			🔵 DEBUG Hi
//	Hi		Info			ℹ️ INFO  Hi
//	Hi		Warn			🤔 WARN  Hi
//	Hi		Error			⛔ ERROR Hi
func NewEmojiMessageHandler() MessageHandler {

	var msgHandler = func(msg string, logLevel logapi.Level) string {
		if runtime.GOOS == windowsOSName {
			return fmt.Sprintf(logPrefixPad(logLevel.String()) + msg)
		}
		switch logLevel {
		case logapi.InfoLogging:
			return formatEmoji(defaultInfoEmoji, msg, logLevel)
		case logapi.ErrorLogging:
			return formatEmoji(defaultErrorEmoji, msg, logLevel)
		case logapi.DebugLogging:
			return formatEmoji(defaultDebugEmoji, msg, logLevel)
		default:
			return formatEmoji(defaultWarnEmoji, msg, logLevel)
		}
	}
	return msgHandler
}

func formatEmoji(emoji, msg string, logLevel logapi.Level) string {
	return emoji + " " + logPrefixPad(logLevel.String()) + msg
}

type LineNumMessageHandlerParms struct {
	BaseDomain  string
	PadLogLevel bool
}

const assumedRepoName = "github.com"

func NewLineNumMessageHandler(parms LineNumMessageHandlerParms) MessageHandler {

	if parms.BaseDomain == "" {
		parms.BaseDomain = assumedRepoName
	}

	var msgHandler = func(msg string, logLevel logapi.Level) string {
		_, file, line, _ := runtime.Caller(2)
		fileNameFromAssumedRepo := parms.BaseDomain + strings.Split(file, parms.BaseDomain)[1]
		if parms.PadLogLevel {
			return fmt.Sprintf("%s [%s:%d]\n%v", logPrefixPad(logLevel.String()), fileNameFromAssumedRepo, line, msg)

		}
		return fmt.Sprintf("%s [%s:%d]\n%v", logLevel.String(), fileNameFromAssumedRepo, line, msg)
	}
	return msgHandler
}

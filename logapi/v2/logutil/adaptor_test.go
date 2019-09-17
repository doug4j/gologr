package logutil

import (
	"log"
	"testing"

	"github.com/doug4j/gologr/logapi/v2"
	"github.com/doug4j/gologr/logapi/v2/logfmtgo"
	"github.com/doug4j/gologr/logapi/v2/logfmtutil"
)

func TestAdaptorBasicNoParms(t *testing.T) {
	// common.LogLevel = common.InfoLogging
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)

	logr := NewLogCreating(logapi.DebugLogging, logPrintWriter{})
	// logr := NewLogCreating(logapi.DebugLogging, logOutWriter{callDepth: 3})
	logr.Debug(fmtDefaultMsg())
	logr.Info(fmtDefaultMsg())
	logr.Warn(fmtDefaultMsg())
	logr.Error(fmtDefaultMsg())
	logr.Info(fmtNonDefaultMsg())
}

func TestAdaptorBasicWithParms(t *testing.T) {
	// common.LogLevel = common.InfoLogging
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)

	// logCtx := AppLogContext {
	// 	//I18nTag       string
	// 	LogLevel      : logapi.DebugLogging,
	// 	StringCreator LogAwareStringCreating
	// 	Logr          : NewLogCreating(logapi.DebugLogging, logPrintWriter{}),
	// }

	logr := NewLogCreating(logapi.DebugLogging, logPrintWriter{})
	// logr := NewLogCreating(logapi.DebugLogging, logOutWriter{callDepth: 3})
	logr.Debug(fmtDefaultWithParmsMsg(1, 2, 3, "Go!"))
	logr.Info(fmtDefaultWithParmsMsg(4, 5, 6, "Stop!"))
	logr.Warn(fmtDefaultWithParmsMsg(7, 8, 9, "Hmmm"))
	logr.Error(fmtDefaultWithParmsMsg(10, 11, 12, "Start Again"))
	logr.Info(fmtDefaultWithParmsMsg(13, 14, 15, "Done"))
}

func TestAdaptorFancyNoParms(t *testing.T) {
	// common.LogLevel = common.InfoLogging
	// testName := common.GetCallingName()
	// common.StartTest(testName)
	// defer common.EndTest(testName)

	logr := NewLogCreating(logapi.DebugLogging, logPrintWriter{})
	// logr := NewLogCreating(logapi.DebugLogging, logOutWriter{callDepth: 3})
	logr.Debug(fancyDefaultMsg())
	logr.Info(fancyDefaultMsg())
	logr.Warn(fancyDefaultMsg())
	logr.Error(fancyDefaultMsg())
	logr.Info(fancyDefaultMsg())
}

func fancyDefaultMsg() logapi.LogAwareStringing {
	return logfmtgo.NewSrcFmtWithSkip(logfmtutil.LogFormat{
		Value: logfmtutil.NewSimpleToStringer("Hello from Basic Info"),
		//Parms:        ,
		PrefixStereo: logfmtutil.DefaultStereo,
	}, 2)
}

func fmtDefaultMsg() logapi.LogAwareStringing {
	return logfmtgo.NewSrcFmtWithSkip(logfmtutil.LogFormat{
		Value: logfmtutil.NewSimpleToStringer("Hello from Basic Info"),
		//Parms:        ,
		PrefixStereo: logfmtutil.DefaultStereo,
	}, 2)
}

func fmtNonDefaultMsg() logapi.LogAwareStringing {
	return logfmtgo.NewSrcFmtWithSkip(logfmtutil.LogFormat{
		Value: logfmtutil.NewSimpleToStringer("Hello from Basic Info"),
		//Parms:        ,
		PrefixStereo: logfmtutil.OKStereo,
	}, 2)
}

func fmtDefaultWithParmsMsg(item1, item2, item3 int, actionWord string) logapi.LogAwareStringing {
	parms := make([]interface{}, 4)
	parms[0] = item1
	parms[1] = item2
	parms[2] = item3
	parms[3] = actionWord

	prefixStereo := logfmtutil.DefaultStereo
	return logfmtgo.NewSrcFmtWithSkip(logfmtutil.LogFormat{
		Value: logfmtutil.NewPosReplaceStringer(logfmtutil.LogFormat{
			Value: logfmtutil.NewSimpleToStringer("Hello from $1, $2, $3... $4 from $1"),
			Parms: parms,
		}),
		PrefixStereo: prefixStereo,
		Parms:        parms,
	}, 2)
}

// TimeStereo
// 	//OKStereo is the logging option above info
// 	OKStereo
// 	//WorkingStereo is the logging option above warn
// 	WorkingStereo
// 	//WaitingForUserStereo is the logging option above error
// 	WaitingForUserStereo
// 	//ExitStereo is the logging option above error
// 	ExitStereo
// 	//NotImplementedStereo is the logging option above error
// 	NotImplementedStereo

// func basicInfoMsg() NewSrcFmtWithSkip(logapi.LogAwareStringing {
// 	return logfmtgoNewSrcFmt( .NewSprintfStringer(logfmtutil.LogFormat{
// 		Value:        logfmtutil.NewSimpleToStringer("Hello from Basic Info"),
// 		Parms:        nil,
// 		PrefixStereo: logfmtutil.DefaultStereo,
// 	}), 3)
// }

// var basicInfoMsg = logfmtgo.NewSprintfStringer(logfmtutil.LogFormat{
// 	Value:        logfmtutil.NewSimpleToStringer("Hello from Basic Info"),
// 	Parms:        nil,
// 	PrefixStereo: logfmtutil.DefaultStereo,
// })

type logPrintWriter struct {
}

func (logOut logPrintWriter) Write(msg string) {
	log.Print(msg)
}

// type logOutWriter struct {
// 	callDepth int
// }

// func (logOut logOutWriter) Write(msg string) {
// 	log.Output(logOut.callDepth, msg)
// }

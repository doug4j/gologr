package logfmtgo

import (
	"fmt"
	"runtime"

	"github.com/doug4j/OmniFocusTools/go/oft/pkg/logapi/v2"
	"github.com/doug4j/OmniFocusTools/go/oft/pkg/logapi/v2/logfmtutil"
)

func NewSrcFmt(format logfmtutil.LogFormat) logapi.LogAwareStringing {
	return srcFormatter{
		format: format,
		skip:   1,
	}
}

func NewSrcFmtWithSkip(format logfmtutil.LogFormat, skip int) logapi.LogAwareStringing {
	return srcFormatter{
		format: format,
		skip:   skip,
	}
}

// srcFormatter represents a formatter that is source code aware
type srcFormatter struct {
	format logfmtutil.LogFormat
	skip   int
}

func (src srcFormatter) LogAwareToString(level logapi.Level, parms ...interface{}) string {
	_, file, line, _ := runtime.Caller(src.skip)
	if src.format.PrefixStereo == logfmtutil.DefaultStereo {
		return src.logWithLastFunc(file, line, src.format.Value.String(), level.String())
	}
	return src.logWithLastFunc(file, line, src.format.Value.String(), src.format.PrefixStereo.String())

}

// func (fmt srcFmt) String() string {
// 	pc, file, line, _ := runtime.Caller(1)
// 	if fmt.PrefixStereo == logfmtutil.DefaultStereo {

// 	}
// }

func (src srcFormatter) logWithLastFunc(file string, line int, msg, prefix string) string {
	// fullPCName := runtime.FuncForPC(pc).Name()
	// lastIndexOfPc := strings.LastIndex(fullPCName, "/") + 1
	// justPcName := fullPCName[lastIndexOfPc:len(fullPCName)]
	// justPcName := fullPCName

	// lastIndexOfFile := strings.LastIndex(file, "/") + 1
	// justFileName := file[lastIndexOfFile:len(file)]
	justFileName := file

	// log.Printf("%s [%s:%d] [%s] %v", prefix, justFileName, line, justPcName, msg)
	return fmt.Sprintf("%s [%s:%d]\n%v", prefix, justFileName, line, msg)

	// return justPcName, justFileName
}

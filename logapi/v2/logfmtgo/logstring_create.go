package logfmtgo

import "github.com/doug4j/OmniFocusTools/go/oft/pkg/logapi/v2"

func NewGoConstLogAwareStringCreating() logapi.LogAwareStringCreating {
	return logAwareStringCreater{}
}

type logAwareStringCreater struct {
	logAwareStringsByPackageAndKey map[string]map[string]logapi.LogAwareStringing
}

func (logAwareStringCreater) Create(key logapi.LogItem, i18nTag string) logapi.LogAwareStringing {
	return nil
}

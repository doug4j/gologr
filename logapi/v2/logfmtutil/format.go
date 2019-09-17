package logfmtutil

import "github.com/doug4j/OmniFocusTools/go/oft/pkg/logapi/v2"

type LogFormat struct {
	Value        logapi.ToString
	PrefixStereo Stereotype
	Parms        []interface{}
}

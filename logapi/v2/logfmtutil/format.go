package logfmtutil

import "github.com/doug4j/gologr/logapi/v2"

type LogFormat struct {
	Value        logapi.ToString
	PrefixStereo Stereotype
	Parms        []interface{}
}

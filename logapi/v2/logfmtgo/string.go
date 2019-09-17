package logfmtgo

import (
	"fmt"

	"github.com/doug4j/gologr/logapi/v2"
	"github.com/doug4j/gologr/logapi/v2/logfmtutil"
)

func NewSprintfStringer(raw logfmtutil.LogFormat) logapi.ToString {
	return sprintfStringer{
		raw: raw,
	}
}

type sprintfStringer struct {
	raw logfmtutil.LogFormat
}

func (str sprintfStringer) String() string {
	val := str.raw.Value.String()
	parms := str.raw.Parms
	return fmt.Sprintf(val, parms)
}

package logfmtutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/doug4j/gologr/logapi/v2"
)

func NewSimpleToStringer(value string) logapi.ToString {
	return simpleToStringer{
		val: value,
	}
}

func NewPosReplaceStringer(raw LogFormat) logapi.ToString {
	return sprintfStringer{
		raw: raw,
	}
}

type simpleToStringer struct {
	val string
}

func (to simpleToStringer) String() string {
	// log.Printf("value is [%v]", to.val)
	return to.val
}

type sprintfStringer struct {
	raw LogFormat
}

func (str sprintfStringer) String() string {
	val := str.raw.Value.String()
	for index, parm := range str.raw.Parms {
		val = strings.ReplaceAll(val, "$"+strconv.Itoa(index+1), fmt.Sprint(parm))
	}
	return val
}

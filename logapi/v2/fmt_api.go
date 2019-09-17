package logapi

// ToString represents a late binding on a string value
// type ToString interface {
// 	String() string
// }

type LogAwareStringing interface {
	LogAwareToString(level Level, parms ...interface{}) string
	// GetLogItem() LogItem
}

type LogAwareStringCreating interface {
	Create(key LogItem, i18nTag string) LogAwareStringing
}

type LogItem struct {
	Group      string
	Key        string
	Stereotype Stereotype
}

// var DefaultLogItem = LogItem{
// 	Key:        "test",
// 	Stereotype: DefaultStereo,
// }

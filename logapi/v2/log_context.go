package logapi

type AppLogContext struct {
	I18nTag       string
	LogLevel      Level
	StringCreator LogAwareStringCreating
	Logr          LogCreatingV2
}

func (app AppLogContext) CreateLogString(key LogItem) LogAwareStringing {
	return app.StringCreator.Create(key, app.I18nTag)
}

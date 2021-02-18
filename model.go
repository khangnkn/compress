package compress

type dummyLog struct {
	ServiceName   string
	APIName       string
	CorrelationID string
	ReturnCode    int32
	ExecutionTime int64
	TemplateID    int32
	Data          string
}

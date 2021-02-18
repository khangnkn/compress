package compress

func serializeWithProtobuf(logs []dummyLog) *Payload {
	var protoLogs = make([]*Log, len(logs))
	for i, log := range logs {
		protoLogs[i] = log2payload(log)
	}
	return &Payload{
		Logs: protoLogs,
	}
}

func log2payload(log dummyLog) (l *Log) {
	return &Log{
		AppName:       log.ServiceName,
		ApiName:       log.APIName,
		CorrelationId: log.CorrelationID,
		ReturnCode:    log.ReturnCode,
		ExecutionTime: log.ExecutionTime,
		TemplateId:    log.TemplateID,
		Data:          log.Data,
	}
}

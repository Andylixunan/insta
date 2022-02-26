package log

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func New() (*Logger, error) {
	l, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &Logger{l.Sugar()}, nil
}

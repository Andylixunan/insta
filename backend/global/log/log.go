package log

import "go.uber.org/zap"

func New() (*zap.SugaredLogger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}

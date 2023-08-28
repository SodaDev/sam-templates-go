package handler

import (
	"context"
	"github.com/Ryanair/gofrlib/log"
)

type LambdaHandler struct {
	loggerConfig log.Configuration
}

func New(loggerConfig log.Configuration) *LambdaHandler {
	return &LambdaHandler{loggerConfig: loggerConfig}
}

func (lh *LambdaHandler) Handle(ctx context.Context) error {
	log.Init(lh.loggerConfig)

	return nil
}

package event

import (
	"context"
	"time"
)

type BaseGrpc struct {
	//
}

func (event *BaseGrpc) Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	timeout := time.Second * 30
	if len(customTimeout) > 0 {
		timeout = customTimeout[0]
	}

	return context.WithTimeout(context.Background(), timeout)
}

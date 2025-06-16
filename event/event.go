package event

import (
	"context"
	"test/logger"
)

type EventCenter interface {
	SendMsg(ctx context.Context, msg Event) error
	GetMsg(ctx context.Context) (Event, error)
}

type Event interface {
}

func NewDefaultEventCenter() EventCenter {
	return globalMemoryEventCenter
}

var globalMemoryEventCenter = &MemoryEventCenter{
	events: make(chan Event, 10),
}

type MemoryEventCenter struct {
	events chan Event
}

func (m *MemoryEventCenter) SendMsg(ctx context.Context, event Event) error {
	fun := "MemoryEventCenter.SendMsg"

	select {
	case m.events <- event:
	default:
		logger.Ex(ctx, "%s there are too much msgs to handler", fun)
		m.events <- event
	}
	return nil
}

func (m *MemoryEventCenter) GetMsg(ctx context.Context) (Event, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case event := <-m.events:
		return event, nil
	}
}

package domainEvent

import (
	"lifelogger/server/util/mattermost"

	"context"
	"errors"
	"fmt"
	"runtime/debug"
)

var Subscribers map[string]DomainEventHandler = map[string]DomainEventHandler{}

type (
	DomainEvent interface {
		Name() string
	}

	DomainEventHandler interface {
		Handle(event DomainEvent)
	}

	EventProcessor interface {
		GetCtx() context.Context
		SetCtx(ctx context.Context)
		Publish(event DomainEvent)
		ProcessAndRemove()
	}

	eventProcessorImpl struct {
		ctx context.Context
	}
)

func NewEventProcessor() EventProcessor {
	return new(eventProcessorImpl)
}

func (p *eventProcessorImpl) GetCtx() context.Context {
	return p.ctx
}

func (p *eventProcessorImpl) SetCtx(ctx context.Context) {
	p.ctx = ctx
}

func (p *eventProcessorImpl) Publish(event DomainEvent) {
	events := []DomainEvent{}
	value := p.ctx.Value("events")
	if value != nil {
		events = value.([]DomainEvent)
	}

	events = append(events, event)
	p.ctx = context.WithValue(p.ctx, "events", events)
}

func (p *eventProcessorImpl) ProcessAndRemove() {
	var events []DomainEvent
	value := p.ctx.Value("events")
	if value != nil {
		events = value.([]DomainEvent)

		for _, event := range events {
			handler := Subscribers[event.Name()]
			if handler == nil {
				panic(errors.New("subscribe " + event.Name()))
			}
			go func(event DomainEvent) {
				defer func() {
					if err := recover(); err != nil {
						mattermost.Noti(
							event.Name(),
							fmt.Sprintf("data: %v\nerror: %v\nstackTrace: %v",
								event,
								err,
								string(debug.Stack()),
							),
						)
					}
				}()

				handler.Handle(event)
			}(event)
		}

		p.ctx = context.WithValue(p.ctx, "events", nil)
	}
}

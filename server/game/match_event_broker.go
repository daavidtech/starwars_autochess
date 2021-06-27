package game

import "context"

type MatchEvent struct {
	EventID string
}

type MatchEventBroker struct {
	pubch   chan MatchEvent
	subch   chan chan MatchEvent
	unsubch chan chan MatchEvent
	ctx     context.Context
}

func NewMatchEventBroker(ctx context.Context) MatchEventBroker {
	return MatchEventBroker{
		pubch: make(chan MatchEvent),
		ctx:   ctx,
	}
}

func (eventBus *MatchEventBroker) Run() {
	subs := map[chan MatchEvent]struct{}{}

	for {
		select {
		case <-eventBus.ctx.Done():
			return
		case ch := <-eventBus.subch:
			subs[ch] = struct{}{}
		case ch := <-eventBus.unsubch:
			delete(subs, ch)
		case matchEvent := <-eventBus.pubch:
			for sub := range subs {
				select {
				case sub <- matchEvent:
				default:
				}
			}
		}
	}
}

func (eventBus *MatchEventBroker) publishEvent(gameEvent MatchEvent) {
	eventBus.pubch <- gameEvent
}

func (eventBus *MatchEventBroker) Subscribe(matchID string) <-chan MatchEvent {
	ch := make(chan MatchEvent)
	return ch
}

func (eventBus *MatchEventBroker) Unsubscribe(ch chan MatchEvent) {
	eventBus.unsubch <- ch
	close(ch)
}

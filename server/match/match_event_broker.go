package match

import (
	"context"
	"log"
)

type UnitBought struct {
}

type BarrackUnitAdded struct {
	UnitID     string
	UnitType   string
	Rank       int
	HP         int
	Mana       int
	AttackRate int
}

type BarrackUnitRemoved struct {
	UnitID string
}

type BarrackUnitUpgraded struct {
	UnitID     string
	UnitType   string
	Tier       int
	Rank       int
	HP         int
	Mana       int
	AttackRate int
}

type ShopRefilled struct {
	ShopUnits []ShopUnit
}

type MatchEvent struct {
	EventID             string
	UnitBought          *UnitBought
	BarrackUnitAdded    *BarrackUnitAdded
	BarrackUnitRemoved  *BarrackUnitRemoved
	BarrackUnitUpgraded *BarrackUnitUpgraded
	ShopRefilled        *ShopRefilled
}

type MatchEventBroker struct {
	pubch   chan MatchEvent
	subch   chan chan MatchEvent
	unsubch chan chan MatchEvent
	ctx     context.Context
}

func NewMatchEventBroker(ctx context.Context) *MatchEventBroker {
	return &MatchEventBroker{
		pubch:   make(chan MatchEvent),
		subch:   make(chan chan MatchEvent),
		unsubch: make(chan chan MatchEvent),
		ctx:     ctx,
	}
}

func (eventBus *MatchEventBroker) Run() {
	subs := map[chan MatchEvent]struct{}{}

	for {
		select {
		case <-eventBus.ctx.Done():
			log.Println("Turning event broker off")

			return
		case ch := <-eventBus.subch:
			log.Println("Run sub")

			subs[ch] = struct{}{}
		case ch := <-eventBus.unsubch:
			log.Println("Run unsub")

			delete(subs, ch)
		case matchEvent := <-eventBus.pubch:
			log.Printf("Run publish matchEvent %v", matchEvent)

			for sub := range subs {
				log.Println("Publishing event to subscriber")

				select {
				case sub <- matchEvent:
				default:
				}
			}
		}
	}
}

func (eventBus *MatchEventBroker) publishEvent(matchEvents ...MatchEvent) {
	log.Printf("Publish matchEvent")

	for _, matchEvent := range matchEvents {
		eventBus.pubch <- matchEvent
	}
}

func (eventBus *MatchEventBroker) Subscribe(matchID string) <-chan MatchEvent {
	log.Println("MatchEventBroker Subscribe")

	ch := make(chan MatchEvent)

	eventBus.subch <- ch

	return ch
}

func (eventBus *MatchEventBroker) Unsubscribe(ch chan MatchEvent) {
	log.Println("MatchEventBroker Unsubscribe")

	eventBus.unsubch <- ch
	close(ch)
}

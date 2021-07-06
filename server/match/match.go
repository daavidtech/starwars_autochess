package match

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MatchPhase string

const (
	InitPhase      MatchPhase = "InitPhase"
	LobbyPhase     MatchPhase = "LobbyPhase"
	ShoppingPhase  MatchPhase = "ShoppingPhase"
	PlacementPhase MatchPhase = "PlacementPhase"
	BattlePhase    MatchPhase = "BattlePhase"
)

type Match struct {
	ctx context.Context

	id                 string
	currentRoundNumber int
	phase              MatchPhase
	players            map[string]*Player

	TierProbabilities *TierProbabilities
	UnitPropertyStore *UnitPropertyStore

	eventBroker *MatchEventBroker

	mu sync.Mutex
}

func NewMatch(ctx context.Context) *Match {
	eventBroker := NewMatchEventBroker(ctx)

	go eventBroker.Run()

	return &Match{
		ctx: ctx,

		id:          uuid.New().String(),
		eventBroker: eventBroker,
		players:     make(map[string]*Player),
		phase:       InitPhase,
	}
}

func (match *Match) GetEventBroker() *MatchEventBroker {
	return match.eventBroker
}

func (match *Match) BuyUnit(playerID string, id int) {
	match.mu.Lock()
	defer match.mu.Unlock()

	player := match.players[playerID]

	if player.IsBarrackFull() {
		log.Println("Barrack is full")

		return
	}

	shopUnit := player.shop.Pick(id)

	events := player.AddShopUnit(shopUnit)

	events = append(events, MatchEvent{
		ShopUnitRemoved: &ShopUnitRemoved{
			ShopUnitID: id,
		},
	})

	match.eventBroker.publishEvent(events...)
}

func (match *Match) SellUnit(playerID string, unitID string) {

}

func (match *Match) PlaceUnit(playerID string, unitID string, x int, y int) {
	match.mu.Lock()
	defer match.mu.Unlock()

	player := match.players[playerID]

	unit := player.battleUnits[unitID]

	unit.placement = &Placement{
		x: x,
		y: y,
	}

	match.eventBroker.publishEvent(MatchEvent{
		UnitPlaced: &UnitPlaced{
			UnitID: unitID,
			X:      x,
			Y:      y,
		},
	})
}

func (match *Match) BuyLevelUp(playerID string) {

}

func (match *Match) RecycleShopUnits(playerID string) {

}

func (match *Match) CreateSnapshop() MatchSnapshot {
	return MatchSnapshot{}
}

func (match *Match) CreatePlayer() *Player {
	newPlayer := NewPlayer()

	newPlayer.shop.UnitPropertyStore = match.UnitPropertyStore
	newPlayer.shop.TierProbabilities = match.TierProbabilities

	match.players[newPlayer.GetID()] = newPlayer

	return newPlayer
}

func (match *Match) SetPhase(phase MatchPhase) {
	match.phase = phase
}

func (match *Match) GetID() string {
	return match.id
}

func (match *Match) CountPlayers() int {
	return len(match.players)
}

func (match *Match) IsFull() bool {
	return len(match.players) > 7
}

func (match *Match) moveToShoppingPhase() {
	match.mu.Lock()
	defer match.mu.Unlock()

	match.phase = ShoppingPhase

	match.eventBroker.publishEvent(MatchEvent{
		PhaseChanged: &PhaseChanged{
			MatchPhase: ShoppingPhase,
		},
	})

	for _, player := range match.players {
		shopRefilled := player.shop.Fill(player.GetLevel())

		match.eventBroker.publishEvent(MatchEvent{
			ShopRefilled: &shopRefilled,
		})
	}
}

func (match *Match) moveToPlacementPhase() {
	match.mu.Lock()
	defer match.mu.Unlock()

	match.phase = PlacementPhase

	match.eventBroker.publishEvent(MatchEvent{
		PhaseChanged: &PhaseChanged{
			MatchPhase: PlacementPhase,
		},
	})

	match.eventBroker.publishEvent(MatchEvent{
		CountdownStarted: &CountdownStarted{
			StartValue: 10,
			Interval:   1.0,
		},
	})
}

func (match *Match) moveToBattlePhase() {
	match.mu.Lock()
	defer match.mu.Unlock()

	match.phase = BattlePhase

	match.eventBroker.publishEvent(MatchEvent{
		PhaseChanged: &PhaseChanged{
			MatchPhase: BattlePhase,
		},
	})
}

func (match *Match) Run() {
	<-time.NewTimer(500 * time.Millisecond).C

	match.eventBroker.publishEvent(MatchEvent{
		PhaseChanged: &PhaseChanged{
			MatchPhase: InitPhase,
		},
	})

	<-time.NewTimer(2 * time.Second).C

	match.eventBroker.publishEvent(MatchEvent{
		PhaseChanged: &PhaseChanged{
			MatchPhase: LobbyPhase,
		},
	})

	match.eventBroker.publishEvent(MatchEvent{
		CountdownStarted: &CountdownStarted{
			StartValue: 2,
			Interval:   1.0,
		},
	})

	select {
	case <-time.NewTimer(2 * time.Second).C:
	case <-match.ctx.Done():
		return
	}

	log.Printf("Refilling shop")

	for {
		match.moveToShoppingPhase()

		match.eventBroker.publishEvent(MatchEvent{
			CountdownStarted: &CountdownStarted{
				StartValue: 10,
				Interval:   1.0,
			},
		})

		<-time.NewTimer(10 * time.Second).C

		match.moveToPlacementPhase()

		<-time.NewTimer(10 * time.Second).C

		match.moveToBattlePhase()
	}
}

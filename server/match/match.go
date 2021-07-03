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
	Lobby          MatchPhase = "Lobby"
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

	shop              Shop
	TierProbabilities TierProbabilities

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
	}
}

func (match *Match) GetEventBroker() *MatchEventBroker {
	return match.eventBroker
}

func (match *Match) BuyUnit(playerID string, index int) {
	match.mu.Lock()
	defer match.mu.Unlock()

	// match.shop.Pick(1)

	player := match.players[playerID]

	if player.IsBarrackFull() {
		log.Println("Barrack is full")

		return
	}

	events := player.AddShopUnit(ShopUnit{
		UnitType:   "unit_droid",
		Tier:       1,
		HP:         100,
		Mana:       100,
		AttackRate: 1,
	})

	match.eventBroker.publishEvent(events...)
}

func (match *Match) SellUnit(playerID string, unitID string) {

}

func (match *Match) PlaceUnit(playerID string, unitID string, x int, y int) {

}

func (match *Match) BuyLevelUp(playerID string) {

}

func (match *Match) RecycleShopUnits(playerID string) {

}

func (match *Match) CreatePlayer() *Player {
	newPlayer := NewPlayer()

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

func (match *Match) Run() {
	<-time.NewTimer(1 * time.Second).C

	log.Printf("Refilling shop")

	match.eventBroker.publishEvent(MatchEvent{
		ShopRefilled: &ShopRefilled{
			ShopUnits: []ShopUnit{
				ShopUnit{
					UnitType: "unit_clone",
				},
			},
		},
	})

	for {
		select {
		case <-match.ctx.Done():
			return
		}
	}
}

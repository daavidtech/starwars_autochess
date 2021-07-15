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
	LobbyPhase     MatchPhase = "LobbyPhase"
	ShoppingPhase  MatchPhase = "ShoppingPhase"
	PlacementPhase MatchPhase = "PlacementPhase"
	BattlePhase    MatchPhase = "BattlePhase"
	EndPhase       MatchPhase = "EndPhase"
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
		phase:       LobbyPhase,
	}
}

func (match *Match) GetPlayers() []*Player {
	players := []*Player{}

	for _, player := range match.players {
		players = append(players, player)
	}

	return players
}

func (match *Match) GetMatchPhase() MatchPhase {
	return match.phase
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
			PlayerID:   playerID,
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

	unit := player.units[unitID]

	unit.Placement = &Point{
		X: float32(x),
		Y: float32(y),
	}

	match.eventBroker.publishEvent(MatchEvent{
		UnitPlaced: &UnitPlaced{
			PlayerID: playerID,
			UnitID:   unitID,
			X:        x,
			Y:        y,
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

func (match *Match) CreatePlayer(name string, lobbyAdmin bool) *Player {
	newPlayer := NewPlayer()

	newPlayer.name = name
	newPlayer.lobbyAdmin = lobbyAdmin

	newPlayer.shop.UnitPropertyStore = match.UnitPropertyStore
	newPlayer.shop.TierProbabilities = match.TierProbabilities
	newPlayer.unitPropStore = match.UnitPropertyStore

	match.players[newPlayer.GetID()] = newPlayer

	match.eventBroker.publishEvent(MatchEvent{
		PlayerJoined: &PlayerJoined{
			Player: *newPlayer,
		},
	})

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

		shopRefilled.PlayerID = player.GetID()

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
			StartValue: 5,
			Interval:   1.0,
		},
	})
}

func (match *Match) moveToBattlePhase() bool {
	match.mu.Lock()
	defer match.mu.Unlock()

	match.phase = BattlePhase

	match.eventBroker.publishEvent(MatchEvent{
		PhaseChanged: &PhaseChanged{
			MatchPhase: BattlePhase,
		},
	})

	rounds := []*Round{}

	playerPool := copyPlayersWithoutDead(match.players)

	for len(playerPool) > 0 {
		player1 := popPlayer(playerPool)

		player2 := popPlayer(playerPool)

		if player2 == nil {
			player2 = picRandomPlayer(match.players)
		}

		battleUnits := []*BattleUnit{}

		for _, unit := range player1.units {
			if unit.Placement != nil {
				battleUnits = append(battleUnits, createBattleUnit(unit, 1, player1.id))
			}
		}

		for _, unit := range player2.units {
			if unit.Placement != nil {
				battleUnits = append(battleUnits, createBattleUnit(unit, 2, player2.id))
			}
		}

		round := CreateRound(match.ctx, match.eventBroker, battleUnits)

		round.player1ID = player1.id
		round.player2ID = player2.id

		roundCreated := RoundCreated{
			PlayerID: round.player1ID,
			Units:    []BattleUnit{},
		}

		for _, battleUnit := range battleUnits {
			roundCreated.Units = append(roundCreated.Units, *battleUnit)
		}

		rounds = append(rounds, round)

		match.eventBroker.publishEvent(MatchEvent{
			RoundCreated: &roundCreated,
		})

		if round.player1ID != round.player2ID {
			roundCreated2 := RoundCreated{
				PlayerID: round.player2ID,
				Units:    []BattleUnit{},
			}

			for _, battleUnit := range battleUnits {
				b := *battleUnit

				b.Y = float32(invertY(b.Y))

				roundCreated2.Units = append(roundCreated2.Units, b)
			}

			match.eventBroker.publishEvent(MatchEvent{
				RoundCreated: &roundCreated2,
			})
		}
	}

	for _, round := range rounds {
		log.Printf("Starting round %v", round.id)

		result := round.run()

		log.Printf("%v round finished", round.id)

		player := match.players[round.player1ID]

		player.payDay()

		if result.whoWon == 2 {
			player.health -= 10

			if player.health <= 0 {
				player.dead = true
			}
		}

		roundFinished := RoundFinished{
			PlayerID:         round.player1ID,
			NewCreditsAmount: player.credits,
			NewPlayerHealth:  player.health,
		}

		for _, unit := range player.units {
			roundFinished.Units = append(roundFinished.Units, *unit)
		}

		match.eventBroker.publishEvent(MatchEvent{
			RoundFinished: &roundFinished,
		})

		if round.player2ID != round.player1ID {
			player2 := match.players[round.player2ID]

			player2.payDay()

			if result.whoWon == 1 {
				player2.health -= 10

				if player2.health <= 0 {
					player2.dead = true
				}
			}

			roundFinished2 := RoundFinished{
				PlayerID:         round.player2ID,
				NewCreditsAmount: player2.credits,
				NewPlayerHealth:  player2.health,
			}

			for _, unit := range player2.units {
				roundFinished2.Units = append(roundFinished2.Units, *unit)
			}

			match.eventBroker.publishEvent(MatchEvent{
				RoundFinished: &roundFinished2,
			})
		}
	}

	if match.CountAlivePlayers() < 1 {
		log.Println("Match ended")

		match.eventBroker.publishEvent(MatchEvent{
			PhaseChanged: &PhaseChanged{
				MatchPhase: EndPhase,
			},
		})

		return false
	}

	return true
}

func (match *Match) CountAlivePlayers() int {
	count := 0

	for _, player := range match.players {
		if player.dead {
			continue
		}

		count += 1
	}

	return count
}

func (match *Match) StartLobby() {
	match.mu.Lock()
	defer match.mu.Unlock()

	match.phase = LobbyPhase

	match.eventBroker.publishEvent(MatchEvent{
		PhaseChanged: &PhaseChanged{
			MatchPhase: LobbyPhase,
		},
	})

}

func (match *Match) Start(playerID string) {
	match.mu.Lock()
	defer match.mu.Unlock()

	if match.phase != LobbyPhase {
		log.Println("Cannot start match if not in lobby phase")

		return
	}

	if len(match.players) < 1 {
		log.Println("Too few players to start the game")

		return
	}

	player := match.players[playerID]

	if !player.lobbyAdmin {
		log.Println("Player is not lobby admin cannot start match")

		return
	}

	match.phase = ShoppingPhase

	go match.Run()
}

func (match *Match) Run() {
	log.Printf("Refilling shop")

	gameContinue := true

	for gameContinue {
		match.moveToShoppingPhase()

		match.eventBroker.publishEvent(MatchEvent{
			CountdownStarted: &CountdownStarted{
				StartValue: 5,
				Interval:   1.0,
			},
		})

		<-time.NewTimer(5 * time.Second).C

		match.moveToPlacementPhase()

		<-time.NewTimer(5 * time.Second).C

		gameContinue = match.moveToBattlePhase()
	}
}

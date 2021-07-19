package game

import (
	"context"
	"log"

	"github.com/daavidtech/starwars_autochess/match"
)

type GameCoordinator struct {
	ctx     context.Context
	matches map[string]*match.Match

	UnitPropertyStore *match.UnitPropertyStore
	TierProbabilities *match.TierProbabilities
}

func NewGameCoordinator(ctx context.Context) *GameCoordinator {
	return &GameCoordinator{
		ctx:     ctx,
		matches: make(map[string]*match.Match),
	}
}

func (gameCoordinator *GameCoordinator) FindNewMatch() *match.Match {
	for matchID, m := range gameCoordinator.matches {
		if m.IsFull() {
			continue
		}

		if m.GetMatchPhase() == match.EndPhase {
			delete(gameCoordinator.matches, matchID)
			continue
		}

		log.Println("Found existing match")

		return m
	}

	log.Println("No existing match found creating new")

	newMatch := match.NewMatch(gameCoordinator.ctx)

	newMatch.UnitPropertyStore = gameCoordinator.UnitPropertyStore
	newMatch.TierProbabilities = gameCoordinator.TierProbabilities

	gameCoordinator.matches[newMatch.GetID()] = newMatch

	return newMatch
}

func (gameCoordinator *GameCoordinator) FindMatch(matchID string) *match.Match {
	return gameCoordinator.matches[matchID]
}

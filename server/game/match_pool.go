package game

import "github.com/google/uuid"

type MatchPool struct {
	matches []*Match
}

func (matchPool *MatchPool) findMatch() *Match {
	for _, match := range matchPool.matches {
		if match.IsFull() {
			continue
		}

		return match
	}

	newMatch := &Match{
		id: uuid.New().String(),
	}

	matchPool.matches = append(matchPool.matches, newMatch)

	return newMatch
}

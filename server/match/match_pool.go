package match

type MatchPool struct {
	matches map[string]Match
}

func (matchPool *MatchPool) FindNewMatch() *Match {
	for _, match := range matchPool.matches {
		if match.IsFull() {
			continue
		}

		return match
	}

	newMatch := NewMatch()

	matchPool.matches[newMatch.id] = newMatch

	return newMatch
}

func (matchPool *MatchPool) FindMatch(matchID string) *Match {
	return matchPool.matches[matchID]
}

package networking

import "github.com/daavidtech/starwars_autochess/match"

func convertMatchPhase(matchPhase match.MatchPhase) GamePhase {
	switch matchPhase {
	case match.LobbyPhase:
		return LobbyPhase
	case match.ShoppingPhase:
		return ShoppingPhase
	case match.PlacementPhase:
		return PlacementPhase
	case match.BattlePhase:
		return BattlePhase
	case match.EndPhase:
		return EndPhase
	}

	return LoginPhase
}

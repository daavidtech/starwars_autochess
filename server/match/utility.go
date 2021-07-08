package match

import "math"

func calcDist(x, y, x2, y2 float32) float32 {
	return float32(math.Max(float64(x2-x), float64(y2-y)))
}

func copyPlayers(players map[string]*Player) map[string]*Player {
	newPlayers := make(map[string]*Player, len(players))

	for id, player := range players {
		newPlayers[id] = player
	}

	return newPlayers
}

func popPlayer(players map[string]*Player) *Player {
	for id, player := range players {
		delete(players, id)

		return player
	}

	return nil
}

func picRandomPlayer(players map[string]*Player) *Player {
	for _, player := range players {
		return player
	}

	return nil
}

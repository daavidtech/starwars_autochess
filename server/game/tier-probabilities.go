package game

type TierProbabilities struct {
	probabilities [][]int
}

func NewTierProbabilities(probabilities [][]int) TierProbabilities {
	return TierProbabilities{
		probabilities: probabilities,
	}
}

func (t *TierProbabilities) PickLevel(level int) []int {

	return []int{}
}

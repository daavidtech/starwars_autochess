package match

type TierProbabilities struct {
	probabilities [][]int
}

func NewTierProbabilities(probabilities [][]int) TierProbabilities {
	return TierProbabilities{
		probabilities: probabilities,
	}
}

func (t *TierProbabilities) PickLevel(level int) []int {
	if level < 1 || level > 9 {
		return []int{}
	}

	return t.probabilities[level-1]
}

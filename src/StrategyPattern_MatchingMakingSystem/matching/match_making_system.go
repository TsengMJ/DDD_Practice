package matching

type MatchingSystem struct {
	matchingAlgo MatchingStrategy
}

func (ms *MatchingSystem) SetMatchingAlog(algo MatchingStrategy) {
	ms.matchingAlgo = algo
}

func (ms *MatchingSystem) Match(me Individual, others []Individual) Individual {
	return ms.matchingAlgo.matching(me, others)
}

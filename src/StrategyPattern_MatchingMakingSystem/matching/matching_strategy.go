package matching

type MatchingStrategy interface {
	matching(Individual, []Individual) Individual
}

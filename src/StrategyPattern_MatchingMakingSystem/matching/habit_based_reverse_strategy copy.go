package matching

import "math"

type HabitBasedReverseStrategy struct {
}

func (s *HabitBasedStrategy) matching(me Individual, others []Individual) Individual {
	var matchPerson Individual
	maxOverlap := math.MaxInt

	for _, p := range others {
		count := 0
		for key := range me.habits {
			if p.habits[key] {
				count += 1
			}
		}

		if count < maxOverlap {
			maxOverlap = count
			matchPerson = p
		}
	}

	return matchPerson
}

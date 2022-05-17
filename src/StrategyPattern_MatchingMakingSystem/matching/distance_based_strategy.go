package matching

import "math"

type DistanceBasedStrategy struct {
}

func (s *DistanceBasedStrategy) matching(me Individual, others []Individual) Individual {
	var matchPerson Individual
	dist := math.MaxFloat32

	for _, p := range others {
		distX := math.Pow(float64(me.coord[0]-p.coord[0]), 2)
		distY := math.Pow(float64(me.coord[1]-p.coord[1]), 2)
		tmpDist := math.Sqrt(distX + distY)

		if tmpDist < dist {
			dist = tmpDist
			matchPerson = p
		}
	}

	return matchPerson
}

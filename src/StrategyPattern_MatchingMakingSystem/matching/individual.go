package matching

type Individual struct {
	id     int
	gender string
	age    int
	intro  string
	habits map[string]bool
	coord  []float32
}

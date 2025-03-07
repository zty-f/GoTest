package felidae

import "go-pprof-practice/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}

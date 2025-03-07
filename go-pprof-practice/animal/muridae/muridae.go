package muridae

import "go-pprof-practice/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}

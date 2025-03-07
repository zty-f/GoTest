package canidae

import "go-pprof-practice/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}

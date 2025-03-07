package animal

import (
	"go-pprof-practice/animal/canidae/dog"
	"go-pprof-practice/animal/canidae/wolf"
	"go-pprof-practice/animal/felidae/cat"
	"go-pprof-practice/animal/felidae/tiger"
	"go-pprof-practice/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},

		&cat.Cat{},
		&tiger.Tiger{},

		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}

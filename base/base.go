package base

import "fmt"

func m1() {
	a := 1
	if a = 4; false {

	} else if b := 2; false {

	} else {
		println(a, b)
	}
	println(a)
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func m2() {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

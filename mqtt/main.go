package main

import (
	mqt "test/mqtt/mqtt"
	"time"
)

func main() {
	go mqt.ConsumerPoint()
	go mqt.ProducerPoint()
	time.Sleep(30 * time.Second)
}

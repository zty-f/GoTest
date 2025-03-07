package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"go-pprof-practice/animal"
)

/*
go tool pprof “http://localhost:6060/debug/pprof/profile?seconds=10
top \ list
go tool pprof -http=:8082 "http://localhost:6060/debug/pprof"
go tool pprof -http=:8082 "http://localhost:6060/debug/pprof/heap"
go tool pprof -http=:8082 "http://localhost:6060/debug/pprof/goroutine"
go tool pprof -http=:8082 "http://localhost:6060/debug/pprof/mutex"
go tool pprof -http=:8082 "http://localhost:6060/debug/pprof/block"
*/

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range animal.AllAnimals {
			v.Live()
		}
		time.Sleep(time.Second)
	}
}

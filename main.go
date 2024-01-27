package main

import (
	"fmt"
	"log"
	"os"
	. "randomWalk/saw"
	. "randomWalk/saw/algorithms"
	"strconv"
	"time"
)

func main() {

	n := 100
	var err error
	if len(os.Args) >= 2 {
		if n, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatalf("Invalid walk length")
		}
	}

	procs := 10
	if len(os.Args) >= 3 {
		if procs, err = strconv.Atoi(os.Args[2]); err != nil {
			log.Fatalf("Invalid process count")
		}
	}

	walker := ZombieWalkers["mapless"]
	if len(os.Args) >= 4 {
		if testWalker, ok := ZombieWalkers[os.Args[3]]; ok {
			walker = testWalker
		}
	}

	if len(os.Args) > 4 {
		log.Fatal("Too many arguments")
	}

	//fmt.Printf("%v, %v, %v", n, procs, walker)

	fun := walker
	if procs > 1 {
		fun = func(n int) *[]Vec2 {
			return ZombiePool(n, walker, procs)
		}
	}

	start := time.Now()
	pos := fun(n)

	WriteJSON(*pos, fmt.Sprintf("output/sa%v", n))

	elapsed := time.Since(start)
	log.Printf("Search took %s", elapsed)
}

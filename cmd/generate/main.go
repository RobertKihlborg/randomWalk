package main

import (
	"fmt"
	"log"
	"os"
	"randomWalk/pkg"
	"strconv"
	"time"
)

func main() {

	argCount := len(os.Args)
	if argCount < 2 {
		log.Fatalf("missing builder name")
	}

	walker, name, walkerArgs, err := pkg.BuildWalker(os.Args[1:])
	fmt.Println(walkerArgs)
	if err != nil {
		log.Fatal(err)
	}
	if argCount <= walkerArgs+1 {
		log.Fatal("Too few arguments")
	}

	var n = 0
	if n, err = strconv.Atoi(os.Args[walkerArgs+1]); err != nil {
		log.Fatalf("Invalid walk length")
	}
	//usedArgs++

	start := time.Now()
	walk := walker(n)

	if err = WriteWalk(walk, fmt.Sprint(name, n)); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Search took %s", elapsed)
}

func WriteWalk(walk pkg.Walk, name string) error {
	file, err := os.Create(fmt.Sprintf("%v.walk", name))
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Fatal("Error closing file:", cerr)
		}
	}()

	_, err = file.WriteString(fmt.Sprint(walk))
	if err != nil {
		return err
	}

	return nil
}

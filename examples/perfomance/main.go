package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/hand/effective"
	"github.com/dnovikoff/tempai-core/hand/shanten"
	"github.com/dnovikoff/tempai-core/hand/tempai"
	"github.com/dnovikoff/tempai-core/tile"
)

func testShanten() {
	const repeat = 100000
	data := make([]compact.Instances, repeat)
	// prepare
	source := rand.NewSource(123)
	rnd := rand.New(source)
	instances := compact.NewAllInstances().Instances()
	for k := range data {
		tile.Shuffle(instances, rnd)
		data[k] = compact.NewInstances().Add(instances[:13])
	}
	start := time.Now()
	for _, v := range data {
		shanten.Calculate(v)

	}
	elapsed := time.Since(start)
	fmt.Println("================== Test shanten")
	fmt.Printf("Repeat: %v\n", repeat)
	fmt.Printf("Elapsed: %v\n", elapsed)
	fmt.Printf("Estemated speed: %v per second\n", repeat/elapsed.Seconds())
}

func testTempai() {
	const repeat = 100000
	data := make([]compact.Instances, repeat)
	// prepare
	source := rand.NewSource(123)
	rnd := rand.New(source)
	instances := compact.NewAllInstancesFromTo(tile.Sou1, tile.Sou9+1).Instances()
	for k := range data {
		tile.Shuffle(instances, rnd)
		data[k] = compact.NewInstances().Add(instances[:13])
	}
	cnt := 0
	start := time.Now()
	for _, v := range data {
		r := tempai.Calculate(v)
		if r != nil {
			cnt++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("================== Test tempai")
	fmt.Printf("Repeat: %v\n", repeat)
	fmt.Printf("Elapsed: %v\n", elapsed)
	fmt.Printf("Estemated speed: %v per second\n", repeat/elapsed.Seconds())
	fmt.Printf("Tempai hand count: %v\n", cnt)
}

func testEffective() {
	const repeat = 10000
	data := make([]compact.Instances, repeat)
	// prepare
	source := rand.NewSource(123)
	rnd := rand.New(source)
	instances := compact.NewAllInstances().Instances()
	for k := range data {
		tile.Shuffle(instances, rnd)
		data[k] = compact.NewInstances().Add(instances[:14])
	}
	start := time.Now()
	for _, v := range data {
		effective.Calculate(v)

	}
	elapsed := time.Since(start)
	fmt.Println("================== Test effectivity")
	fmt.Printf("Repeat: %v\n", repeat)
	fmt.Printf("Elapsed: %v\n", elapsed)
	fmt.Printf("Estemated speed: %v per second\n", repeat/elapsed.Seconds())
}

func main() {
	testShanten()
	testTempai()
	testEffective()
}

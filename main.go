package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const DEFAULT_MAP_SIZE = 10

func main() {
	fmt.Println("Let's find the way out here!")
	fmt.Println("")
	fmt.Println("Want a same map? Just give a seed like this: ./go-find-the-way -seed=5675435675")
	fmt.Println("Adjust the map size, like 15x15: ./go-find-the-way --size=15")
	fmt.Println("")

	mapSize := flag.Int("size", DEFAULT_MAP_SIZE, "Map size n*n")
	seed := flag.Int64("seed", time.Now().Unix(), "Map seed")
	flag.Parse()

	fmt.Printf("Current seed: %d\n", *seed)
	rand.Seed(*seed)

	fmt.Printf("Map size: %d\n", *mapSize)
	area := make([][]int, *mapSize)

	for row := 0; row < *mapSize; row++ {
		area[row] = make([]int, *mapSize)

		for col := 0; col < *mapSize; col++ {
			area[row][col] = rand.Intn(2)
		}
	}

	for row := 0; row < *mapSize; row++ {
		fmt.Println(area[row])
	}
}

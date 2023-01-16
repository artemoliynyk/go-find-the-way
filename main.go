package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	// "math/rand"
	"time"

	screenHelper "github.com/artemoliynyk/go-find-the-way/screen"

	"github.com/gdamore/tcell/v2"
)

const DEFAULT_MAP_SIZE = 10

var screen tcell.Screen;
var style tcell.Style;

func printIntro(seed int64, size int) {
	fmt.Println("Let's find the way out here!")
	fmt.Println("")
	fmt.Println("Want a same map? Just give a seed like this: ./go-find-the-way -seed=5675435675")
	fmt.Println("Adjust the map size, like 15x15: ./go-find-the-way --size=15")
	fmt.Println("")
	fmt.Printf("Current seed: %d\n", seed)
	fmt.Printf("Map size: %d\n", size)
	fmt.Println("")
}

func main() {

	quiet := flag.Bool("q", false, "No into text")
	mapSize := flag.Int("size", DEFAULT_MAP_SIZE, "Map size n*n")
	mapSeed := flag.Int64("seed", time.Now().Unix(), "Map seed")
	flag.Parse()

	if !*quiet {
		printIntro(*mapSeed, *mapSize)
	}

	screen, style = screenHelper.InitScreen()

	// seed our map
	rand.Seed(*mapSeed)

	// make map
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

	screenHelper.DrawText(screen, 0, 0, 10, 10, style, "Test")

	quit := func() {
		screen.Fini()
		os.Exit(0)
	}
	for {
		// Update screen
		screen.Show()
	
		// Poll event
		ev := screen.PollEvent()
	
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}

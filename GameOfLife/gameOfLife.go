// IMPORTANT //
/*
	After program has been started. It reads the first line from file(path : "rules.txt") which contains parametres.
	You have 5 seconds to change rules (Sleep time(can be changed) between iterations).
  Parameters should be separated by comma without spaces!
	To change rules you need to change the numbers in file (the parameters concerning the rules are signed in rules.txt)
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Initial conditions
var (
	rulesFile      = "rules.txt"
	paramLineNum   = 0
	xAxis          = 40 // The width of the grid
	yAxis          = 40 // The height of the grid
	iterationCount = 10 // Number of iterations
	sleepTime      = 5  // Sleep time between iterations in seconds
	initLiveCells  = 5  // Percentage of game board with living cells as initial
)

type Game struct {
	generation   int
	xSize, ySize int
	cells        []bool
}

//initializes a new game with an empty grid
func NewGame(x, y int) *Game {
	cells := make([]bool, x*y)
	return &Game{generation: 0, xSize: x, ySize: y, cells: cells}
}

//sets a given percentage of the cells in the grid to "alive".
func (gb *Game) RandInit(percentage int) {

	//Calculate number of living cells
	numAlive := percentage * len(gb.cells) / 100

	//Insert living cells at the beginning
	for i := 0; i < numAlive; i++ {
		gb.cells[i] = true
	}

	vals := gb.cells
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(vals); n > 0; n-- {
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
	}

	gb.cells = vals
}

// Reading parametres for the rules from file
func ReadLine(lineNum int) []int {
	values := make([]int, 0)

	f, err := os.Open(rulesFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		if line == lineNum {
			in := scanner.Text()
			parts := strings.Split(in, ",")
			for i := range parts {
				x, err := strconv.ParseInt(parts[i], 10, 0)
				values = append(values, int(x))
				if err != nil {
					fmt.Println("ERROR:", err)
				}
			}
			break
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return values
}

//Rules implementation
func (gb *Game) Iterate() {

	rules := ReadLine(paramLineNum)

	gbOld := NewGame(gb.xSize, gb.ySize)
	copy(gbOld.cells, gb.cells)

	for y := 0; y < gb.ySize; y++ { //Rows
		for x := 0; x < gb.xSize; x++ { //Collumns

			//Any dead cell with exactly rules[0] live neighbors becomes a live cell
			if !gbOld.Get(x, y) && gbOld.Neighbours(x, y) == rules[0] {
				gb.Set(x, y, true)
				continue
			}

			//Any live cell with fewer than rules[1] live neighbors dies
			if gbOld.Get(x, y) && gbOld.Neighbours(x, y) < rules[1] {
				gb.Set(x, y, false)
				continue
			}

			//Any live cell with rules[2] or rules[3] live neighbors lives
			if gbOld.Get(x, y) && ((gbOld.Neighbours(x, y) == rules[2]) || (gbOld.Neighbours(x, y) == rules[3])) {
				//No need to set, already alive
				continue
			}

			//Any live cell with more than rules[4] live neighbors dies
			if gbOld.Get(x, y) && (gbOld.Neighbours(x, y) > rules[4]) {
				gb.Set(x, y, false)
				continue
			}
		}
	}
}

//Set sets a cell defined by it's x and y coordinates to a given state (alive: true, dead: false)
func (gb *Game) Set(x, y int, val bool) {
	if !gb.InBounds(x, y) {
		log.Fatal("Invalid Coordinate")
	}

	gb.cells[y*(gb.xSize)+x] = val
}

//return cell's state by it's x and y coordinates (alive: true, dead: false)
func (gb *Game) Get(x, y int) bool {
	if !gb.InBounds(x, y) {
		log.Fatal("Invalid Coordinate")
	}

	return gb.cells[y*(gb.xSize)+x]
}

//returns the number of alive neighbours of a given cell
func (gb *Game) Neighbours(x, y int) int {
	count := 0
	arr := []int{-1, 0, 1}

	for _, v1 := range arr {
		for _, v2 := range arr {
			if gb.InBounds(x+v1, y+v2) {
				if gb.Get(x+v1, y+v2) && !(v1 == 0 && v2 == 0) {
					count++
				}
			}
		}
	}
	return count
}

//  Check if the coordinates inside
func (gb *Game) InBounds(x int, y int) bool {
	return (x >= 0 &&
		x < gb.xSize &&
		y >= 0 &&
		y < gb.ySize)
}

func (gb *Game) Print() {
	//Top margin
	fmt.Print("_")
	for x := 1; x <= gb.xSize; x++ {
		fmt.Print("__")
	}
	fmt.Println("_")

	//Rows
	for y := 0; y < gb.ySize; y++ {
		fmt.Print("|")
		//Collumns
		for x := 0; x < gb.xSize; x++ {
			if gb.Get(x, y) {
				fmt.Print("1")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("|")
	}

	//Bottom margin
	fmt.Print("_")
	for x := 1; x <= gb.xSize; x++ {
		fmt.Print("__")
	}
	fmt.Println("_")
}

func main() {
	g := NewGame(xAxis, yAxis)
	g.RandInit(initLiveCells)
	sleepTime := time.Duration(1000*sleepTime) * time.Millisecond

	i := iterationCount // Initial number of itterations
	for {
		if i == 0 {
			break
		}
		i--
		g.Print()
		fmt.Printf("Generation: %d\n", -i)
		g.Iterate()
		time.Sleep(sleepTime)
	}
}

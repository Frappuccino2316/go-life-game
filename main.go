package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	height = 22
	width  = 62
	clear  = "\033[2J"
	head   = "\033[1;1H"
)

type field [height][width]int

var box, tmp field

func main() {
	initialize()

	for {
		fmt.Print(head)
		render()
		update()
		time.Sleep(time.Millisecond * 500)
	}
}

func initialize() {
	fmt.Print(clear)
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			box[i][j] = rand.Intn(2)
		}
	}
}

func render() {
	for _, row := range box {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("■ ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
}

func update() {
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			cnt := 0
			dead := 0
			alive := 1

			cnt = box[i-1][j-1] + box[i-1][j] + box[i-1][j+1] + box[i][j-1] + box[i][j+1] + box[i+1][j-1] + box[i+1][j] + box[i+1][j+1]
			if (box[i][j] == 1 && cnt == 2) || cnt == 3 {
				tmp[i][j] = alive
			} else {
				tmp[i][j] = dead
			}
		}
	}
	box = tmp
}

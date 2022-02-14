package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 10
	height = 10
)

type field [width][height]int

var box field

func main() {
	initialize()

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

func initialize() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			box[i][j] = rand.Intn(2)
		}
	}
}

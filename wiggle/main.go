package main

import (
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

func main() {
	x, y := robotgo.GetScreenSize()

	for true {
		robotgo.MoveMouseSmooth(rand.Intn(x), rand.Intn(y))
		time.Sleep(time.Second * 2)
	}
}

package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	x, y := robotgo.GetScreenSize()

	for true {
		robotgo.MoveMouseSmooth(rand.Intn(x), rand.Intn(y))
		time.Sleep(time.Second * 2)
	}
}

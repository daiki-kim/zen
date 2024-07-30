package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func getRondomPointOnCircle(radius float64) (float64, float64) {
	theta := 2 * math.Pi * rand.Float64()   // ランダムな角度を取得
	r := radius * math.Sqrt(rand.Float64()) // ランダムな半径を取得
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)
	return x, y
}

func main() {
	// 半径5の円上のランダムな座標を10個取得
	for i := 0; i < 10; i++ {
		x, y := getRondomPointOnCircle(5.0)
		fmt.Printf("(%f, %f)\n", x, y)
	}
}

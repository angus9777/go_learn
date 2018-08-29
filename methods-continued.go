package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	f2 := MyFloat(-2)
	fmt.Println(f2.Abs())
	f3 := MyFloat(3)
	fmt.Println(f3.Abs())
}


package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (es ErrNegativeSqrt) Error() string {
	fmt.Println("enter Error!!!")
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(es))
}

func selfSqrt(x ErrNegativeSqrt) (ErrNegativeSqrt, error) {
	fmt.Println("Sqrt ----- 0")
	if x < 0.0 {
		fmt.Println("Sqrt ----- 1")
		return x, x
	} else {
		fmt.Println("Sqrt ----- 2")
	}
	return x, nil
}

func main() {
	fmt.Println("main --------1")
	if re, err := selfSqrt(2.0); err != nil {
		fmt.Println("main --------2.1")
		fmt.Println(err)
	} else {
		fmt.Println("main --------2.2")
		fmt.Println(float64(re))
	}

	if re, err := selfSqrt(-2.0); err != nil {
		fmt.Println("main --------3.1")
		fmt.Println(err)
	} else {
		fmt.Println("main --------3.2")
		fmt.Println(float64(re))
	}
	fmt.Println("main --------4")
	fmt.Println(ErrNegativeSqrt(-2).Error())
	fmt.Println("main --------5")
}

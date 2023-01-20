package main

import (
	"fmt"
	maths_shots "github.com/rahul0tripathi/go-shots/maths"
)

func main() {
	fmt.Println(maths_shots.EuclideanGcd(18, 12))
	fmt.Println(maths_shots.ExtendedEuclidean(1759, 550, 1, 0, 0, 1))
}

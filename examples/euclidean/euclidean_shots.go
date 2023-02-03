package main

import (
	"fmt"
	maths_shots "github.com/rahul0tripathi/go-shots/maths/euclidean"
)

func main() {
	fmt.Println(maths_shots.EuclideanGcd(18, 12))
	fmt.Println(maths_shots.ExtendedEuclidean(1759, 550))
	fmt.Println(maths_shots.ExtendedEuclidean(7, 3))
}

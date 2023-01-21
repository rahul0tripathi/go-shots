package main

import (
	"fmt"
	maths_shots "github.com/rahul0tripathi/go-shots/maths"
)

func main() {
	field := maths_shots.NewField(37)
	a := field.Item(3)
	b := field.Item(18)
	fmt.Println(maths_shots.Add(a, b).Value)
	fmt.Println(maths_shots.Mul(a, b).Value)
	fmt.Println(maths_shots.Sub(a, b).Value)
	fmt.Println(maths_shots.Div(a, b).Value)
}

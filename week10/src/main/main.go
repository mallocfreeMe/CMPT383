package main

import (
	"exer8"
	"fmt"
)

func main() {
	fmt.Printf("%d\n", exer8.Hailstone(18))

	fmt.Println(exer8.HailstoneSequenceAppend(5))

	fmt.Println(exer8.HailstoneSequenceAllocate(5))

	pt := exer8.NewPoint(3, 4.5)
	fmt.Println(pt)                        // should print (3, 4.5)
	fmt.Println(pt.String() == "(3, 4.5)") // should print true

	pt2 := exer8.NewPoint(3, 4)
	fmt.Println(pt2.Norm() == 5.0)
}

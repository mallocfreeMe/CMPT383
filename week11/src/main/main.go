package main

import (
	"exer9"
	"fmt"
	// "math"
)

func main() {
	// p := exer9.NewPoint(1, 2)
	// p.Scale(5)
	// fmt.Println(p)

	// p := exer9.NewPoint(1, 0)
	// p.Rotate(math.Pi / 2)
	// fmt.Println(p)
	// p.Rotate(math.Pi / 2)
	// fmt.Println(p)

	// arr := exer9.RandomArray(10, 10)
	// fmt.Println(arr)

	arr := []int{1,2,3,4,5,6,7,8,9,10}
	mean, stddev := exer9.MeanStddev(arr, 5)

	fmt.Println(mean)
	fmt.Println(stddev)

	arr2 := []int{11,12,13,14,15,16,17,18,19,20}
	mean2, stddev2 := exer9.MeanStddev(arr2, 5)

	fmt.Println(mean2)
	fmt.Println(stddev2)

	// arr2 := []float64{12, 11, 13, 5, 6 }
	// fmt.Println(arr2)
	// exer9.InsertionSort(arr2)
	// fmt.Println(arr2)

	// arr3 := []float64{10,9,8,7,6,5,4,3,2,1}
	// fmt.Println(arr3)
	// exer9.QuickSort(arr3)
	// fmt.Println(arr3)
}

package exer9

import (
	"math"
	"math/rand"
	"time"
)

type Sum struct {
	a float64
	b float64
}

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.
	arr := make([]int, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxInt)
	}

	return arr
}

func CalculateSum(arr []int, result chan Sum) {
	var sumA, sumB float64
	for i := 0; i < len(arr); i++ {
		sumA += float64(arr[i])
		sumB += float64(arr[i] * arr[i])
	}
	s := Sum{sumA, sumB}
	result <- s
}

func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.
	result := make(chan Sum)
	chunkSize := len(arr) / chunks

	for i := 0; i < len(arr); i += chunkSize {
		chunkArr := arr[i : i+chunkSize]
		go CalculateSum(chunkArr, result)
	}

	var sumA, sumB float64

	for i := 0; i < chunks; i++ {
		res := <-result
		sumA += res.a
		sumB += res.b
	}

	mean = sumA / float64(len(arr))
	stddev = math.Sqrt(sumB/float64(len(arr)) - (mean * mean))

	return
}

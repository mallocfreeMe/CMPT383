package exer9

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRandomArrays(t *testing.T) {
	length := 1000
	maxint := 100
	arr1 := RandomArray(length, maxint)
	assert.Equal(t, length, len(arr1))
	for _, v := range arr1 {
		assert.True(t, 0 <= v, "contains a negative integer")
		assert.True(t, v < maxint, "contains an integer >=maxint")
	}

	// check that different calls return different results
	arr2 := RandomArray(length, maxint)
	assert.False(t, reflect.DeepEqual(arr1, arr2))
}

func TestArrayStatParallel(t *testing.T) {
	length := 30000000
	maxint := 10000
	arr2 := RandomArray(length, maxint)

	// call MeanStddev single-threaded
	start := time.Now()
	mean2, stddev2 := MeanStddev(arr2, 1)
	end := time.Now()
	dur1 := end.Sub(start)

	// now turn on cuncurrency and make sure we get the same results, but faster
	start = time.Now()
	mean3, stddev3 := MeanStddev(arr2, 3)
	end = time.Now()
	dur2 := end.Sub(start)

	speedup := float64(dur1) / float64(dur2)
	assert.True(t, speedup > 1.25, "Running MeanStddev with concurrency didn't speed up as expected. Sped up by %g.", speedup)
	assert.Equal(t, float32(mean2), float32(mean3)) // compare as float32 to avoid rounding differences
	assert.Equal(t, float32(stddev2), float32(stddev3))
}

// copied from https://golang.org/src/math/rand/rand.go?s=7456:7506#L225 for Go <1.10 compatibility
func shuffle(n int, swap func(i, j int)) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if n < 0 {
		panic("invalid argument to shuffle")
	}

	i := n - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		swap(i, j)
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

func benchmarkSorting(b *testing.B, sort func(arr []float64)) {
	const length = 1000
	arr := make([]float64, length)
	for i := 0; i < length; i++ {
		arr[i] = float64(i)
	}

	// run the benchmark
	for n := 0; n < b.N; n++ {
		shuffle(length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
		sort(arr)
	}
}

func BenchmarkInsertionSort(b *testing.B) { benchmarkSorting(b, InsertionSort) }
func BenchmarkQuickSort(b *testing.B)     { benchmarkSorting(b, QuickSort) }
func BenchmarkFloat64s(b *testing.B)      { benchmarkSorting(b, sort.Float64s) }

func TestPointScale(t *testing.T) {
	points := make([]Point, 0)
	points = append(points, NewPoint(1, 2))
	points = append(points, NewPoint(3, 4))
	points = append(points, NewPoint(5, 6))

	results := make([]Point, 0)
	results = append(results, NewPoint(5, 10))
	results = append(results, NewPoint(15, 20))
	results = append(results, NewPoint(25, 30))

	for i := 0; i < len(points); i++ {
		points[i].Scale(5)
		assert.Equal(t, points[i], results[i], "The two struct should be equal.")
	}
}

func TestPointRotate(t *testing.T) {
	points := make([]Point, 0)
	points = append(points, NewPoint(1, 0))
	points = append(points, NewPoint(0, 1))

	results := make([]Point, 0)
	results = append(results, NewPoint(0, 1))
	results = append(results, NewPoint(-1, 0))

	for i := 0; i < len(points); i++ {
		points[i].Rotate(math.Pi / 2)
		assert.Equal(t, math.Round(points[i].x), results[i].x, "The two struct x should be equal.")
		assert.Equal(t, math.Round(points[i].y), results[i].y, "The two struct y should be equal.")
	}
}

func TestMeanStddev(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mean, stddev := MeanStddev(arr, 5)
	assert.Equal(t, mean, 5.5)
	assert.Equal(t, math.Round(stddev), float64(3))

	arr2 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	mean2, stddev2 := MeanStddev(arr2, 5)
	assert.Equal(t, mean2, 15.5)
	assert.Equal(t, math.Round(stddev2), float64(3))
}

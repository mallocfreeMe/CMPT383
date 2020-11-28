package exer10

func Helper(n uint, ch chan uint) {
	var result uint
	ch1 := make(chan uint)
	ch2 := make(chan uint)

	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		go Helper(n-1, ch1)
		go Helper(n-2, ch2)
		r1 := <-ch1
		r2 := <-ch2
		result = r1 + r2
	}

	ch <- result
}

func Fib(n uint, cutoff uint) uint {
	var result uint

	ch := make(chan uint)

	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		if n > cutoff {
			go Helper(n, ch)
			result = <-ch
		} else {
			result = Fib(n-1, cutoff) + Fib(n-2, cutoff)
		}
	}

	return result
}

func Fibonacci(n uint) uint {
	return Fib(n, 37)
}

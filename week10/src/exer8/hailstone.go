package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions
/*
	HailstoneSequenceAppend releative spped is 2579 ns/op
	HailstoneSequenceAllocate releative speed is 1164 ns/op
*/
func Hailstone(n uint) uint {
	var result uint

	if n%2 == 0 {
		result = n / 2
	} else {
		result = 3*n + 1
	}

	return result
}

func helper(n uint, result *[]uint) []uint {
	if n > 1 {
		*result = append(*result, n)
		if Hailstone(n) > 1 {
			helper(Hailstone(n), result)
		} else if Hailstone(n) == 1 {
			*result = append(*result, 1)
		}
	}

	return *result
}

func HailstoneSequenceAppend(n uint) []uint {
	result := []uint{}
	helper(n, &result)
	return result
}

func helper2(n uint, size *uint) uint {
	if n > 1 {
		*size++
		if Hailstone(n) > 1 {
			helper2(Hailstone(n), size)
		} else if Hailstone(n) == 1 {
			*size++
		}
	}

	return *size
}

func HailstoneSequenceAllocate(n uint) []uint {
	var size uint
	helper2(n, &size)
	arr := make([]uint, size)
	for i := 0; i < int(size); i++ {
		arr[i] = n
		n = Hailstone(n)
	}
	return arr
}

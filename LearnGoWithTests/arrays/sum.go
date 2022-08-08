package main

func Sum(numbers []int) (sum int) {
	// index and value
	// Using blank identifier to ignore index here
	for _, number := range numbers {
		sum += number
	}
	return
}

func Sum5(numbers [5]int) (sum int) {
	// index and value
	// Using blank identifier to ignore index here
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(sliceOfNumbers ...[]int) (sums []int) {
	sums = make([]int, len(sliceOfNumbers))
	for iter, numbers := range sliceOfNumbers {
		sums[iter] = Sum(numbers)
	}
	return sums
}

func main() {}

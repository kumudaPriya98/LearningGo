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

func SumAll(sliceOfNumbers ...[]int) []int {
	var sums []int
	for _, numbers := range sliceOfNumbers {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func main() {}

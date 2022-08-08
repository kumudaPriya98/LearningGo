package main

func Sum(numbers [5]int) (sum int) {
	// index and value
	// Using blank identifier to ignore index here
	for _, number := range numbers {
		sum += number
	}
	return
}

func main() {}
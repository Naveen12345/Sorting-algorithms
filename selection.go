package main

import "fmt"

func main() {
	var sample [8]int
	fmt.Println("Enter", 8, "elements:")
	for i := 0; i < 8; i++ {
		fmt.Scanln(&sample[i])
	}
	len := len(sample)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if sample[j] < sample[minIndex] {
				sample[j], sample[minIndex] = sample[minIndex], sample[j]
			}
		}
	}
	fmt.Println("\nAfter SelectionSort")
	for _, val := range sample {
		fmt.Println(val)
	}
}

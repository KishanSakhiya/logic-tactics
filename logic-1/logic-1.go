package main

import "fmt"

func main() {
	input := []int{3, 4, 500, 501, 6, 7, 400, 8, 9}
	intermediateValues := make([]int, 0)
	intermediateValues = append(intermediateValues, input[0])

	intermediateIndexCount := 0
	i, j := 0, 1
	for j < len(input) {
		if input[i] <= input[j] {
			intermediateValues[intermediateIndexCount] += input[j]
		} else {
			intermediateIndexCount++
			intermediateValues = append(intermediateValues, input[j])
		}
		i++
		j++
	}

	result := intermediateValues[0]
	for _, v := range intermediateValues {
		if result < v {
			result = v
		}
	}
	fmt.Println("Output : ", result)
}

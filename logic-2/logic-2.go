package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	ok := reader.Scan()
	if !ok {
		fmt.Println("scan is not working")
	}
	num := reader.Text()

	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Please enter proper number")
	}

	result := 0
	if n > 2 {
		if n%2 == 0 {
			result = n - 3
		} else {
			result = n - 1
		}
	}

	fmt.Println("Output : ", result)

}

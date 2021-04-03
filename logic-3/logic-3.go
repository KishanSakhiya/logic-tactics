package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number of rows : ")
	ok := reader.Scan()
	if !ok {
		fmt.Println("scanner is not working")
	}
	str := reader.Text()

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Please enter proper number")
	}

	n := num * (num + 1) / 2

	count := 0
	for i, j := 0, 1; i < n; i++ {
		fmt.Print("* ")
		count++
		if j == count {
			fmt.Println()
			count = 0
			j++
		}
	}
}

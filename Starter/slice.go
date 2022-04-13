package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var scanValue string
	intSlice := make([]int, 0, 3)

	for {
		fmt.Println("enter an integer: ")
		if _, err := fmt.Scan(&scanValue); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if scanValue == "X" {
			break
		}
		n, err := strconv.Atoi(scanValue)
		if err != nil {
			fmt.Println(err)
			continue
		}
		intSlice = append(intSlice, n)
		sort.Ints(intSlice)
		fmt.Println(intSlice)
	}
}

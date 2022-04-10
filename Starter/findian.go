package main

import (
	"bufio"
	"fmt"
	"os"
)
import "strings"


func main() {

	fmt.Printf("Enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	fmt.Printf("\tstring entered is: \t\"%v\" \n", inputString)
    lowerString := strings.ToLower(inputString)
    formattedString := strings.ReplaceAll(lowerString, " ", "")
	
	if strings.Contains(formattedString, "a") { 
		if strings.HasPrefix(formattedString, "i") { 
			if strings.HasSuffix(formattedString, "n") { 
				fmt.Println("\tContains (‘i’, ‘a’, and ‘n’), Found!")
				os.Exit(0)
			}
		}
	}
	fmt.Println("\tDoesn't contain (‘i’, ‘a’, and ‘n’), Not Found!")
}
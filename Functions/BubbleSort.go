package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
   "strings"
)



func BubbleSort(array[] int)[]int {
   for i:=0; i< len(array); i++ {
      for j:=0; j < len(array)-1-i; j++ {
         if array[array + j] > array[j]) {
          swap(array, j)
         }
      }
   }
}


func swap(array []int, j int) {
   var temp int
   temp = array[j]
   array[j] = array[j+1]
   array[j+1] = temp
}

func main() {
   fmt.Println("Input numbers separated by space:=> ")
   scanner := bufio.NewReader(os.Stdin)
   input, _, _ := scanner.ReadLine()
   newString := strings.Split(string(input), " ")
   var values []int
   for _, s := range newString {
      n, _ := strconv.Atoi(s)
      values = append(values, n)
   }
   BubbleSort(values)
   fmt.Println(values)
   
}
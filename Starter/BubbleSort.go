package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
   "strings"
)



func BubbleSort(array[] int)[]int {
   for i:=0; i< len(array)-1; i++ {
      for j:=0; j < len(array)-i-1; j++ {
         if (array[j] > array[j+1]) {
          swap(array, j)
         }
      }
   }
   return array
}


func swap(a []int, j int) {
   var temp int
   temp = a[j]
   a[j] = a[j+1]
   a[j+1] = temp
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
package main

import (
   "fmt"
   "bufio"
   "os"
   "strconv"
   "strings"
)


func main() {
   fmt.Println("Input numbers separated by space:=> ")
   scanner := bufio.NewReader(os.Stdin)
   input, _, _ := scanner.ReadLine()
   newString := strings.Split(string(input), " ")
   var values []int
   for _, s := range(newString) {
      n, _ := strconv.Atoi(s)
      values = append(values, n)
   }
   BubbleSort(values)
   fmt.Println(values)
   
}



func BubbleSort(arr []int) {
   for i := 0; i < len(arr); i++ {
      for j := 0; j < len(arr)-1-i; j++ {
         if arr[j + 1] < arr[j] {
          swap(arr, j)
         }
      }
   }
}


func swap(a []int, j int) {
   var temp int
   temp = a[j]
   a[j] = a[j+1]
   a[j+1] = temp
}

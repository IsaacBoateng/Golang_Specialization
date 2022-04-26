package main

import (
    "fmt"
    "encoding/json"
    "os"
    "bufio"
)

func main() {
    var m map[string]string
    m = make(map[string]string)
    fmt.Println("Input your name:")
    br := bufio.NewReader(os.Stdin)
    bname, _, _ := br.ReadLine()
    name:=string(bname)
    m["name"]=name
    fmt.Println("Input your adress:")
    br1 := bufio.NewReader(os.Stdin)
    badress, _, _ := br1.ReadLine()
    adress:=string(badress)
    m["adress"]=adress
    b, err := json.MarshalIndent(m, "", " ")
    if err != nil {
        fmt.Println("Encoding faild")
    } else {
        fmt.Println("Encoded data : ")
        fmt.Println(b)
        fmt.Printf("JSON Data : %v", string(b))
    }
}
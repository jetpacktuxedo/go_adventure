package main

import (
  "bufio"
  "fmt"
  "os"
)


func main() {
  fmt.Printf("You are standing in an open cube inside of a white office full of sound-dampening walls. There are new emails in your inbox.\n")
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Printf("> ")
    input, err := reader.ReadString('\n')
    if(err != nil){
      fmt.Println(err)
    }
    fmt.Println(input)
  }
}

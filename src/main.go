package main

import (
     "fmt"
)

func main() {
     a := 5
     b := &a
     a = 7
     fmt.Println(a)
     fmt.Println(*b)
}

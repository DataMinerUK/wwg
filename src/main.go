package main

import (
     "fmt"
)

type Kitten struct {
     Name string
}

func (k *Kitten) SetName(name string) {
     k.Name = name
}

func (k *Kitten) GetName() string {
     return k.Name
}

func main() {
     kitty := Kitten{}
     kitty.SetName("Mr Tiggles")
     fmt.Println(kitty.GetName())
}

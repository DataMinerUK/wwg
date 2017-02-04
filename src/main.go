package main

import (
     "fmt"
     "github.com/DataMinerUK/wwg/src/animals"
)

func main() {
     kitty := animals.Kitten{}
     kitty.SetName("Mr Tiggles")
     fmt.Println(kitty.GetName())

     kitty2 := &kitty
     kitty2.SetName("Mr Tom")
     fmt.Println(kitty.GetName())
     fmt.Println(kitty2.GetName())

     deffer := animals.Dog{}
     deffer.SetName("Defer")
     fmt.Println(deffer.GetName(), "goes", deffer.MakeNoise())
}

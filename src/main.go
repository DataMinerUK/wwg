package main

import (
     "net/http"
     "github.com/DataMinerUK/wwg/src/animals"
     "encoding/json"
)

var kittens []animals.Kitten

func main() {
     kittens = []animals.Kitten{
       animals.Kitten{
         Name: "Mr Tiggles",
         Hobbies: []string{
           "Playing with wool",
           "Eating",
         },
       },
     }
     http.HandleFunc("/list", ListKittens)
     http.ListenAndServe(":9000", http.DefaultServeMux)
}

func ListKittens(rw http.ResponseWriter, r *http.Request) {
    data, err := json.Marshal(kittens)
    if err != nil {
      rw.WriteHeader(http.StatusInternalServerError)
      return
    }
    rw.Write(data)
}

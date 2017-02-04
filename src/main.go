package main

import (
     "net/http"
     "github.com/DataMinerUK/wwg/src/animals"
     "encoding/json"
)

var pets []animals.Pet

func main() {
     pets = []animals.Pet{
       &animals.Kitten{
         Name: "Mr Tiggles",
         Hobbies: []string{
           "Playing with wool",
           "Eating",
         },
       },
       &animals.Kitten{
         Name: "Ms Tiddles",
         Hobbies: []string{
           "Scratching",
           "Mouse hunt",
         },
       },
       &animals.Dog{
         Name: "Deffer",
         Hobbies: []string{
           "Wag tail",
           "Sniff bum",
         },
       },
     }
     http.HandleFunc("/list", ListPets)
     http.ListenAndServe(":9000", http.DefaultServeMux)
}

func ListPets(rw http.ResponseWriter, r *http.Request) {
    data, err := json.Marshal(pets)
    if err != nil {
      rw.WriteHeader(http.StatusInternalServerError)
      return
    }
    rw.Write(data)
}

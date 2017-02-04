package main

import (
     "fmt"
     "net/http"
)

func main() {
       http.HandleFunc("/hello", HelloWorld)
       http.ListenAndServe(":9000", http.DefaultServeMux)
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
      fmt.Fprint(rw, "Hello World!")
}

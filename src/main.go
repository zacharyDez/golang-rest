package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World!")
    fmt.Println("Endpoint Hit: Hello World")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {
    handleRequests()
}
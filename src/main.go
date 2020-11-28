package main

import (
	"fmt"
	"log"
	"net/http"
)

var name string

func getGreeting() string {
	if name == "" {
		return "Hello World!"
	}

	return "Hello " + name + "!"
}

func homePage(w http.ResponseWriter, r *http.Request) {
	names, ok := r.URL.Query()["name"]

	// helloworld if name is not present
	if !ok || len(names[0]) < 1 {
		log.Println("Url Param 'name' is missing. Showing hello world!")
		name = ""
	} else {
		name = names[0]
		log.Println("Url Param 'name' is: " + string(name))
	}

	fmt.Fprintf(w, getGreeting())
	fmt.Println("Endpoint Hit: Homepage")
}

func handlerHomePage() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {
	handlerHomePage()
}

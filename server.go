// a simple RESTful application for creating todo list
package main

import "net/http"

type Todo struct {
	Id        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var TodoIndex = make(map[string]*Todo)

func main() {
	router := NewRouter()
	http.ListenAndServe(":3000", router)
}

/*



*/

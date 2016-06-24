package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Todo struct {
	Id        string   `json:"id"`
	Task      []string `json:"task"`
	Completed bool     `json:"completed"`
}

type Todos []Todo

var todoId = 1

var TodoList Todos

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadFile("public/todo.html")
	if err != nil {
		fmt.Fprintf(w, "invalid :", err)
		return
	}
	w.Header().Add("Content Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func saveTodos(w http.ResponseWriter, req *http.Request) {
	fp, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error saving file", 501)
		return
	}
	json.NewEncoder(fp).Encode(TodoList)
	w.Write([]byte("saved successfully"))
	http.RedirectHandler("http://localhost:8080/showTodos", http.StatusMovedPermanently)
}

func showTodos(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	for k, v := range req.Form {
		var temp = Todo{}
		temp.Id = fmt.Sprintf("%d", todoId)
		todoId++
		temp.Task = v
		if k == "completed" {
			fmt.Println("button ticked :", v)
			temp.Completed = true
		}
		fmt.Println(k, v)
		TodoList = append(TodoList, temp)
	}

	w.Header().Add("Content Type", "application/json")
	data, _ := json.MarshalIndent(TodoList, "", "    ")

	w.Write(data)
	fp, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error saving file", 501)
		return
	}

	json.NewEncoder(fp).Encode(TodoList)
	w.Write([]byte("saved successfully"))
}

func deleteTodo(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var newList []Todo
	delid := req.URL.Path[len("/deleteTodo/"):] // extract id coming after Todo/
	for _, v := range TodoList {
		if v.Id != delid {
			newList = append(newList, v)
		}
	}
	TodoList = newList
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/new", CreateTodo)
	mux.HandleFunc("/showTodos", showTodos)
	mux.HandleFunc("/deleteTodo/", deleteTodo)
	http.ListenAndServe(":3000", mux)
}

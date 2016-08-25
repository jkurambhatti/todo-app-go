// a simple RESTful application for creating todo list
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todoId = 1

type TodoList []Todo

func main() {
	mux := mux.NewRouter()
	//mux.HandleFunc("/new", CreateTodo)
	//mux.HandleFunc("/showTodo", ShowTodo)
	//mux.HandleFunc("/deleteTodo/", DeleteTodo)
	mux.HandleFunc("/", Index)
	//mux.HandleFunc("/loadTodo", LoadTodo)
	//http.HandleFunc("/", Index)
	fmt.Println("listening at :3000")
	http.ListenAndServe(":3000", mux)
}

func Index(w http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile("public/index.html")
		if err != nil {
			fmt.Errorf("error reading index file : %s", err)
		}
		w.Header().Add("Content Type", " text/html")
		w.WriteHeader(200)
		w.Write(data)
}

/*

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

func SaveTodos(w http.ResponseWriter, req *http.Request) {
	fp, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		http.Error(w, "error saving file", 501)
		return
	}
	defer fp.Close()
	json.NewEncoder(fp).Encode(TodoList)
	w.Write([]byte("saved successfully"))
	// http.RedirectHandler("http://localhost:8080/showTodos", http.StatusMovedPermanently)
}

func LoadTodo(w http.ResponseWriter, req *http.Request) {
	var newload Todos
	f, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		http.Error(w, "error loading todos :", http.StatusBadRequest)
		return
	}
	json.NewDecoder(f).Decode(&newload)
	fmt.Println(newload)
	if newload == nil {
		w.Write([]byte("no records found"))
	} else {
		w.Write([]byte("records loaded successfully"))
		json.NewEncoder(w).Encode(newload)
	}

	TodoList = newload
}

func ShowTodo(w http.ResponseWriter, req *http.Request) {
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
	defer fp.Close()
	json.NewEncoder(fp).Encode(TodoList)
	w.Write([]byte("saved successfully"))
}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
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

*/

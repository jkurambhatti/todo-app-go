// a simple RESTful application for creating todo list
package main

import "net/http"

type Todo struct {
	Id        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todoId = 1

type TodoList []*Todo

var TodoIndex = make(map[string]*Todo)

func main() {

	//mux := mux.NewRouter().StrictSlash(true)
	//mux.HandleFunc("/insert", CreateTodo)
	////mux.HandleFunc("/list", ShowTodo)
	////mux.HandleFunc("/delete/", DeleteTodo)
	////mux.HandleFunc("/loadTodo", LoadTodo)
	//mux.HandleFunc("/", Index)
	//fmt.Println("listening at :3000")
	router := NewRouter()
	http.ListenAndServe(":3000", router)
}

/*
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


*/

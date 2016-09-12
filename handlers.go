package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"os"
)

func Index(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("public/index.html")
	//data, err := ioutil.ReadFile("public/todo.html")

	if err != nil {
		fmt.Errorf("error reading index file : %s", err)
	}
	w.Header().Add("Content Type", " text/html")
	w.WriteHeader(200)
	w.Write(data)
}

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	var t = new(Todo)
	fmt.Println("reached insert path")
	json.NewDecoder(req.Body).Decode(t)
	v, err := json.Marshal(*t)
	if err != nil {
		fmt.Println("error marshaling json")
	}
	TodoIndex[t.Id] = t
	fmt.Println("added new task to the taskindex : ", string(v))

	w.Header().Add("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(v)
}

func GetTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("reached getTodo path")
	vars := mux.Vars(req)
	tid := ""
	todo := new(Todo)
	if v, ok := vars["id"]; ok {
		tid = v
	}
	todo = TodoIndex[tid]
	fmt.Println(tid, todo)

	// tdtsk, err := json.Marshal(todo)
	// if err != nil {
	// 	fmt.Println("error marshaling todo")
	// }
	w.Header().Add("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
	// w.Write(tdtsk)

}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("reached delete path")
	vars := mux.Vars(req)
	did := ""
	if v, ok := vars["id"]; ok {
		did = v
		delete(TodoIndex, v)
	}
	msg := fmt.Sprintf("task %s has been deleted", did)
	w.Write([]byte(msg))
}

func ShowTodo(w http.ResponseWriter, req *http.Request) {
	tdl, err := json.MarshalIndent(TodoIndex, " ", "    ")
	if err != nil {
		fmt.Println("error marshaling todolistindex")
	}
	w.Header().Add("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tdl)
}

func UpdateTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("reached Update path")
	var t = new(Todo)
	json.NewDecoder(req.Body).Decode(t)
	v, err := json.MarshalIndent(*t, " ", "    ")
	if err != nil {
		fmt.Println("error marshaling json")
	}
	vars := mux.Vars(req)
	updateid := ""
	if id, ok := vars["id"]; ok {
		updateid = id
		TodoIndex[id] = t
	}
	msg := fmt.Sprintf("task %s has been updated : %s \n", updateid, string(v))
	w.Write([]byte(msg))
}

func SaveTodo(w http.ResponseWriter, req *http.Request) {
	fp, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		http.Error(w, "error opening file", 501)
		return
	}
	defer fp.Close()
	json.NewEncoder(fp).Encode(TodoIndex)
	w.Write([]byte("saved successfully"))
}

func LoadTodo(w http.ResponseWriter, req *http.Request) {
	var newload = make(map[string]*Todo)
	f, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		http.Error(w, "error loading todos :", http.StatusNotFound)
		return
	}
	json.NewDecoder(f).Decode(&newload)
	fmt.Println(newload)
	if newload == nil {
		w.Write([]byte("no records found"))
	} else {
		// w.Write([]byte("records loaded successfully"))
		json.NewEncoder(w).Encode(newload)
	}

}

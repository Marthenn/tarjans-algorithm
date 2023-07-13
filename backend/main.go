package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang-collections/collections/set"
)

type Message struct {
	U       string              `json:"u"`
	V       string              `json:"v"`
	AdjList map[string][]string `json:"adjList"`
	Names   []string            `json:"names"`
}

func main() {
	http.HandleFunc("/file", FileHandler)
	http.HandleFunc("/addEdge", AddEdgeHandler)
	port := ":8080"
	http.ListenAndServe(port, nil)
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var message Message
	err = json.Unmarshal(body, &message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if message.U == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	names_set := set.New()
	adjList := FileToAdjList(message.U, names_set)

	message.AdjList = adjList
	message.Names = SetToList(names_set)
	message.U = ""
	message.V = ""

	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func AddEdgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var message Message
	err = json.Unmarshal(body, &message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if message.AdjList == nil || message.Names == nil || message.U == "" || message.V == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	names_set := ListToSet(message.Names)

	message.AdjList = AddEdge(message.AdjList, names_set, message.U, message.V)

	message.Names = SetToList(names_set)

	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

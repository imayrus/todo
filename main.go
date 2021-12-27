package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Entry struct {
	ID   string    `json:"id"`
	Job  string    `json:"job"`
	Time time.Time `json:"time"`
}

var entryList []Entry

func getEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entryList)
}

func getEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range entryList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	l := len(entryList)
	var newEntry Entry
	_ = json.NewDecoder(r.Body).Decode(&newEntry)
	newEntry.ID = strconv.Itoa((l + 1))
	newEntry.Time = time.Now()
	entryList = append(entryList, newEntry)
	fmt.Printf("%+v", newEntry)
	json.NewEncoder(w).Encode(newEntry)
}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i := 0; i < len(entryList); i++ {
		if params["id"] == entryList[i].ID {
			entryList = append(entryList[:i], entryList[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(entryList)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/list", getEntries).Methods("GET")
	r.HandleFunc("/get/{id}", getEntry).Methods("GET")
	r.HandleFunc("/create", createEntry).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteEntry).Methods("DELETE")

	http.ListenAndServe(":8100", r)
}

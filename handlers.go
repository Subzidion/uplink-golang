package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func GenerationIndex(w http.ResponseWriter, r *http.Request) {
    generations := Generations {
        Generation{Name: "Test"},
    }

     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
     w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(generations); err != nil {
        panic(err)
    }
}

func GenerationById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    generationId := vars["generationId"]
    fmt.Fprintln(w, "Generation ID:", generationId)
}

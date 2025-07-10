package main

import (
    "fmt"
    "log"
    "net/http"
    "referee-service/handlers"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/referees", handlers.CreateReferee).Methods("POST")
    r.HandleFunc("/referees", handlers.GetReferees).Methods("GET")

    fmt.Println("Referee service running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

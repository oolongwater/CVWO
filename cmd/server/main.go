// cmd/server/main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/CVWO/sample-go-app/internal/router"
)

type Data struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    r := router.Setup()
    r.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
        data := []Data{
            {ID: 1, Name: "Sample Data 1"},
            {ID: 2, Name: "Sample Data 2"},
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    })

    fmt.Print("Listening on port 8000 at http://localhost:8000!")
    log.Fatalln(http.ListenAndServe(":8000", r))
}
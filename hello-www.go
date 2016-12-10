package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloWWW(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World Wide Web!")
}

func main() {
    http.HandleFunc("/", helloWWW)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

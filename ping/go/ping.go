package main

import (
    "log"
    "net/http"
    "io"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "go pong")
}

func main() {
    http.HandleFunc("/", PingHandler)
    log.Fatal(http.ListenAndServe(":80", nil))
}

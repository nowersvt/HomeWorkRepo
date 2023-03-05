package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "time"
)

var bgColor string
//
func main() {
    hostname, _ := os.Hostname()

    colors := []string{"black", "blue", "orange"}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if bgColor == "" {
            rand.Seed(time.Now().UnixNano())
            bgColor = colors[rand.Intn(len(colors))]
        }
        fmt.Fprintf(w, "<html><body style='background-color: %s; color: white'><h1>Hostname: %s</h1></body></html>", bgColor, hostname)
        // Сохраняем текущий цвет фона
        w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
        w.Header().Set("Pragma", "no-cache")
        w.Header().Set("Expires", "0")
    })

    if err := http.ListenAndServe(":8003", nil); err != nil {
        panic(err)
    }
}

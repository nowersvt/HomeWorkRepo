package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "time"
)

var bgColor string

func main() {
    hostname, _ := os.Hostname()
<<<<<<< HEAD
<<<<<<< HEAD

=======
//comit 1
//comit 2
//comit 3
<<<<<<< HEAD
//comit 4
//comit 5
<<<<<<< HEAD
>>>>>>> efbcede (5)
=======
>>>>>>> 5b25036 (3)
=======
//comit 1
>>>>>>> 3e510d3 (1)
=======
//commit to test
// a
//b
>>>>>>> 97221a1 (//b)
    colors := []string{"black", "blue", "orange"}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if bgColor == "" {
            rand.Seed(time.Now().UnixNano())
            bgColor = colors[rand.Intn(len(colors))]
        }
        fmt.Fprintf(w, "<html><body style='background-color: %s; color: white'><h1>Hostname: %s</h1></body></html>", bgColor, hostname)
        w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
        w.Header().Set("Pragma", "no-cache")//
        w.Header().Set("Expires", "0")//
    })//

    if err := http.ListenAndServe(":8003", nil); err != nil {
        panic(err)
    }
}

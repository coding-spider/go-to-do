package main

import (
    "log"
    "net/http"

    "go-to-do/server/routes"
)

// bootsrap app
func main() {
    log.Fatal(http.ListenAndServe(":8000", routes.GetRouter()))
}

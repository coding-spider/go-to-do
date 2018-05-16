package routes

import (
    // "encoding/json"
    "github.com/gorilla/mux"
    // "net/http"
)

func GetRouter() *mux.Router {
    router := mux.NewRouter()
    return router
}

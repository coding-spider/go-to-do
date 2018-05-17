package routes

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    // "fmt"

    "go-to-do/server/models"
)

// Get All Task List
func GetTaskLists(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.TaskLists)
}

// Creating a task List
func CreateTaskList(w http.ResponseWriter, r *http.Request) {
    var newTaskList models.TaskList
    _ = json.NewDecoder(r.Body).Decode(&newTaskList)
    createdTaskList := models.CreateNewTaskList(&newTaskList)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdTaskList)
}

func GetRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/tasklists", GetTaskLists).Methods("GET")
    router.HandleFunc("/tasklists", CreateTaskList).Methods("POST")
    return router
}
package routes

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
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
    createdTaskList := models.CreateNewTaskList(newTaskList)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdTaskList)
}

// Creating task for a task list
func CreateTaskForATaskList(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var newTask models.Task
    _ = json.NewDecoder(r.Body).Decode(&newTask)
    newTask.TaskListId = params["id"]
    createdTask := models.CreateNewTask(newTask)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdTask)
}

func GetTaskListById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    taskId, _ := strconv.Atoi(params["id"])
    w.Header().Set("Content-Type", "application/json")
    for _, item := range models.TaskLists {
        if item.ID == taskId {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(nil)
}

func GetRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/tasklists", GetTaskLists).Methods("GET")
    router.HandleFunc("/tasklists/{id}", GetTaskListById).Methods("GET")
    router.HandleFunc("/tasklists", CreateTaskList).Methods("POST")
    router.HandleFunc("/tasklists/{id}/createTask", CreateTaskForATaskList).Methods("POST")
    return router
}

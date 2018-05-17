package routes

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "fmt"

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
    newTask.TaskListId, _ = strconv.Atoi(params["id"])
    createdTask := models.CreateNewTask(newTask)
    var modifiedTaskList models.TaskList
    modifiedTaskList = models.GetTaskListDetailsById(createdTask.TaskListId)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(modifiedTaskList)
}

// Fetching Task List by Id
func GetTaskListById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    taskListId, _ := strconv.Atoi(params["id"])
    var foundTaskList models.TaskList
    foundTaskList = models.GetTaskListDetailsById(taskListId)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(foundTaskList)
}

// Delete Task List By Id
func DeleteTaskListById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    taskId, _ := strconv.Atoi(params["id"])

    //Delete TaskList
    for idx, item := range models.TaskLists {
        if item.ID == taskId {
            models.TaskLists = append(models.TaskLists[:idx], models.TaskLists[idx+1:]...)
            break
        }
    }

    //Delete Tasks
    for idx, item := range models.Tasks {
        if item.TaskListId == taskId {
            models.Tasks = append(models.Tasks[:idx], models.Tasks[idx+1:]...)
            break
        }
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(true)
}

func GetRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/tasklists", GetTaskLists).Methods("GET")
    router.HandleFunc("/tasklists/{id}", GetTaskListById).Methods("GET")
    router.HandleFunc("/tasklists/{id}", DeleteTaskListById).Methods("DELETE")
    router.HandleFunc("/tasklists", CreateTaskList).Methods("POST")
    router.HandleFunc("/tasklists/{id}/createTask", CreateTaskForATaskList).Methods("POST")
    fmt.Println("Server is running");
    return router
}

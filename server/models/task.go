package models

import (
  "errors"
)

// Task
type Task struct {
    ID        int   `json:"id,omitempty"`
    Name string   `json:"name,omitempty"`
    Description  string   `json:"description,omitempty"`
    Status string `json:"status,omitempty"`
    TaskListId   int `json:"taskListId,omitempty"`
}

var Tasks []Task

func CreateNewTask (newTask Task) (bool, error){

    taskListFound := false
    //Find TaskList
    for _, item := range TaskLists {
        if item.ID == newTask.TaskListId {
            taskListFound = true
            break
        }
    }

    if(taskListFound == false) {
      return false, errors.New("Task Not Found !")
    }

    newTask.ID = len(Tasks) + 1
    Tasks = append(Tasks, newTask)
    return true, nil
}

func UpdateTaskStatus (taskListId int, taskId int, status string) (bool, error){
    //Update Task Status
    for idx := 0; idx < len(Tasks); idx++ {
        item := &Tasks[idx]
        if item.TaskListId == taskListId && item.ID == taskId {
            item.Status = status
            return true, nil
        }
    }
    return false, errors.New("Invalid Details Passed !")
}

func DeleteTask (taskListId int, taskId int) (bool, error){
    deleted := false
    //Delete TaskList
    for idx, item := range Tasks {
        if item.TaskListId == taskListId && item.ID == taskId {
            Tasks = append(Tasks[:idx], Tasks[idx+1:]...)
            deleted = true
            break
        }
    }

    if(deleted == true) {
      return deleted, nil
    }
    return deleted, errors.New("Task Not Found !")
}

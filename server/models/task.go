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

func CreateNewTask (newTask Task) Task{
  newTask.ID = len(Tasks) + 1
  Tasks = append(Tasks, newTask)
  return newTask;
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

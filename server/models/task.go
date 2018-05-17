package models

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

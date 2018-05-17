package models

// Task
type Task struct {
    ID        int   `json:"id,omitempty"`
    Name string   `json:"name,omitempty"`
    Description  string   `json:"description,omitempty"`
    Status bool `json:"status,omitempty"`
    TaskListId   string `json:"id_task_list,omitempty"`
}

var Tasks []Task

func CreateNewTask (newTask Task) Task{
  newTask.ID = len(Tasks) + 1
  Tasks = append(Tasks, newTask)
  return newTask;
}

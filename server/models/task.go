package models

// Task
type Task struct {
    ID        int   `json:"id,omitempty"`
    Name string   `json:"name,omitempty"`
    Description  string   `json:"description,omitempty"`
    Status bool `json:"status,omitempty"`
    TaskListId   string `json:"id_task_list,omitempty"`
}

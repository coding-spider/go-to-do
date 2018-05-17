package models

// TaskList
type TaskList struct {
    ID        int   `json:"id,omitempty"`
    Name string   `json:"name,omitempty"`
    Description  string   `json:"description,omitempty"`
    Tasks   []Task `json:"tasks"`
}

var TaskLists []TaskList

func CreateNewTaskList (newTaskList TaskList) TaskList{
  newTaskList.ID = len(TaskLists) + 1
  TaskLists = append(TaskLists, newTaskList)
  return newTaskList;
}

func GetTaskListDetailsById (taskListId int) TaskList {
  var existingTaskList TaskList
  var tasks []Task

  for _, item := range TaskLists {
      if item.ID == taskListId {
          existingTaskList = item
          break
      }
  }

  for _, item := range Tasks {
      if item.TaskListId == taskListId {
          tasks = append(tasks, item);
      }
  }

  existingTaskList.Tasks = tasks

  return existingTaskList
}


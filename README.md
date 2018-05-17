# go-to-do
TODO App in Golang

### Install Necessary packages
```
go get github.com/gorilla/mux
```

### Start the server. Default Port is 8000
```
go run server/server.go
```

### Features
 - There are 2 models ``TaskList`` and ``Task`` with the relation as ```TaskList hasMany Tasks```
 - All the data is is stored in application variables to mock persistence


### Sample Requests
> 1. Creating a new TaskList

````
url - POST localhost:8000/tasklists

Request -
{
  "name": "Home TODO List",
  "description": "Home TODO List",
  "tasks": []
}

Respnse -
{
    "id": 1,
    "name": "Home TODO List",
    "description": "Home TODO List",
    "tasks": []
}
````


> 2. Adding a Task to a TaskList

````
url - POST localhost:8000/tasklists/1/createTask

Request -
{
  "name": "Need to Complete TOTO Golang App",
  "description": "Need to Complete TOTO Golang App",
  "status": "Created"
}

Response -
{
    "id": 1,
    "name": "Home TODO List",
    "description": "Home TODO List",
    "tasks": [
        {
            "id": 2,
            "name": "Need to Complete TOTO Golang App",
            "description": "Need to Complete TOTO Golang App",
            "status": "Created",
            "taskListId": 1
        }
    ]
}
````

> 3. Deleting a Task from a TaskList

````
url - DELETE localhost:8000/tasklists/1/deleteTask/2

Response -
{
    "id": 1,
    "name": "Home TODO List",
    "description": "Home TODO List",
    "tasks": null
}
````

> 4. Updating a Task in a TaskList. Marking Task Status as completed

````
url - PUT localhost:8000/tasklists/1/updateTaskStatus/1/status/Completed

Response -
{
    "id": 1,
    "name": "Home TODO List",
    "description": "Home TODO List",
    "tasks": [
        {
            "id": 1,
            "name": "Need to Complete TOTO Golang App",
            "description": "Need to Complete TOTO Golang App",
            "status": "Completed",
            "taskListId": 1
        }
    ]
}
````

> 5. Deleting TaskList

````
url - DELETE localhost:8000/tasklists/1

Response -

true
````

> 6. Fetching TaskList Details

````
url - GET localhost:8000/tasklists/3

Request -
{
  "name": "Need to Complete TOTO Golang App",
  "description": "Need to Complete TOTO Golang App",
  "status": "Created"
}

Response -
{
    "id": 3,
    "name": "Office TODO List",
    "description": "Office TODO List",
    "tasks": [
        {
            "id": 3,
            "name": "Complete Client Onboarding",
            "description": "Complete Client Onboarding",
            "status": "Created",
            "taskListId": 3
        },
        {
            "id": 4,
            "name": "Client Meeting with Rajesh@12",
            "description": "Api Discusssion",
            "status": "Created",
            "taskListId": 3
        }
    ]
}
````

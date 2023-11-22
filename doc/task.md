# Task API Spec

## Create a New Task
Endpoint : POST /api/task

Request Header :

- BEARER_TOKEN : Token (Mandatory) 

Description: endpoint to create a new task that belong to spesific user that login

Request Body (Success) :

```json
{
    "title": "Task A", 
    "description":" this is Task A", 
    "due_date":"2023-02-12T23:53:15Z"
}
```

Response Body(Success)

```json
{
    {
    "code": 200,
    "status": "success",
    "data": {
        "id": 0,
        "title": "Task B2",
        "description": " this is Task A",
        "status": "pending",
        "user": {
            "id": 0,
            "name": "",
            "email": "",
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "created_at": "2023-02-12T23:53:15Z"
    }
}
}
```

Response Body (Failed, 401) :

```json
{
    "code": 401,
    "status": "failure",
    "message": "Your token is expired"
}
```

Response Body (Failed, 400) :
```json
{
    "code": 400,
    "status": "failed",
    "message": "Bad Request",
    "errors": {
        "id": "Title already exist"
    }
}
```

## Update Existing Tasks
Endpoint: PUT /api/task/:id

Request Header :

- BEARER_TOKEN : Token (Mandatory) 

Description: endpoint to update existing tasks and set the task status to 'completed'. and user can only update the task status if the status belong to them

Request Body (Success) :

```json
{}
```

Response Body(Success)

```json
{
    "code": 200,
    "status": "success",
    "data": {
        "id": 1,
        "title": "Task A",
        "description": " this is Task A",
        "status": "completed",
        "user": {
            "id": 1,
            "name": "Dirga Meligo",
            "email": "dirga@gmail.com",
            "created_at": "2023-02-12T23:53:15Z",
            "updated_at": "0001-01-01T00:00:00Z"
        },
        "created_at": "2023-02-12T23:53:15Z"
    }
}
```

Response Body (Failed, 401) :

```json
{
    "code": 401,
    "status": "failure",
    "message": "Your token is expired"
}
```

Response Body (Failed, 400) :

```json
{
    "code": 400,
    "status": "failed",
    "message": "Bad Request",
    "errors": {
        "id": "The task is not belong to this user"
    }
}
```
Response Body (Failed, 400) :

```json
{
    "code": 400,
    "status": "failed",
    "message": "Bad Request",
    "errors": {
        "status": "The task should be pending"
    }
}
```


## List Existing Tasks
Endpoint: GET /api/task

Querystring:

    - per_page=10 -> show all list with limit 10

    - page=1 -> show list in first 10 and so on

    - search:"{{ search by title }}" -> you can search task by it's title

Request Header :

- BEARER_TOKEN : Token (Mandatory) 

Description: show all list of task for all user. you can use paginagination for better experience

Response Body (Success) :

```json
{
  {
    "code": 200,
    "status": "success",
    "data": [
        {
            "id": 1,
            "title": "Task A",
            "description": " this is Task A",
            "status": "completed",
            "user": {
                "id": 1,
                "name": "Dirga Meligo",
                "email": "dirga@gmail.com",
                "created_at": "2023-02-12T23:53:15Z",
                "updated_at": "0001-01-01T00:00:00Z"
            },
            "created_at": "2023-02-12T23:53:15Z"
        },
        {
            "id": 2,
            "title": "Task B",
            "description": " this is Task A",
            "status": "completed",
            "user": {
                "id": 2,
                "name": "Dirga Fithub",
                "email": "dirga@example.com",
                "created_at": "2023-11-21T15:44:05Z",
                "updated_at": "0001-01-01T00:00:00Z"
            },
            "created_at": "2023-02-12T23:53:15Z"
        },
        {
            "id": 3,
            "title": "Task B2",
            "description": " this is Task A",
            "status": "pending",
            "user": {
                "id": 1,
                "name": "Dirga Meligo",
                "email": "dirga@gmail.com",
                "created_at": "2023-02-12T23:53:15Z",
                "updated_at": "0001-01-01T00:00:00Z"
            },
            "created_at": "2023-02-12T23:53:15Z"
        }
    ],
    "total": 3,
    "page": 1,
    "per_page": 10
}
}
```

Response Body (Failed, 401) :

```json
{
    "code": 401,
    "status": "failure",
    "message": "Your token is expired"
}
```
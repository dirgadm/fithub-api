# Task API Spec

## Create a New Task
Endpoint : POST /api/task

Request Header :

- BEARER_TOKEN : Token (Mandatory) 

Request Body (Success) :

```json
{
    "title": "Task A", 
    "description":" this is Task A", 
    "due_date":
}
```

Response Body(Success)

```json
{
    "data":"OK"
}
```

Response Body (Failed, 401) :

```json
{
    "data":"Failed"
}
```

## Update Existing Tasks
Endpoint: PUT /api/task/:id

Request Header :

- BEARER_TOKEN : Token (Mandatory) 

Request Body (Success) :

```json
{
    "title": "Task A", 
    "description":" this is Task A", 
    "due_date":
}
```

Response Body(Success)

```json
{
    "data":"OK"
}
```

Response Body (Failed, 401) :

```json
{
    "data":"Failed"
}
```


## List Existing Tasks
Endpoint: GET /api/task

Querystring: perpage=10,page=0,search:"{{ search by title }}",order_by

Request Header :

- BEARER_TOKEN : Token (Mandatory) 

Response Body (Failed, 401) :

Response Body (Success) :

```json
{
  "data" : {
  }
}
```

```json
{
  "errors" : "Unauthorized"
}
```
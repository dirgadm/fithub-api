# FITHUB API ASSESSMENT

### Technology Stack:
| Teknologi   | Version | Link |
| ----------- | ---------------- | ------------------- |
| Golang      | v1.19 or later   | [Go Download](https://go.dev/dl)  |
| Go Echo Framework     | v4     | [Echo Installation](https://echo.labstack.com/guide/#installation) | 
| MySql | v5.7 or later |  |
| Docker | v24.0.6 or later |  |
| Docker compose| v12.21.0 or later |  |
<br>

## To Do
    install docker dan docker-compose
    Install postman
    Install git
    clone repo [https://github.com/dirgadm/fithub-api.git]

## Running Server
    1. command: *docker compose up -d*, Running compose yaml file in background side, and then do the migration to the mysql. the file path is in .database.sql
    2. command: *go run main.go* , Running server in 8000 port
    3. command: *go test* , go to the repository folder (./internal/reppository) and then do the unit test
    4. command: *go test -cover*, in folder you want to do unit test, you can see the coverage of unit test
    ```

## Endpoint Testing 
    - available in `./doc/Fithub.postman_collection.json` and ready to import to postman
    - base_url: http://localhost:8000/v1

## Technical Documentation
- There are file task.md and user.md in folder .doc
- There are consist of Technical specification for each endpoint in this API, and also the implementation

## List Endpoint
### User
1. [METHOD:POST] Register: User can register and create new user and store in database
2. [METHOD:POST] Login: Login to the system and provide BEARER token(expired in 24 hours) for all endpoint from task endpoint

### Task
1. [METHOD:GET] List: show all list of task for all user. you can use paginagination for better experience
2. [METHOD:POST] Create: endpoint to create a new task that belong to spesific user that login, and store it to database
3. [METHOD:PUT] Update: endpoint to update existing tasks and set the task status to 'completed'. and user can only update the task status if the status belong to them. and store it to database
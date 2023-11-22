# User API SPEC

## Register User

Endpoint:  POST /v1/auth/register

Description: Create new user for login to the system

Request body:
```json
{
    "email":"dirga@example.com",
    "password":"rahasia",
    "name":"dirga",
    "phone":"+6285319076822"
}
```

Response Body (Success) :

```json
{
    "code": 200,
    "status": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJleHBpcmVfYXQiOjE3MDA2MjE5NDZ9.FQK4X0R-NWfV9Yf3nsACfcgO5SXy8myJMoOHH8SzKew",
        "expired_at": "2023-11-22T09:59:06.816092409+07:00",
        "user": {
            "id": 0,
            "name": "Dirga Fithub",
            "email": "dirga2@example.com",
            "created_at": "2023-11-22T08:59:06.81098641+07:00",
            "updated_at": "2023-11-22T08:59:06.81098651+07:00"
        }
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
        "email": "The email is already exists"
    }
}
```

## Login User

Endpoint:  POST /v1/auth/login

Description: Login to the system and provide BEARER token(expired in 24 hours) for all endpoint from task endpoint

Request body:
```json
{
    "email":"dirga@example.com",
    "password":"rahasia"
}
```

Response Body (Success) :

```json
{
    "code": 200,
    "status": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJleHBpcmVfYXQiOjE3MDA2MjE5Njd9.f0GAqx-6tPXabLay4mrjoI-bekI7uU0UJHRPnHsa9Lc",
        "expired_at": "2023-11-22T09:59:27.089727642+07:00",
        "user": {
            "id": 2,
            "name": "Dirga Fithub",
            "email": "dirga@example.com",
            "created_at": "2023-11-21T15:44:05Z",
            "updated_at": "2023-11-21T15:44:05Z"
        }
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
        "password": "The password is not match"
    }
}
```
# User API SPEC

## Register User

Endpoint:  POST /v1/auth/register

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
  "data" : "OK"
}
```

Response Body (Failed, 400) :

```json
{
  "errors" : "Bad Request"
}
```

## Login User

Endpoint:  POST /v1/auth/login

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
  "data" : {
    "token" : "TOKEN",
    "expiredAt" : 2342342423423 // milliseconds
  }
}
```

Response Body (Failed, 401) :

```json
{
  "errors" : "Unauthorized"
}
```
# bookstore_users-api
USERS API

[//]: # (Format json)

#### error

```json
{
  "message" : "user 123 not found",
  "status" : 404,
  "error" : "not_found"
}
```

```json
{
  "message" : "invalid json body",
  "status" : 400,
  "error" : "bad_request"
}
```

```json
{
  "message" : "database is down",
  "status" : 500,
  "error" : "internal_server_error"
}
```
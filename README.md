# Architecture
- Frontend is written in Javascript using React.
- Backend is written in Go.
- Frontend and Backend are in the same repository but are separately developed  

# How to start
## Local Development
```sh
make up
make mod
make migrate-up
make be-run
```
Then go to `http://127.0.0.1:3000`


# Backend
## Spec
### Auth
- (Doing) JWT in Cookie. ref: https://medium.com/@bzzs.1120/jwt-for-session-and-csrf-c13c2885320e

### Response Format
- (Doing) Basically follow JSON API. ref: https://jsonapi.org/format/

```json
{
  "meta": {
    "code": 20001,
    "message": "success"
  },
  "data": {
    "id": 1,
    "title": "sample"
  }
}
```

```json
{
  "meta": {
    "code": 20002,
    "message": "success"
  },
  "data": [
    {
      "id": 1,
      "title": "sample"
    }
  ]
}
```

```json
{
  "meta": {
    "code": 42201,
    "message": "failure"
  },
  "errors": [
    {
      "source": { "pointer": "/data/title" },
      "detail": "the field is required"
    }
  ]
}
```

## Technical Document
TODO

## Directory Structure
TODO

# Frontend
## Technical Document
TODO

## Directory Structure
TODO

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
    "code": 20001,
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

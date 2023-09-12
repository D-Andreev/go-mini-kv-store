# go-mini-kv-store

## Usage

### Put

```curl
curl -X PUT -H "Content-Type: application/json" -d '{"key": "mykey", "value": "myval24"}' http://localhost:3000/mykey
```

### Get

```curl
curl http://localhost:3000/mykey
```

### Delete

```curl
curl -X DELETE http://localhost:3000/mykey
```

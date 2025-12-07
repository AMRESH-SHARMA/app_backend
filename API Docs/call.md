# Call APIs

---

## POST `/api/v1/call/start`

### Request
```json
{
  "callerId": "1",
  "calleeId": "2",
  "channel": "test123"
}
```

### Response
```json
{
  "message": "Call started",
  "data": {
    "callId": "uuid",
    "status": "RINGING"
  }
}
```

### Errors
```json
{ "error": "Invalid request" }
```

---

## POST `/api/v1/call/accept`

### Request
```json
{
  "callId": "uuid"
}
```

### Response
```json
{
  "message": "Call accepted",
  "data": {
    "callId": "uuid",
    "status": "ACCEPTED",
    "channel": "test123"
  }
}
```

### Errors
```json
{ "error": "Invalid request" }
```
```json
{ "error": "Call not found" }
```

---

## POST `/api/v1/call/reject`

### Request
```json
{
  "callId": "uuid"
}
```

### Response
```json
{
  "message": "Call rejected",
  "data": {
    "callId": "uuid",
    "status": "REJECTED"
  }
}
```

### Errors
```json
{ "error": "Invalid request" }
```
```json
{ "error": "Call not found" }
```

---

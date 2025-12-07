# RTC Token APIs

---

## POST `/api/v1/rtc/token`

### Request
```json
{
  "channel": "test123",
  "uid": "user-123"
}
```

### Response
```json
{
  "message": "Token generated",
  "data": {
    "token": "agora-token-string"
  }
}
```

### Errors
```json
{ "error": "Invalid request body" }
```
```json
{ "error": "Agora environment variables missing" }
```
```json
{ "error": "Token generation failed" }
```

---

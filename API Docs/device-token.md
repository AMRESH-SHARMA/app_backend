# Device Token APIs

---

## POST `/api/v1/user/device-token`

### Request
```json
{
  "userId": "1",
  "deviceToken": "fcm-token-here"
}
```

### Response
```json
{
  "message": "Device token updated",
  "data": {
    "userId": "1",
    "deviceToken": "fcm-token-here"
  }
}
```

### Errors
```json
{ "error": "UserId and DeviceToken required" }
```
```json
{ "error": "Failed to update device token" }
```

---

## DELETE `/api/v1/user/device-token`

### Request
```json
{
  "userId": "1"
}
```

### Response
```json
{
  "message": "Device token removed"
}
```

### Errors
```json
{ "error": "UserId required" }
```

---

## POST `/api/v1/user/device-token/refresh`

### Request
```json
{
  "userId": "1",
  "oldToken": "previous-token",
  "newToken": "new-fcm-token"
}
```

### Response
```json
{
  "message": "Device token refreshed",
  "data": {
    "userId": "1",
    "deviceToken": "new-fcm-token"
  }
}
```

### Errors
```json
{ "error": "UserId and NewToken required" }
```
```json
{ "error": "Failed to refresh device token" }
```

---

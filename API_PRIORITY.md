# API Priority List for Listener App Backend

## High Priority
1. **User Registration & Login (with OTP verification)**
   - POST /auth/register
   - POST /auth/login
   - POST /auth/refresh-token
   - POST /auth/verify-otp
   - POST /auth/resend-otp

   **CALL**
   POST /call/start
   Body:
   - callerId
   - calleeId
   - channelName

   POST /call/accept
   Body:
   - callId

   POST /call/reject
   Body:
   - callId

   Cancel (caller ends before callee picks up)
   POST /call/cancel
   Body:
   - callId

   End Call (after connected)
   POST /call/end
   Body:
   - callId

   Missed / Timeout (optional)
   POST /call/timeout
   Body:
   - callId


   Optional Presence APIs (Online/Offline)
   POST /presence/update
   GET  /presence/{userId}

   Call History APIs (Optional but common)
   GET  /calls/history?userId={id}
   GET  /calls/{callId}

   Notification APIs (Optional)  
   POST /notifications/send

start

accept

reject

busy state

missed call

end call

offline notifications

presence hooks

listener availability

RTC integration

Firebase push events

WebSocket signaling

   **Backend Files:**
   token service
   call signaling
   call storage
   user permission checks

   **Agora Token**
   POST /agora/token
      channelName={channel}&uid={uid}&role={publisher/subscriber}

   **Accept call: Flow**
   updates call status
   notifies the caller via push or websocket
   Then the app:
   requests Agora token from your backend
   joins Agora channel

   🧠 Flow in your backend
         user tries to start call:

         user calls /call/start

         backend checks permissions:

         is caller authenticated?

         is callee valid?

         is callee online?

         is listener busy?

         if allowed, backend:

         creates call record

         sends push notification

         returns call_id

         user accepts:

         backend checks if call_id exists

         check if still valid

         then generate Agora token

         return token

2. **User Profile**
   - GET /user/profile
   - PUT /user/profile

3. **Balance Management**
   - GET /user/balance
   - POST /user/balance/topup
   - GET /user/transactions

4. **Listener Discovery & Details**
   - GET /listeners
   - GET /listeners/{id}

5. **Session Management**
   - POST /sessions/start (initiate paid session with listener)
   - POST /sessions/end (end session, calculate cost, update balance)
   - GET /sessions/history

6. **Payment Integration**
   - POST /payments/initiate
   - POST /payments/webhook (handle payment gateway callbacks)

## Medium Priority
7. **Agora Integration (real-time audio/video)**
   - POST /agora/token (get RTC token for session)
   - POST /agora/callback (handle events from Agora)

8. **Notifications**
   - POST /notifications/send
   - GET /notifications

9. **Listener Management (for listeners/admin)**
   - POST /listeners/register
   - PUT /listeners/profile
   - GET /listeners/sessions

## Low Priority
10. **Admin APIs**
    - GET /admin/users
    - GET /admin/listeners
    - GET /admin/sessions
    - POST /admin/ban-user

11. **Support/Feedback**
    - POST /support/ticket
    - GET /support/tickets

12. **Miscellaneous**
    - GET /health
    - GET /version

---

## Suggested Project Modules

- **auth** – User registration, login, OTP verification, password reset
- **user** – User profile, balance, transaction history, settings
- **listener** – Listener registration, profile, availability, session history
- **session** – Session creation, tracking, ending, billing, history
- **payment** – Payment initiation, webhook/callback handling, top-up, refunds
- **agora** – Real-time audio/video token generation, event callbacks
- **notification** – Push, SMS, or email notifications for events (OTP, session start/end, etc.)
- **admin** – Admin dashboard, user/listener/session management, analytics
- **support** – Support ticket creation, FAQ, feedback
- **config** – Environment/configuration management (Viper, .env, etc.)
- **middleware** – Auth, logging, error handling, rate limiting, CORS, etc.
- **util** – Utility functions (helpers, validators, etc.)
- **repository** – Database access and queries (user, session, payment, etc.)
- **service** – Business logic layer (optional, for clean architecture)
- **route** – Route registration and grouping
- **analytics** – Usage tracking, reporting, metrics
- **referral** – Referral codes, bonuses, affiliate tracking
- **promotion** – Coupons, discounts, special offers
- **audit** – Audit logs for sensitive actions
- **test** – Test routes, mocks, or integration test helpers

Use this section to plan and track your project structure and module responsibilities.

## API
http://127.0.0.1:3001/api/v1/listeners/
http://127.0.0.1:3001/api/v1/listeners/gsearch?lang=hindi&gender=male


UI (Compose)
   ↓
ViewModel
   ↓
UseCase
   ↓
Repository
   ↓
Remote API (Retrofit)
   ↓
Backend response → DTO
   ↑
Mapper converts DTO → Domain model
   ↑
ViewModel updates StateFlow
   ↑
Compose recomposes UI


## Loading Image optimize:
If you want, I can show how to:

preload images

use Coil image loader

show shimmer

prefetch avatar on splash

cache aggressively like WhatsApp
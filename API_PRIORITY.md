# API Priority List for Listener App Backend

## High Priority
1. **User Registration & Login (with OTP verification)**
   - POST /auth/register
   - POST /auth/login
   - POST /auth/verify-otp
   - POST /auth/resend-otp

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

Use this file to track what has been built and what needs to be built for the backend.

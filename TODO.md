1. New User
   Register a new account in the system. (DONE)

2. User (Attendee/Presenter)
   Account Management
   ✔️ Login to the system (DONE)
   ✔️ Edit their profile (DONE)
   ✔️ View other users’ profiles (DONE)

   Session Interaction
   ✔️ View all conference sessions (DONE)
   ✔️ Leave feedback on sessions (DONE)
   ✔️ Register for sessions only if: (DONE)
       - Seats are available.
       - No overlapping time slots (one session per time period).
   Session Proposals
   ✔️ Create, edit, or delete their session proposals (DONE)
   ✔️ Submit only one proposal within a specific time period. (DONE)

3. Event Coordinator
   Session Management
   ✔️ View all session proposals
   ✔️ Accept or reject user-submitted proposals
   ✔️ Remove sessions (e.g., due to policy violations or cancellations)

   Feedback Moderation
   ✔️ Remove inappropriate user feedback

4. Admin
   User Management
   ✔️ Add new Event Coordinators
   ✔️ Remove users or Event Coordinators


Key Constraints to Implement
----------------------------
Session Registration:
- Users cannot book overlapping sessions.
- Registration is only allowed during the conference’s open enrollment period.

Session Proposals:
- Users may submit only one proposal per designated time window (e.g., per conference cycle). (DONE)

Permissions:
- Only Admins can assign/revoke Event Coordinator roles.
- Event Coordinators cannot alter user accounts (only Admins can remove users).
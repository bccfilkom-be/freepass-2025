# BCC University

### ⚠️⚠️⚠️

```
Submissions from 2024 students will have much higher priority than submissions from 2023, SAP, or higher students.
Please take note of this before planning to attempt this freepass challenge.
```

## 💌 Invitation Letter

Embracing the ever-evolving conference management landscape, we recognize the need for a seamless and engaging experience in academic meetings especially in the organizational space. We are embarking on an innovative project to transform the way conferences are hosted and experienced, and we want you to be a part of this journey!

We aim to create a dynamic conference platform that revolutionizes session management, attendee engagement, and administrative oversight. Your contributions will help shape the future of conference management. Together, we can create a platform that enhances knowledge sharing and professional networking while maintaining the highest standards of academic discourse.

Join us in revolutionizing the conference experience. Your insights and expertise are key to making this transformation happen!

## **⭐** Minimum Viable Product (MVP)

As we have mentioned earlier, we need technology that can support BCC Conference in the future. Please consider these features below:

- New user can register account to the system ✔️
- User can login to the system ✔️
- User can edit their profile account ✔️
- User can view all conference sessions ✔️
- User can leave feedback on sessions ✔️
- User can view other user's profile ✔️
- Users can register for sessions during the conference registration period if seats are available ✔️
- Users can only register for one session within a time period ✔️
- Users can create, edit, delete their session proposals ✔️
- Users can only create one session proposal within a time period ✔️
- Users can edit, delete their session ✔️
- Event Coordinator can view all session proposals ✔️
- Event Coordinator can accept or reject user session proposals ✔️
- Event Coordinator can remove sessions ✔️
- Event Coordinator can remove user feedback ✔️
- Admin can add new event coordinators ✔️
- Admin can remove users/event coordinators ✔️

## **🌎** Service Implementation

```
GIVEN => I am a new user
WHEN  => I register to the system
THEN  => System will record and return the user's registration details

GIVEN => I am a user
WHEN  => I log in to the system
THEN  => System will authenticate and grant access based on user credentials

GIVEN => I am a user
WHEN  => I edit my profile account
THEN  => The system will update my account with the new information

GIVEN => I am a user
WHEN  => I view all available conference's sessions
THEN  => System will display all conference sessions with their details

GIVEN => I am a user
WHEN  => I leave feedback on a session
THEN  => System will record my feedback and display it under the session

GIVEN => I am a user
WHEN  => I view other user's profiles
THEN  => System will show the information of other user's profiles

GIVEN => I am a user
WHEN  => I register for conference's sessions
THEN  => System will confirm my registration for the selected session

GIVEN => I am a user
WHEN  => I create a new session proposal
THEN  => System will record and confirm the session creation

GIVEN => I am a user
WHEN => I see my session's proposal details
THEN => System will display my session's proposal details

GIVEN => I am a user
WHEN  => I update my session's proposal details
THEN  => System will apply the changes and confirm the update

GIVEN => I am a user
WHEN  => I delete my session's proposal
THEN  => System will remove the session's proposal

GIVEN => I am a user
WHEN => I see my session details
THEN => System will display my session details

GIVEN => I am a user
WHEN  => I update my session details
THEN  => System will apply the changes and confirm the update

GIVEN => I am a user
WHEN  => I delete my session
THEN  => System will remove the session

GIVEN => I am an event coordinator
WHEN  => I view session proposals
THEN  => System will display all submitted session proposals

GIVEN => I am an event coordinator
WHEN  => I accept or reject the session proposal
THEN  => The system will make the session either be available or unavailable

GIVEN => I am an event coordinator
WHEN  => I remove a session
THEN  => System will delete the session

GIVEN => I am an event coordinator
WHEN  => I remove user feedback
THEN  => System will delete the feedback from the session

GIVEN => I am an admin
WHEN  => I add new event coordinator
THEN  => System will make the account to the system

GIVEN => I am an admin
WHEN  => I remove a user or event coordinator
THEN  => System will delete the account from the system
```

## **👪** Entities and Actors

### Entities

1. **User**
   - `userid`: Unique identifier for the user
   - `username`: Username of the user
   - `email`: Email address of the user
   - `password`: Hashed password of the user
   - `role`: Role of the user (e.g., user, event_coordinator, admin)
   - `created_at`: Timestamp when the user was created

2. **Session**
   - `sessionid`: Unique identifier for the session
   - `title`: Title of the session
   - `description`: Description of the session
   - `speaker`: Speaker of the session
   - `start_time`: Start time of the session
   - `end_time`: End time of the session
   - `max_seats`: Maximum number of seats available for the session
   - `created_by`: User who created the session
   - `created_at`: Timestamp when the session was created

3. **Session Proposal**
   - `proposalid`: Unique identifier for the session proposal
   - `title`: Title of the proposed session
   - `description`: Description of the proposed session
   - `speaker`: Speaker of the proposed session
   - `start_time`: Proposed start time of the session
   - `end_time`: Proposed end time of the session
   - `max_seats`: Maximum number of seats for the proposed session
   - `status`: Status of the proposal (e.g., pending, accepted, rejected)
   - `userid`: User who proposed the session
   - `proposed_at`: Timestamp when the proposal was created

4. **Feedback**
   - `feedbackid`: Unique identifier for the feedback
   - `sessionid`: Session for which the feedback is given
   - `userid`: User who gave the feedback
   - `comment`: Comment provided by the user
   - `rating`: Rating provided by the user (1-5)
   - `created_at`: Timestamp when the feedback was created

5. **Session Registration**
   - `registrationid`: Unique identifier for the session registration
   - `sessionid`: Session for which the user registered
   - `userid`: User who registered for the session
   - `registered_at`: Timestamp when the registration was created

### Actors

1. **User**
   - Can register an account
   - Can log in and log out
   - Can edit their profile
   - Can view all conference sessions
   - Can leave feedback on sessions
   - Can view other users' profiles
   - Can register for sessions if seats are available
   - Can create, edit, and delete their session proposals
   - Can edit and delete their sessions

2. **Event Coordinator**
   - Can view all session proposals
   - Can accept or reject session proposals
   - Can remove sessions
   - Can remove user feedback

3. **Admin**
   - Can add new event coordinators
   - Can remove users and event coordinators

## **📘** References

You might be overwhelmed by these requirements. Don't worry, here's a list of some tools that you could use (it's not required to use all of them nor any of them):

1. [Example Project](https://github.com/meong1234/fintech)
2. [Git](https://try.github.io/)
3. [Cheatsheets](https://devhints.io/)
4. [REST API](https://restfulapi.net/)
5. [Insomnia REST Client](https://insomnia.rest/)
6. [Test-Driven Development](https://www.freecodecamp.org/news/test-driven-development-what-it-is-and-what-it-is-not-41fa6bca02a2/)
7. [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
8. [GraphQL](https://graphql.org/)
9. [gRPC](https://grpc.io/)
10. [Docker Compose](https://docs.docker.com/compose/install/)

## **🔪** Accepted Weapons

> BEFORE CHOOSING YOUR LANGUAGE, PLEASE VISIT OUR [CONVENTION](CONVENTION.md) ON THIS PROJECT
>
> **Any code that did not follow the convention will be rejected!**
>
> 1. Golang (preferred)
> 2. Java
> 3. NodeJS
> 4. PHP

You are welcome to use any libraries or frameworks, but we appreciate if you use the popular ones.

## **🎒** Tasks

```
The implementation of this project MUST be in the form of a REST, gRPC, or GraphQL API (choose AT LEAST one type).
```

1. Fork this repository
2. Follow the project convention
3. Finish all service implementations
4. Write the installation guide of your back-end service in the section below

## **🧪** API Documentation
[Api documentation](https://documenter.getpostman.com/view/37017335/2sAYQiBTfu)

## **🧪** API Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-username/freepass-2025.git
   cd freepass-2025
   ```

2. **Install dependencies:**
   ```sh
   npm install
   ```

3. **Set up environment variables:**
   Create a `.env` file in the root directory and add the following:
   ```env
   PORT=3001
   JWT_SECRET=your_jwt_secret
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_HOST=your_db_host
   DB_PORT=5432
   DB_DATABASE=your_db_name
   ```

4. **Run the database migrations:**
   Ensure your PostgreSQL database is running and execute the necessary SQL scripts to create the tables.

5. **Start the server:**
   ```sh
   node app.js
   ```

6. **Access the API:**
   Open your browser or API client (e.g., Postman, Insomnia) and navigate to `http://localhost:3001`.



## **📞** Contact

Have any questions? You can contact either [Tyo](https://www.instagram.com/nandanatyo/) or [Ilham](https://www.instagram.com/iilham_akbar/).

## **🎁** Submission

Please follow the instructions on the [Contributing guide](CONTRIBUTING.md).

![cheers](https://gifsec.com/wp-content/uploads/2022/10/cheers-gif-11.gif)

> This is not the only way to join us.
>
> **But, this is the _one and only way_ to instantly pass.**

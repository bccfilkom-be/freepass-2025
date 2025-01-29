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

We want to see your perspective about these problems. You can define various types of entities or actors. One thing for sure, there is no
true or false statement to define the entities. As long as the results are understandable, then go for it! 🚀
```
(Entities)
User : Represents an individuals who interact with the system. User have different roles. Attributes include
- user_id
- name
- email
- password
- address
- created_at
- updated_at
- role_id 

Role : Represent the role that user have. Attributes include :
- role_id
- role_name

Session : Represent the conference that user can register to. Attributes include :
- session_id
- title
- description
- session_owner
- time_slot
- max_seats
- available_seats
- created_at
- updated_at
- proposal_id

Session Registration : Represent session that user already register to. Attributes include :
- registration_id
- timestamp
- session_id
- user_id

Session Proposal : Represent a proposal of a session made by user. Attributes include :
- proposal_id
- title
- description
- time_slot
- status
- created_at
- updated_at
- user_id

Feedback : Represent a feedback made by user on session. Attributes include :
- feedback_id
- comment
- created_at
- session_id
- user_id

(Actor)
- User : An individuals who interact with the system. User can have a different role
- Event Coordinator : A user who can manage session and feedback of the user.
- Admin : A user who manage reguler user and event coordinator
```
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

## **🧪** API Installation

> Write how to run your service in local or development environment here. If you use Docker to serve your DBMS or your server, you will receive bonus points for your submission.

## Running the code
Before you begin, ensure you have the following installed on your machine
1. Go
2. Gin
3. Gorm
4. MariaDB

Here are the step :
1. Make sure your MariaDB/MySQL is configured correctly
2. Clone this repository or download manually
3. Navigate to the project directory
```
cd freepass-2025\cmd\app
```
4. Install Gin and GORM
```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
```
5. Configure your credential and databse configuration in your .env
6. Start the application
```
go run main.go
```

## **📞** Contact

Have any questions? You can contact either [Tyo](https://www.instagram.com/nandanatyo/) or [Ilham](https://www.instagram.com/iilham_akbar/).

## **🎁** Submission

Please follow the instructions on the [Contributing guide](CONTRIBUTING.md).

![cheers](https://gifsec.com/wp-content/uploads/2022/10/cheers-gif-11.gif)

> This is not the only way to join us.
>
> **But, this is the _one and only way_ to instantly pass.**

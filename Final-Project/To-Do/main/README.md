# To-Do list api

A simple RESTful API for task management. Allows you to create, delete users, create, update tasks for users, show lists of tasks and delete users or tasks.

## Beginning of work

These instructions will help you run a copy of the project on your local machine for development and testing.

### Prerequisites

To start this project you will need:

- Go (version 1.15 or higher)
- Postman (to send requests)
- PostgreSQL (to connect your database)
- Docker (for containerizing a Postgres database)

### Installation

Steps to set up a project locally:

1. Running a Postgres Database via Docker:
   docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
2. Installing dependencies Go:
   go mod tidy
3. Launching the application:
   go run . (to run all go files)

### Usage

API endpoints:
  User management:

    POST /users — creating a new user.
    GET /users — getting a list of all users.
    DELETE /users/{id} — deleting a user by ID (requires a JWT token in the Authorization header). 

  User Login:

    POST /login — user authentication and receiving a JWT token.

  Task management:

    POST /tasks — creating a new task (requires a JWT token in the Authorization header).
    GET /tasks — getting a list of all user tasks (requires a JWT token in the Authorization header).
    DELETE /tasks/{id} — deleting a task by ID (requires a JWT token in the Authorization header).
	PUT /tasks/{id} — updating a task by ID (requires a JWT token in the Authorization header). 

### Database structure

![Database structure](https://i.imgur.com/PsZrSTZ.png)

### Example requests:
  Creating a new user:
  ```
    POST /users
    Content-Type: application/json

    {
      "name": "John",
      "email": "john@example.com",
      "password": "password123"
    }
  ```

  Authentication and receipt of JWT token:
  ```
    POST /login
    Content-Type: application/json

    {
      "email": "john@example.com",
      "password": "password123"
    }
  ```
  Create a new task:
  ```
    POST /tasks
    Content-Type: application/json
    Authorization: Bearer <your_jwt_token>

    {
      "title": "New Task",
      "description": "Task description",
      "status": "new"
    }
  ```

  Getting a list of tasks:
  ```
    GET /tasks
    Content-Type: application/json
    Authorization: Bearer <your_jwt_token>
  ```

  Deleting a task:
  ```
    DELETE /tasks/{id}
    Content-Type: application/json
    Authorization: Bearer <your_jwt_token>
  ```
  
  Updating a task:
  ```
    PUT /tasks/{id}
    Content-Type: application/json
    Authorization: Bearer <your_jwt_token>
	
	{
    "title": "new title",
    "description": "new description",
    "status": "changed status"
    }
  ```

  Getting a list of users:
  ```
    GET /users
    Content-Type: application/json
  ```

  Deleting a user:
  ```
    DELETE /users/{id}
    Content-Type: application/json
    Authorization: Bearer <your_jwt_token>
  ```
    
### License:
- Skorokhod Maxim Andreevich 22B030614
- Zhunisov Ernur Erboluly 22B030365



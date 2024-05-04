# MessengerApp API Documentation

## Overview
MessengerApp - is a messaging API that allows users to register, log in, send messages, add friends, and view messages with various filtering and sorting options.

## Creators
- Skorokhod Maxim Andreevich 22B030614
- Zhunisov Ernur Erboluly 22B030365

## Technologies Used
- **Go (Golang)**: Main programming language.
- **Gin Framework**: HTTP web framework for creating APIs.
- **PostgreSQL**: Database to store all user data and messages.
- **JWT (JSON Web Tokens)**: Used for authentication and authorization.
- **Viper**: Configuration management library.

## Installation

### Prerequisites
- Install Go (version 1.x or higher).
- Install PostgreSQL and create a database for the application.
- Set up your Go environment and PostgreSQL database.

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/51mans0n/GoLang/tree/main/Project
2. Go to the project directory:
   ```bash
   cd messengerApp
3. Set up the config.yaml configuration file according to your database settings.
4. Install dependencies:
   ```bash
   go mod tidy
5. Run database migrations:
   ```bash
   go run migrations/migration.go
6. Start the server:
   ```bash
   go run main.go

## API Endpoints

### Authentication

- **Login**
    - **URL**: `/login`
    - **Method**: `POST`
    - **Example**: `http://localhost:8080/login`
    - **Data Params**:
      ```json
      {
        "username": "example",
        "password": "password"
      }
      ```
    - **Success Response**:
        - **Code**: 200
        - **Content**:
          ```json
          {
            "token": "jwt_token_here"
          }
          ```

- **Register**
    - **URL**: `/register`
    - **Method**: `POST`
    - **Example**: `http://localhost:8080/register`
    - **Data Params**:
      ```json
      {
        "username": "example",
        "password": "password"  
      }
      ```
    - **Success Response**:
        - **Code**: 201
        - **Content**:
          ```json
          {
            "message": "User registered successfully"
          }
          ```


### Friends

- **Add Friend**
    - **URL**: `/add-friend`
    - **Method**: `POST`
    - **Authorization**: Bearer Token
    - **Example**: `http://localhost:8080/add-friend`
    - **Data Params**:
      ```json
      {
        "friend_id": 10
      }
      ```
    - **Success Response**:
        - **Code**: 200
        - **Content**:
          ```json
          {
            "message": "Friend added successfully"
          }
          ```
          

### Messages

- **Send Message**
    - **URL**: `/send-message`
    - **Method**: `POST`
    - **Authorization**: Bearer Token
    - **Example**: `http://localhost:8080/send-message`
    - **Data Params**:
      ```json
      {
        "receiver_id": 10,
        "message": "Hello!"
      }
      ```
    - **Success Response**:
        - **Code**: 200
        - **Content**:
          ```json
          {
            "message": "Message sent successfully"
          }
          ```
          

### View Messages (Super Users Only)

- **Get Messages**
    - **URL**: `/messages`
    - **Method**: `GET`
    - **Authorization**: Bearer Token
    - **Query Params**: `page`, `pageSize`, `sortBy`, `sortDir`
    - **Example**: `http://localhost:8080/messages`
    - **Success Response**:
        - **Code**: 200
        - **Content**: List of all messages


- **Get Sorted Messages**
    - **URL**: `/messages`
    - **Method**: `GET`
    - **Authorization**: Bearer Token
    - **Query Params**: `page`, `pageSize`, `sortBy`, `sortDir`
    - **Example**: `http://localhost:8080/messages?page=1&pageSize=10&sortBy=timestamp&sortDir=asc`
    - **Success Response**:
        - **Code**: 200
        - **Content**: List of messages sorted by timestamp in ascending order


- **Get Paginated Messages**
    - **URL**: `/messages`
    - **Method**: `GET`
    - **Authorization**: Bearer Token
    - **Query Params**: `page`, `pageSize`
    - **Example**: `http://localhost:8080/messages?page=2&pageSize=2`
    - **Success Response**:
      - **Code**: 200
      - **Content**: Paginated list of messages


- **Get Filtered Messages**
    - **URL**: `/messages`
    - **Method**: `GET`
    - **Authorization**: Bearer Token
    - **Query Params**: `sender_id`
    - **Example**: `http://localhost:8080/messages?sender_id=30`
    - **Success Response**:
      - **Code**: 200
      - **Content**: List of messages filtered by sender ID

### User Management (Super Users Only)
- **Delete User**
    - **URL**: `/users/{id}`
    - **Method**: `DELETE`
    - **Authorization**: Bearer Token
    - **Example**: `http://localhost:8080/users/4`
    - **Success Response**:
        - **Code**: 200
        - **Content**: 
          ```json
          {
            "message": "User deleted successfully"
          }
          ```


- **Update User**
    - **URL**: `/users/{id}`
    - **Method**: `PUT`
    - **Authorization**: Bearer Token
    - **Data Params**:
      ```json
      {
        "username": "newusername",
        "password": "newpassword"
      }
      ```
    - **Example**: `http://localhost:8080/users/10`
    - **Success Response**:
        - **Code**: 200
        - **Content**:
          ```json
          {
            "message": "User updated successfully"
          }
          ```

**NOTE!** Super User - users (ID > 50 & prime number. For example 53). You need to add them using SQL requests.

## Errors
The API uses conventional HTTP response codes to indicate success or failure of an API request. In general:
- 200 OK - Everything worked as expected.
- 400 Bad Request - The request was unacceptable.
- 401 Unauthorized - The request lacks valid authentication credentials.
- 500 Internal Server Error - Something went wrong on our end.
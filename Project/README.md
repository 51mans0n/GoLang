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
    - **Data Params**:
      ```json
      {
        "username": "example",
        "password": "password",
        "role": "user"  
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

### Messages

- **Send Message**
    - **URL**: `/send-message`
    - **Method**: `POST`
    - **Authorization**: Bearer Token
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

- **Get Messages**
    - **URL**: `/messages`
    - **Method**: `GET`
    - **Authorization**: Bearer Token
    - **Query Params**: `page`, `pageSize`, `sortBy`, `sortDir`
    - **Success Response**:
        - **Code**: 200
        - **Content**: List of messages

### Friends

- **Add Friend**
    - **URL**: `/add-friend`
    - **Method**: `POST`
    - **Authorization**: Bearer Token
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

## Errors
The API uses conventional HTTP response codes to indicate success or failure of an API request. In general:
- 200 OK - Everything worked as expected.
- 400 Bad Request - The request was unacceptable.
- 401 Unauthorized - The request lacks valid authentication credentials.
- 500 Internal Server Error - Something went wrong on our end.
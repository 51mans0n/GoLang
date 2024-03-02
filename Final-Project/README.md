# Web Application for Registration and Authorization

This web application provides a basic user registration and authorization system using email and password. The application uses Go for the backend, PostgreSQL to store user data, and JWT to manage user sessions.

## Beginning of work

### Prerequisites

To run the application you will need the following components:

- Go (check installation with `go version` command)
- PostgreSQL (check installation with `psql --version`)
- Gorilla Mux library for Go (install with `go get -u github.com/gorilla/mux`)
- Library for working with JWT in Go (install with the command `go get -u github.com/dgrijalva/jwt-go`)
- Library for hashing passwords in Go (install with `go get -u golang.org/x/crypto/bcrypt`)

### Installation

1. Clone the repository to a local directory.
2. Make sure you have all prerequisites installed and configured.
3. Create a PostgreSQL database and table to store user data. Example SQL to create a table:

    ```sql
    CREATE TABLE auth (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL
    );
    ```

4. Edit the `main.go` file, specifying the correct parameters for connecting to your PostgreSQL database in the `host`, `port`, `user`, `password`, `dbname` constants.

### Launch

To start the web server, use the command:

```bash
go run main.go
```

The server will be available at `http://localhost:8080`.

## Usage

The application provides two main endpoints:

- `POST /register`: Accepts JSON with email and password to register a new user. Example request:

   ```json
   {
       "email": "user@example.com",
       "password": "yourpassword"
   }
   ```

- `POST /login`: Accepts JSON with email and password for user authentication. In case of successful authorization, it returns a JWT token. Example request:

   ```json
   {
       "email": "user@example.com",
       "password": "yourpassword"
   }
   ```

## Safety

User passwords are hashed using the `bcrypt` library before being saved to the database, ensuring sensitive information is stored securely.

## Contribution

Any contribution to the project is welcome. To propose changes, create a Pull Request.

## License

[Specify license type]

---

**Note**: Be sure to indicate the type of license under which you want to distribute your project in the "License" section.
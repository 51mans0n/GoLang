CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    description TEXT,
    status VARCHAR(50),
    user_id INTEGER REFERENCES users(id)
);
CREATE USER Simanson WITH PASSWORD '12345a';
GRANT ALL PRIVILEGES ON DATABASE ToDo TO Simanson;


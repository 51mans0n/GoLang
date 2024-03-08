package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/login", loginUser).Methods("POST")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("The server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createTask(w http.ResponseWriter, r *http.Request) {
	// Извлечение user_id из JWT токена
	user_id, err := extractUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var newTask Task
	err = json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := connectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO tasks (title, description, status, user_id) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, newTask.Title, newTask.Description, newTask.Status, user_id).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newTask.ID = id

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	user_id, err := extractUserIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	db := connectDB()
	defer db.Close()

	// Добавление условия WHERE user_id = $1 для фильтрации задач по пользователю
	rows, err := db.Query("SELECT id, title, description, status, user_id FROM tasks WHERE user_id = $1", user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Получаем идентификатор задачи из URL
	params := mux.Vars(r)
	taskID := params["id"]

	// Устанавливаем подключение к базе данных
	db := connectDB()
	defer db.Close()

	// Выполняем SQL запрос на удаление задачи по идентификатору
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем статус HTTP 204 No Content, как подтверждение успешного удаления
	w.WriteHeader(http.StatusNoContent)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	db := connectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, newUser.Name, newUser.Email, string(hashedPassword)).Scan(&id)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	newUser.ID = id
	newUser.Password = "" // Не возвращаем пароль

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, "Error reading users", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var creds User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := connectDB()
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT id, password_hash FROM users WHERE email = $1", creds.Email).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User is not found", http.StatusUnauthorized)
		} else {
			http.Error(w, "User request error", http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Генерация JWT
	token, err := generateJWT(user.ID)
	if err != nil {
		http.Error(w, "JWT generation error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Извлечение user_id из пути запроса
	params := mux.Vars(r)
	userID := params["id"]

	db := connectDB()
	defer db.Close()

	// Выполнение SQL запроса на удаление пользователя
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully."))
}

func generateJWT(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Создание утверждений JWT
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func extractUserIDFromToken(r *http.Request) (int, error) {
	// Получение токена из заголовка Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("Authorization header is missing")
	}

	// Проверка формата заголовка
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return 0, fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	tokenString := bearerToken[1]

	// Получение секретного ключа из переменной окружения
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set in the environment variables")
	}

	// Парсинг и валидация токена
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("Invalid token: %v", err)
	}

	// Извлечение и возврат user_id из токена
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, fmt.Errorf("Invalid token claims")
	}
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      int    `json:"user_id"`
}
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-"`
}

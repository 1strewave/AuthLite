package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Структура пользователя
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Инициализация базы данных
func InitDB(dataSource string) error {
	var err error
	DB, err = sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Создаем таблицу, если она не существует
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			token TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Функция для создания пользователя
func CreateUser(username, password string) (*User, error) {
	token := uuid.New().String() // Генерация токена

	stmt, err := DB.Prepare("INSERT INTO users (username, password, token) VALUES (?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(username, password, token)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return &User{
		ID:       id,
		Username: username,
		Password: password,
		Token:    token,
	}, nil
}

// Функция для получения пользователя по имени
func GetUserByUsername(username string) (*User, error) {
	row := DB.QueryRow("SELECT id, username, password, token FROM users WHERE username = ?", username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	return &user, nil
}

// Функция для получения пользователя по токену
func GetUserByToken(token string) (*User, error) {
	row := DB.QueryRow("SELECT id, username, password, token FROM users WHERE token = ?", token)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	return &user, nil
}

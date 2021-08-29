package api

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v4"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MSG struct {
	User    User   `json:"user,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllUsers(db *pgx.Conn) []User {
	var users []User
	rows, err := db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		fmt.Printf("Error gettings all Users: %v\n", err)
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Age)
		users = append(users, user)
	}

	return users
}

func GetUserById(db *pgx.Conn, id string) MSG{
	var user User
	var msg MSG
	err := db.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
		mess := "Not found"
		msg.Message = mess
		return msg
	}
	msg.User = user
	return msg
}

package api

import (
	"os"
	"fmt"
	"context"
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

type ConnectionRef struct {
	dbname string
	port   string
	host   string
	user   string
	psw    string
}

// create Connection Detils for InitConn
func GetRef(dbname string, port string) *ConnectionRef {
	return &ConnectionRef{dbname, port, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PSW")}
}

// initialize Connection to desired Postgres DB
func InitConn(c *ConnectionRef){
	pgxInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.user, c.psw, c.host, c.port, c.dbname)
	conn, err := pgx.Connect(context.Background(), pgxInfo)
	if err != nil {
		fmt.Printf("Error connecting: %v", err)
	} else {
		fmt.Println("Successfully connected to Postegsql!")
		db = conn
	}
}

// return Connection
func GetConn() *pgx.Conn {
	return db
}

// close Connection
func CloseConn(){
	db.Close(context.Background())
}
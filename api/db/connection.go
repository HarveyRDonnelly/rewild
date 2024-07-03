package db

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// Alias UUID type
type uuid_t = uuid.UUID

type Connection struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Gateway  *sql.DB
}

type Connectioner interface {
	Connect(Connection)
	GetUser(uuid_t)
}

func Connect(conn Connection) Connection {

	conn_str := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conn.Host, conn.Port, conn.User, conn.Password, conn.Database)

	db, err := sql.Open("postgres", conn_str)

	if err != nil {
		panic(err)
	}

	conn.Gateway = db

	return conn
}

func nullIDString(id uuid_t) sql.NullString {
	idString := sql.NullString{
		String: id.String(),
		Valid:  true,
	}
	if id == uuid.Nil {
		idString.Valid = false
	}
	return idString
}

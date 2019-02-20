package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type RealDatabase struct {
	connection *sqlx.DB
}

func NewDatabase() *RealDatabase {
	DB, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
	}

	return &RealDatabase{
		connection: DB,
	}
}
func (d *RealDatabase) GetUser(user string) string {

	return "real user"
}

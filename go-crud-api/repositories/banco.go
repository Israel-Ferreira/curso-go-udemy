package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(username string, password string, host string, port string, database string) (*sql.DB, error) {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, host, port, database)

	dbConn, err := sql.Open("mysql", connectionStr)

	if err != nil {
		return nil, err
	}

	if err = dbConn.Ping(); err != nil {
		return nil, err
	}

	return dbConn, nil

}

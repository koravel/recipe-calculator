package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

type SQLiteDriver struct {
	pathToDB string
	db       *sql.DB
}

func NewSQLiteDriver(pathToDB string) *SQLiteDriver {
	return &SQLiteDriver{
		pathToDB: pathToDB,
	}
}

func (driver *SQLiteDriver) Connect() error {
	db, err := sql.Open("sqlite", driver.pathToDB)
	if err != nil {
		return fmt.Errorf("Error while connecting to recipe db: %v", err)
	}

	driver.db = db
	return nil
}

func (driver *SQLiteDriver) Close() error {
	err := driver.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (driver *SQLiteDriver) Query(query string) (string, error) {
	var result string
	err := driver.db.QueryRow(query).Scan(&result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (driver *SQLiteDriver) Exec(query string) error {
	_, err := driver.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

package bowler

import (
	"database/sql"
	"fmt"
	"test/config"
)

type fieldMap struct {
	name      string
	fieldType string
}

var (
	conf        *config.Config
	knownTables []string
)

const (
	mySQLConnectionPattern = "%s:%s@%s/%s"
)

func init() {
	knownTables = make([]string, 0, 20)
}

func Configure(c *config.Config) error {
	conf = c
	db, err := connect()
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func connect() (*sql.DB, error) {
	cs := fmt.Sprintf(mySQLConnectionPattern, conf.User, conf.Password, conf.Host, conf.Name)
	tdb, err := sql.Open("mysql", cs)
	if err != nil {
		return nil, err
	}
	return tdb, nil
}

// Disconnect from the database
func disconnect(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}

func beginTransaction(db *sql.DB) (*sql.Tx, error) {
	return db.Begin()
}

func commitTransaction(tx *sql.Tx) {
	if tx != nil {
		tx.Commit()
	}
}

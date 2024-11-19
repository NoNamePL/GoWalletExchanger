package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/NoNamePL/GoWalletExchanger/iternal/config"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	// Connected to DB
	connStr := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open(cfg.Driver, connStr)
	if err != nil {
		log.Fatal(err)
	}

	// create user table
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS user (
			userID SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
		)
	`)

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New("can't create user table")
	}

	// Creating base tables in the database
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS wallet(
			walletId SERIAL PRIMARY KEY,
    		amount INT,
			username UNIQUE TEXT REFERENCES user (username),
		)
	`)

	if err != nil {
		return nil, errors.New("can't prepere query of wallet table")
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New("can't create wallet table")
	}

	// create table exchange
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS exchange(
			exchangeId SERIAL PRIMARY KEY,
			currency TEXT NOT NULL,
			rate INT NOT NULL,
		)
	`)

	if err != nil {
		return nil, errors.New("can't prepere query of exchange table")
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New("can't create exchange table")
	}

	

	// // Create index
	// stmt, err = db.Prepare(`
	// 	CREATE UNIQUE INDEX IF NOT EXISTS valIdx on wallet (valletId)
	// `)

	// if err != nil {
	// 	return nil, errors.New("can't prepere query of wallet table")
	// }

	// _, err = stmt.Exec()
	// if err != nil {
	// 	return nil, errors.New("can't create index")
	// }

	return db, nil
}

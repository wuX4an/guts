package database

import "log"

func Create() error {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	instruction := "CREATE TABLE IF NOT EXISTS users (name text, age integer)"
	_, err = db.Exec(instruction)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

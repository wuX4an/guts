package database

import (
	"fmt"
)

func Write(name string, age int) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	instruction := fmt.Sprintf("INSERT INTO users VALUES ('%s', %d)", name, age)
	_, err = db.Exec(instruction)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

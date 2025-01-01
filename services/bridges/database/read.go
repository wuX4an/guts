package database

import (
	"fmt"
	"log"
)

type streamer struct {
	name string
	age  int
}

// FIX: Return
func Read(name string) (string, error) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close the connection after the function finishes

	var streamers []streamer

	// Build the query to select all data from the table
	instruction := fmt.Sprintf("SELECT * FROM users WHERE name='%s'", name)
	rows, err := db.Query(instruction)
	if err != nil {
		return "", err
	}
	defer rows.Close() // Close the rows after iterating through them

	// Scan each row and append data to the streamers slice
	for rows.Next() {
		var s streamer
		err := rows.Scan(&s.name, &s.age)
		if err != nil {
			return "", err
		}
		streamers = append(streamers, s)
	}

	// Check for any errors encountered during iteration
	err = rows.Err()
	if err != nil {
		return "", err
	}
	if streamers == nil {
		return "", err
	}

	return streamers[0].name, nil
}

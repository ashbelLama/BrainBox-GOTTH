package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	database "github.com/ashbelLama/brainbox/db"
	"github.com/ashbelLama/brainbox/model"
)

func CheckDBConnection() error {
	if database.DB == nil {
		return errors.New("database connection not established")
	}
	return nil
}

func CheckBubbleContents(results *sql.Rows) bool {
	var bubble model.Bubble
	for results.Next() {
		err := results.Scan(&bubble.Id, &bubble.Title, &bubble.Description)
		if err != nil {
			return false
		}
	}
	return true
}

func GetBubble() ([]byte, error) {
	if err := CheckDBConnection(); err != nil {
		log.Fatal(err)
	}

	results, err := database.DB.Query("SELECT id, title, description FROM bubble")
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var bubble model.Bubble
		err = results.Scan(&bubble.Id, &bubble.Title, &bubble.Description)
		if err != nil {
			log.Fatal(err)
		}
		// Marshal bubble struct into json
		jsonData, err := json.Marshal(bubble)
		if err != nil {
			log.Fatal(err)
		}

		return jsonData, err
	}
	return nil, nil
}

func PostBubble(item model.Bubble) bool {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	}

	if err := CheckDBConnection(); err != nil {
		log.Fatal(err)
		return false
	}

	stmt, err := tx.Prepare("INSERT INTO bubble (title, description) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.Title, item.Description)
	if err != nil {
		log.Fatal(err)
		return false
	}

	tx.Commit()
	return true
}

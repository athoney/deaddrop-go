package db

import (
	"log"

	"github.com/andey-robins/deaddrop-go/logger"
)

// GetMessagesForUser assumes that a user has already been
// authenticated through a call to session.Authenticate(user)
// and then returns all the messages stored for that user
func GetMessagesForUser(user string) []string {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (data) FROM Messages
		WHERE recipient = (
			SELECT id FROM Users WHERE user = ?
		)
	`, user)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer rows.Close()

	// marshall rows into an array
	messages := make([]string, 0)
	for rows.Next() {
		var message string
		err := rows.Scan(&message)
		if err != nil {
			log.Fatalf("unable to scan row")
		}
		messages = append(messages, message)
	}
	// Log reading a message from a user that exists
	logger.LogReadMessage(user)
	return messages
}

// saveMessage will process the transaction to place a message
// into the database
func SaveMessage(message, recipient, sender string) {
	database := Connect().Db

	database.Exec(`
		INSERT INTO Messages (recipient, data)
		VALUES (
			(SELECT id FROM Users WHERE user = ?), 
			?
		);
	`, recipient, message)
	// Log sending message to a user that exists
	logger.LogSentMessage(recipient, sender)
}

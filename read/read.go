package read

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/session"
)

func ReadMessages(user string) {
	if !db.UserExists(user) {
		log.Println("Failed Read User: " + user + " cannot read messages for non existent user")
		log.Fatalf("User not recognized")
	}

	err := session.Authenticate(user)
	if err != nil {
		log.Println("Failed Read Password: " + user + " attempted to read messages with wrong password")
		log.Fatalf("Unable to authenticate user")
	}

	messages := db.GetMessagesForUser2(user)
	for _, message := range messages {
		if verify([]byte(message.Data), []byte(os.Getenv("KEY")), message.Hash) {
			fmt.Println(message.Sender + " sent: " + message.Data)
		} else {
			log.Println("MAC failure: '" + message.Data + "' from " + message.Sender + " cannot be verified!")
			fmt.Println("WARNING! Message cannot be authenticated!! " + message.Sender + " sent: " + message.Data)
		}
	}
}

func verify(msg, key []byte, hash string) bool {
	sig, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}

	mac := hmac.New(sha256.New, key)
	mac.Write(msg)

	return hmac.Equal(sig, mac.Sum(nil))
}

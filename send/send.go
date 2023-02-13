package send

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/andey-robins/deaddrop-go/db"
	"github.com/andey-robins/deaddrop-go/logger"
	"github.com/andey-robins/deaddrop-go/session"
)

// SendMessage takes a destination username and will
// prompt the user for a message to send to that user
func SendMessage(to, from string) {
	//Auth for sender
	if !db.NoUsers() && !db.UserExists(from) {
		log.Fatalf("Sender not recognized")
	}
	err := session.Authenticate(from)
	if err != nil {
		logger.LogFailedSentMessageSender(from)
		log.Fatalf("Unable to authenticate user")
	}

	if !db.UserExists(to) {
		logger.LogFailedSentMessage(to)
		log.Fatalf("Destination user does not exist")
	}

	message := getUserMessage()

	db.SaveMessage(message, to, from)
}

// getUserMessage prompts the user for the message to send
// and returns it
func getUserMessage() string {
	fmt.Println("Enter your message: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return text
}

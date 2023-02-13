package logger

import (
	"os"
	"time"
)

func addToLog(message string) {
	f, _ := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	f.WriteString(time.Now().Format(time.UnixDate) + " - " + message + "\n")

	f.Close()
}

func LogSentMessage(recipient, sender string) {
	message := "Send: " + recipient + " recieved a message from " + sender
	addToLog(message)
}

func LogFailedSentMessage(recipient string) {
	message := "Failed Send: recipient " + recipient + " is not a user"
	addToLog(message)
}

func LogFailedSentMessageSender(recipient string) {
	message := "Failed Send: " + recipient + " attempted to send a message with the wrong password"
	addToLog(message)
}

func LogReadMessage(recipient string) {
	message := "Read: " + recipient + " read their messages"
	addToLog(message)
}

func LogFailedReadUser(recipient string) {
	message := "Failed Read User: " + recipient + " cannot read messages for non existent user"
	addToLog(message)
}

func LogFailedReadPass(recipient string) {
	message := "Failed Read Password: " + recipient + " attempted to read messages with wrong password"
	addToLog(message)
}

func LogNewUser(creator, newUser string) {
	message := "New User: " + creator + " created " + newUser
	addToLog(message)
}

func LogFailedNewUser(creator string) {
	message := "Failed New User Password: " + creator + " attempted to make a new user with wrong password"
	addToLog(message)
}

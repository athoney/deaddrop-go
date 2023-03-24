# deaddrop-go

A deaddrop utility written in Go. Put files in a database behind a password to be retrieved at a later date.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:
- [Javascript](https://github.com/andey-robins/deaddrop-js)
- [Rust](https://github.com/andey-robins/deaddrop-rs)

## Versioning

`deaddrop-go` is built with:
- go version go1.19.4 linux/amd64

## Usage

`go run main.go --help` for instructions

Then run `go run main.go -new -user <username here>` and you will be prompted to create the initial password.

## Database

Data gets stored into the local database file dd.db. This file will not by synched to git repos. Delete this file if you don't set up a user properly on the first go

## Logging Strategy **(REVISED)**
To add logging functionality, I used the Golang log package. The log reflects the following actions:
- Sent Messages:
  - Successful message with sender and recipient
  - Failed send on account of non exisistent recipient
  - Failed send on account of wrong sender password
- Read Messages:
  - Successful read with username
  - Failed read on account of non exisistent user
  - Failed read on account of wrong password
- New User:
  - Successful creation of new user and the creator
  - Failed user creation on account of wrong password for creator

## Notes
The update that I made was to add an authentication step to send a message. There is now a new flag for the -send flag to include a sender. Usage: `go run main.go -send -to <user1> -from <user2>`. It should be noted that this does not vastly increase the security of this application since the database is still in plaintext.

## MAC Strategy
When a user sends a message, they must authenticate (see Notes) and a MAC is added to the database with the message. This MAC is checked when a user reads their messages and alerts the user if the MAC does not match. Moreover, the sender name is also reported when a user reads their messages. This new functionality was verified by modifying the databse with additional SQL statements that 1) altered the MAC and 2) altered the message to ensure that changes were logged and users were alerted. Moreover, the application does not allow for a user to update a MAC or message data.
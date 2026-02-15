package main

import (
	"fmt"
	"time"
)

type MessageType string

const (
	Chat   MessageType = "chat"
	Join   MessageType = "join"
	Leave  MessageType = "leave"
	Typing MessageType = "typing"
)

type Message struct {
	Type      MessageType
	RoomID    string
	UserID    string
	Username  string
	Content   string
	Timestamp time.Time
}

// constructor to create a new message
func NewMessage(t MessageType, RoomID, UserID, Username, Content string) Message {
	return Message{
		Type:      t,
		RoomID:    RoomID,
		UserID:    UserID,
		Username:  Username,
		Content:   Content,
		Timestamp: time.Now(),
	}
}

// check if the message is valid and has all the required fields
func Validate(msg Message) error {
	//validating message type
	if !msg.Type.IsValid() {
		return fmt.Errorf("Invalid message type %s", msg.Type)
	}
	//validating message data
	if msg.RoomID == "" || msg.UserID == "" || msg.Username == "" || msg.Content == "" {
		return fmt.Errorf("all field are required")
	}
	return nil
}

// Convert a message to a json format
func TOJSON() {}

// parse a JSON into a message struct
func FROMJSON() {}

// check if the message type is a valid entry
func (t MessageType) IsValid() bool {
	switch t {
	case Chat, Join, Leave, Typing:
		return true
	}
	return false
}

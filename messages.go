package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type MessageStatus int

const (
	ActiveMessageStatus MessageStatus = iota
	TrashedMessageStatus
)

func (m MessageStatus) String() string {
	return [...]string{
		"active",
		"trashed",
	}[m]
}

func (m *MessageStatus) FromString(s string) MessageStatus {
	return map[string]MessageStatus{
		"active":  ActiveMessageStatus,
		"trashed": TrashedMessageStatus,
	}[s]
}

func (m *MessageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *MessageStatus) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*m = m.FromString(s)
	return nil
}

func getMessages() ([]Message, error) {
	var msg []Message
	f, err := os.Open("data/messages.json")
	if err != nil {
		return nil, fmt.Errorf("error in opening file: %s", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error with reading: %s", err)
	}

	if err := json.Unmarshal(b, &msg); err != nil {
		return nil, fmt.Errorf("error unmarshaling: %s", err)
	}
	return msg, nil
}

type Message struct {
	Id      int
	Status  MessageStatus
	Title   string
	Content string
	Subject string
}

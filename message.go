package main

import "github.com/davilag/telego/kind"

// Returns the message kind based on the parameters that it has
func (m *Message) GetMessageKind() kind.MessageType {
	if m.Location != nil {
		return kind.Location
	}
	if m.Contact != nil {
		return kind.Contact
	}
	if m.Photo != nil {
		return kind.Photo
	}
	if m.Sticker != nil {
		return kind.Sticker
	}
	if m.Video != nil {
		return kind.Video
	}

	if m.Entities != nil {
		if m.getBotCommandEntity() != nil {
			return kind.Command
		}
	}
	return kind.Text
}

// Returns the Message Entity which type is bot_command
func (m *Message) getBotCommandEntity() *MessageEntity {
	if m.Entities == nil {
		return nil
	}

	for _, e := range *m.Entities {
		if e.Type == "bot_command" {
			return &e
		}
	}

	return nil
}

// From the Message.Text, it retrieves the command
func (m *Message) GetCommand() string {
	me := m.getBotCommandEntity()

	if me == nil {
		return ""
	}

	runes := []rune(m.Text)
	return string(runes[me.Offset+1 : me.Offset+me.Length])
}

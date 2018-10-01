// Package actions provides definitions of actions and their data
package actions

import (
	"encoding/json"
)

// Action represents an operation on a Dnote resource
type Action struct {
	UUID      string          `json:"uuid"`
	Schema    int             `json:"schema"`
	Type      string          `json:"type"`
	Data      json.RawMessage `json:"data"`
	Timestamp int64           `json:"timestamp"`
}

var (
	// ActionAddNote identifies a type of action for adding a note
	ActionAddNote = "add_note"
	// ActionRemoveNote  identifies a type of action for removing a note
	ActionRemoveNote = "remove_note"
	// ActionEditNote identifies a type of action for editing a note
	ActionEditNote = "edit_note"
	// ActionAddBook identifies a type of action for adding a book
	ActionAddBook = "add_book"
	// ActionRemoveBook identifies a type of action for removing a book
	ActionRemoveBook = "remove_book"
)

// NoteSnapshot is a snapshot of a note
type NoteSnapshot struct {
	Content  string `json:"content"`
	BookName string `json:"book_name"`
	Public   bool   `json:"public"`
}

// AddNoteDataV1 is a data for adding a note (v1)
type AddNoteDataV1 struct {
	NoteUUID string `json:"note_uuid"`
	BookName string `json:"book_name"`
	Content  string `json:"content"`
}

// AddNoteDataV2 is a data for adding a note (v2)
type AddNoteDataV2 struct {
	NoteUUID string `json:"note_uuid"`
	BookName string `json:"book_name"`
	Content  string `json:"content"`
	Public   bool   `json:"public"`
}

// EditNoteDataV1 is a data for editing a note (v1)
type EditNoteDataV1 struct {
	NoteUUID string `json:"note_uuid"`
	FromBook string `json:"from_book"`
	ToBook   string `json:"to_book"`
	Content  string `json:"content"`
}

// EditNoteDataV2 is a data for editing a note (v2)
type EditNoteDataV2 struct {
	NoteUUID string  `json:"note_uuid"`
	FromBook string  `json:"from_book"`
	ToBook   *string `json:"to_book"`
	Content  *string `json:"content"`
	Public   *bool   `json:"public"`
}

// EditNoteDataV3 is a data for editing a note (v2)
type EditNoteDataV3 struct {
	NoteUUID string       `json:"note_uuid"`
	Snapshot NoteSnapshot `json:"snapshot"`
	BookName *string      `json:"book_name"`
	Content  *string      `json:"content"`
	Public   *bool        `json:"public"`
}

// RemoveNoteDataV1 is a data for removing a note (v1)
type RemoveNoteDataV1 struct {
	NoteUUID string `json:"note_uuid"`
	BookName string `json:"book_name"`
}

// RemoveNoteDataV2 is a data for removing a note (v2)
type RemoveNoteDataV2 struct {
	NoteUUID string       `json:"note_uuid"`
	Snapshot NoteSnapshot `json:"snapshot"`
}

// AddBookDataV1 is a data for adding a book (v1)
type AddBookDataV1 struct {
	BookName string `json:"book_name"`
}

// RemoveBookDataV1 is a data for removing a book (v1)
type RemoveBookDataV1 struct {
	BookName string `json:"book_name"`
}

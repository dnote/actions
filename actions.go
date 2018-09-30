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

/********* add_note **/

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

// AddNoteDataV3 is a data for adding a note (v3)
type AddNoteDataV3 struct {
	NoteUUID string `json:"note_uuid"`
	BookUUID string `json:"book_uuid"`
	Content  string `json:"content"`
	Public   bool   `json:"public"`
}

/********* edit_note **/

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

// EditNoteDataV3 is a data for editing a note (v3)
type EditNoteDataV3 struct {
	NoteUUID string  `json:"note_uuid"`
	BookUUID *string `json:"book_uuid"`
	Content  *string `json:"content"`
	Public   *bool   `json:"public"`
}

/********* remove_note **/

// RemoveNoteDataV1 is a data for removing a note (v1)
type RemoveNoteDataV1 struct {
	NoteUUID string `json:"note_uuid"`
	BookName string `json:"book_name"`
}

// RemoveNoteDataV2 is a data for removing a note (v2)
type RemoveNoteDataV2 struct {
	NoteUUID string `json:"note_uuid"`
}

/********* add_book **/

// AddBookDataV1 is a data for adding a book (v1)
type AddBookDataV1 struct {
	BookName string `json:"book_name"`
}

// AddBookDataV2 is a data for adding a book (v2)
type AddBookDataV2 struct {
	BookName string `json:"book_name"`
	BookUUID string `json:"book_uuid"`
}

/********* remove_book **/

// RemoveBookDataV1 is a data for removing a book (v1)
type RemoveBookDataV1 struct {
	BookName string `json:"book_name"`
}

// RemoveBookDataV2 is a data for removing a book (v2)
type RemoveBookDataV2 struct {
	BookUUID string `json:"book_uuid"`
}

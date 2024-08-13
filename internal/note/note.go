package note

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/notnmeyer/dngl/internal/db"

	"github.com/google/uuid"
)

type Note struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type CreateNoteResponse struct {
	ID string `json:"id"`
}

type CreateNoteInput struct {
	Content string `json:"content"`
}

type GetNoteInput CreateNoteResponse

func New(body *io.ReadCloser) (*Note, error) {
	var n Note
	err := json.NewDecoder(*body).Decode(&n)

	if err != nil {
		return nil, err
	}

	n.ID = uuid.New().String()

	return &n, nil
}

func (n *Note) Save() error {
	client := db.New()
	if err := client.Save(n.ID, n.Content); err != nil {
		return err
	}

	return nil
}

func Get(id string) (*Note, error) {
	client := db.New()
	content, err := client.Get(id)
	if err != nil {
		return nil, err
	}

	return &Note{
		ID:      id,
		Content: *content,
	}, nil
}

func Delete(id string) error {
	client := db.New()
	if err := client.Delete(id); err != nil {
		return err
	}

	return nil
}

func List() ([]*Note, error) {
	client := db.New()
	result, err := client.GetAll()
	if err != nil {
		return nil, err
	}

	var notes []*Note
	for _, id := range result {
		n, err := Get(id)
		if err != nil {
			fmt.Printf("error getting id %s: %s\n", id, err.Error())
			continue
		}
		fmt.Printf("%+v", n.Content)
		notes = append(notes, n)
	}

	return notes, err
}

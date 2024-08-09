package note

import (
	"encoding/json"
	"io"

	"github.com/notnmeyer/dngl/internal/db"

	"github.com/google/uuid"
)

type Note struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
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

	n.ID = uuid.New()

	return &n, nil
}

func (n *Note) Save() error {
	client := db.New()
	if err := client.Save(n.ID.String(), n.Content); err != nil {
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

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &Note{
		ID:      parsedID,
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
	for k, v := range result {
		parsedID, err := uuid.Parse(k)
		if err != nil {
			continue
		}
		notes = append(notes, &Note{
			ID:      parsedID,
			Content: v,
		})
	}

	return notes, err
}

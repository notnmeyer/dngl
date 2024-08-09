package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/notnmeyer/dngl/internal/note"
)

func TestCreateNoteHandler(t *testing.T) {
	payloadBytes, err := json.Marshal(&note.Note{
		Content: "blah",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "{\"id\":\"heyo!\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

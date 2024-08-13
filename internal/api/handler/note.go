package handler

import (
	"encoding/json"
	"net/http"

	"github.com/notnmeyer/dngl/internal/note"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// TODO: move this out of the handler
	n, err := note.New(&r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := n.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(note.CreateNoteResponse{
		ID: n.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")
	n, err := note.Get(id)
	if err != nil {
		http.Error(w, "could not retrieve note", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(n); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")
	if err := note.Delete(id); err != nil {
		http.Error(w, "could not delete note: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	n, err := note.List()
	if err != nil {
		http.Error(w, "could not retrieve notes", http.StatusNotFound)
		return
	}

	if len(n) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

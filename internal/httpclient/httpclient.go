package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/notnmeyer/dngl/internal/envhelper"
	"github.com/notnmeyer/dngl/internal/note"
)

func NewRequest(verb, path string, data *string) ([]byte, error) {
	var (
		req     *http.Request
		env     = envhelper.New()
		reqPath = env.DNGL_API_URL + path
	)

	if data == nil {
		var err error
		req, err = http.NewRequest(verb, reqPath, nil)
		if err != nil {
			return nil, err
		}
	} else {
		jsonInput, err := json.Marshal(&note.CreateNoteInput{
			Content: *data,
		})
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(verb, reqPath, bytes.NewBuffer(jsonInput))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", env.DNGL_TOKEN))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

package grabjson

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJson(url string, model interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot get url: %w", err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error in reading response body: %w", err)
	}

	if err = json.Unmarshal(data, model); err != nil {
		return fmt.Errorf("internal error in parsing json: %w", err)
	}

	return nil
}

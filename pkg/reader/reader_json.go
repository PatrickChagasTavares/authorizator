package reader

import (
	"encoding/json"
	"io"
)

// ReadJSON reads data from an io.Reader and decodes it to the provided data structure.
// Receives an io.Reader containing JSON data and the data structure (pointer) to store
// the decoded data. It uses the encoding/json package to perform decoding.
// Returns an error if reading or decoding fails.
func ReadJSON(r io.Reader, data any) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(data)
}

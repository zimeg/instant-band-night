package terminal

import (
	"encoding/json"
	"os"

	"github.com/zimeg/instant-band-night/internal/errors"
)

// ReadFile reads a file at the provided path
func ReadFile(path string) ([]byte, error) {
	n, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return n, errors.ErrMissingFile
	}
	return n, err
}

// ReadJSON reads JSON from filepath into struct reference v
func ReadJSON(filepath string, v any) error {
	content, err := ReadFile(filepath)
	if err != nil {
		if err != errors.ErrMissingFile {
			return errors.ToIBNError(err)
		}
	} else if err = json.Unmarshal(content, v); err != nil {
		return errors.ToIBNError(err)
	}
	return nil
}

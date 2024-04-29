package terminal

import (
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

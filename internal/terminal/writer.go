package terminal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// WriteJSON saves indented JSON to a filename
func WriteJSON(filename string, v any) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0o777)
	if err != nil {
		return err
	}
	marshal, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, marshal, 0o666)
}

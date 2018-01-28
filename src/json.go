package src

import (
	"encoding/json"
	"io"
)

func EmitJson(w io.Writer, hs Highlights) error {
	out, err := json.MarshalIndent(hs, "", "  ")

	if err != nil {
		return err
	}

	w.Write(out)

	return nil
}

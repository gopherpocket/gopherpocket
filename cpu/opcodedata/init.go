package opcodedata

import (
	"bytes"
	_ "embed"
	"encoding/json"
)

//go:embed opcodes.json
var opcodesJSON []byte

func init() {
	decoder := json.NewDecoder(bytes.NewReader(opcodesJSON))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&OpcodeData); err != nil {
		panic(err)
	}
}

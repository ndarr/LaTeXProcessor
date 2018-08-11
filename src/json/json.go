package json

import (
	"encoding/json"
	"data"
)

func Decode(raw string) data.Repository{
	var m data.Message
	json.Unmarshal([]byte(raw), &m)
	return m.Repository
}
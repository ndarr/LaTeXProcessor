package json

import "encoding/json"

func Decode(raw string) Repository{
	var m message
	json.Unmarshal([]byte(raw), &m)
	return m.Repository
}
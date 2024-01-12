package helper

import (
	"encoding/json"
)

func Json(v any) string {
	bs, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bs)
}

func JsonIndent(v any) string {
	bs, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(bs)
}

package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type competition struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SignupFields string `json:"signupFields"`
	SignupLimit  int    `json:"signup_limit"`
	Remark       string
}

var fields = []string{"ID", "Name", "SignupFields", "SignupLimit"} // 大写

func TestStructToMap(t *testing.T) {
	src := &competition{1, "blueStar", "name,mobile", 10, "demo"}
	want := map[string]any{"id": 1, "name": "blueStar", "signupFields": "name,mobile", "signup_limit": 10}
	is := assert.New(t)
	is.Equal(want, StructToMap(src, fields))
}

func TestStructToMapExcept(t *testing.T) {
	src := &competition{1, "blueStar", "name,mobile", 10, "demo"}
	want := map[string]any{"id": 1, "name": "blueStar", "signupFields": "name,mobile", "Remark": "demo"}
	is := assert.New(t)
	is.Equal(want, StructToMapExcept(src, "SignupLimit"))
}

func TestStructsToMaps(t *testing.T) {
	src := []*competition{
		{1, "blueStar", "name,mobile", 10, "demo"},
	}

	want := []map[string]any{
		{"id": 1, "name": "blueStar", "signupFields": "name,mobile", "signup_limit": 10},
	}

	is := assert.New(t)
	is.Equal(want, StructsToMaps(src, fields))
}

//func BenchmarkStructsToMaps(b *testing.B) {
//	src := lo.Times(100, func(i int) *competition {
//		return &competition{i, "blueStar", "name,mobile", 10, "demo"}
//	})
//	for i := 0; i < b.N; i++ {
//		StructsToMaps(src, fields)
//	}
//}

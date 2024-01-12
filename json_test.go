package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type person struct {
	Name string
	Age  int
	Good bool
}

func TestJson(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"int", args{999}, `999`},
		{"string", args{"zcgly"}, `"zcgly"`},
		{"bool", args{true}, `true`},
		{"object", args{person{"zcg", 48, true}}, `{"Name":"zcg","Age":48,"Good":true}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Json(tt.args.v), "Json(%v)", tt.args.v)
		})
	}
}

func TestJsonIndent(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"int", args{999}, `999`},
		{"string", args{"zcgly"}, `"zcgly"`},
		{"bool", args{true}, `true`},
		{"object", args{person{"zcg", 48, true}}, `{
	"Name": "zcg",
	"Age": 48,
	"Good": true
}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, JsonIndent(tt.args.v), "Json(%v)", tt.args.v)
		})
	}
}

func TestDump(t *testing.T) {
	Dump("")
}

func TestDumpCompact(t *testing.T) {
	DumpCompact("")
}

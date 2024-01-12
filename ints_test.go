package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpperTo2Power(t *testing.T) {
	type upperToPowerCase struct {
		value int
		want  int
	}
	cases := []upperToPowerCase{
		{value: 0, want: 0},
		{value: 1, want: 1},
		{value: 2, want: 2},
		{value: 3, want: 4},
		{value: 4, want: 4},
		{value: 5, want: 8},
		{value: 6, want: 8},
		{value: 9, want: 16},
		{value: 12, want: 16},
		{value: 16, want: 16},
		{value: 17, want: 32},
	}
	for _, c := range cases {
		if UpperTo2Power(c.value) != c.want {
			t.Errorf("value: %d, get %d, want %d\n", c.value, UpperTo2Power(c.value), c.want)
		}
	}
}

func TestUpperTo2Power1(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{num: 0}, 0},
		{"1", args{num: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UpperTo2Power(tt.args.num), "UpperTo2Power(%v)", tt.args.num)
		})
	}
}

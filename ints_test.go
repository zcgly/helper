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

func TestStrToInts(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{s: "", sep: ","}, []int{}},
		{"1", args{s: "1", sep: ","}, []int{1}},
		{"2", args{s: "2,3", sep: ","}, []int{2, 3}},
		{"有空格", args{s: "2, 3", sep: ","}, []int{2, 3}},
		{"多余的逗号", args{s: ",2,3,", sep: ","}, []int{2, 3}},
		//{"无效", args{s: "2a,3", sep: ","}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StrToInts(tt.args.s, tt.args.sep), "StrToInts(%v, %v)", tt.args.s, tt.args.sep)
		})
	}
}

func TestIs2Power(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{num: 0}, false},
		{"1", args{num: 1}, true},
		{"2", args{num: 2}, true},
		{"3", args{num: 3}, false},
		{"4", args{num: 4}, true},
		{"5", args{num: 5}, false},
		{"6", args{num: 6}, false},
		{"7", args{num: 7}, false},
		{"8", args{num: 8}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Is2Power(tt.args.num), "Is2Power(%v)", tt.args.num)
		})
	}
}

func TestMoreThanHalf(t *testing.T) {
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
		{"2", args{num: 2}, 2},
		{"3", args{num: 3}, 2},
		{"4", args{num: 4}, 3},
		{"5", args{num: 5}, 3},
		{"6", args{num: 6}, 4},
		{"7", args{num: 7}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MoreThanHalf(tt.args.num), "MoreThanHalf(%v)", tt.args.num)
		})
	}
}

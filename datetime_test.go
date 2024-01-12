package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNextWeekTime(t *testing.T) {
	parseTime := func(s string) time.Time {
		t, _ := time.Parse("2006-01-02", s)
		return t
	}
	type args struct {
		t   time.Time
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"已经是当前日期后面", args{parseTime("2023-04-08"), parseTime("2023-04-03")}, parseTime("2023-04-08")},
		{"隔0天", args{parseTime("2023-04-01"), parseTime("2023-04-01")}, parseTime("2023-04-08")},
		{"隔1天", args{parseTime("2023-04-01"), parseTime("2023-04-02")}, parseTime("2023-04-08")},
		{"隔2天", args{parseTime("2023-04-01"), parseTime("2023-04-03")}, parseTime("2023-04-08")},
		{"隔6天", args{parseTime("2023-04-01"), parseTime("2023-04-07")}, parseTime("2023-04-08")},
		{"隔7天", args{parseTime("2023-04-01"), parseTime("2023-04-08")}, parseTime("2023-04-08")},
		{"隔8天", args{parseTime("2023-04-01"), parseTime("2023-04-09")}, parseTime("2023-04-15")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NextWeekTime(tt.args.t, tt.args.now), "NextWeekTime(%v)", tt.args.t)
		})
	}
}

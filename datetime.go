package helper

import "time"

func NextWeekTime(t time.Time, now time.Time) time.Time {
	if t.After(now) {
		return t
	}
	duration := now.Sub(t)
	days := int(duration.Hours() / 24)
	weeks := days / 7
	if days != 7 {
		weeks++
	}
	return t.AddDate(0, 0, weeks*7)
}

package helper

const MinFloatDiff = 0.00001

func IsFloat32Equal(f1, f2 float32) bool {
	if f1 > f2 {
		return f1-f2 < MinFloatDiff
	} else {
		return f2-f1 < MinFloatDiff
	}
}

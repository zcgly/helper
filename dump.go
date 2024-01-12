package helper

import "fmt"

func Dump(v ...any) {
	fmt.Println(JsonIndent(v))
}

func DumpCompact(v ...any) {
	fmt.Println(Json(v))
}

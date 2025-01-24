package io_P

import (
	"fmt"
	"os"

	gleam_P "example.com/todo/gleam"
)

func Print(s gleam_P.String_t) gleam_P.Nil_c {
	fmt.Print(s)
	return gleam_P.Nil_c{}
}

func Println(s gleam_P.String_t) gleam_P.Nil_c {
	fmt.Println(s)
	return gleam_P.Nil_c{}
}

func PrintError(s gleam_P.String_t) gleam_P.Nil_c {
	fmt.Fprint(os.Stderr, s)
	return gleam_P.Nil_c{}
}

func PrintlnError(s gleam_P.String_t) gleam_P.Nil_c {
	fmt.Fprintln(os.Stderr, s)
	return gleam_P.Nil_c{}
}

package stringʹ_P

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	gleam_P "example.com/todo/gleam"
	string_tree_P "example.com/todo/gleam_stdlib/gleam/string_tree"
)

func Length(x gleam_P.String_t) gleam_P.Int_t {
	return gleam_P.Int_t(len([]rune(x)))
}

func Lowercase(x gleam_P.String_t) gleam_P.String_t {
	return gleam_P.String_t(strings.ToLower(string(x)))
}

func Uppercase(x gleam_P.String_t) gleam_P.String_t {
	return gleam_P.String_t(strings.ToUpper(string(x)))
}

func LessThan(a, b gleam_P.String_t) gleam_P.Bool_t {
	return a < b
}

func doSlice(s gleam_P.String_t, start, length gleam_P.Int_t) gleam_P.String_t {
	if length <= 0 {
		return ""
	}

	runes := []rune(s)
	num_runes := gleam_P.Int_t(len(runes))

	if num_runes == 0 {
		return ""
	}

	for start < 0 {
		start += num_runes
	}

	end := start + length
	if end > num_runes {
		end = num_runes
	}

	return gleam_P.String_t(runes[start:end])
}

func Crop(s, substr gleam_P.String_t) gleam_P.String_t {
	i := strings.Index(string(s), string(substr))
	if i < 0 {
		return s
	}
	return s[i:]
}

func Contains(s, substr gleam_P.String_t) gleam_P.Bool_t {
	return gleam_P.Bool_t(strings.Contains(string(s), string(substr)))
}

func StartsWith(s, prefix gleam_P.String_t) gleam_P.Bool_t {
	return gleam_P.Bool_t(strings.HasPrefix(string(s), string(prefix)))
}

func EndsWith(s, suffix gleam_P.String_t) gleam_P.Bool_t {
	return gleam_P.Bool_t(strings.HasSuffix(string(s), string(suffix)))
}

func SplitOnce(s, sep gleam_P.String_t) gleam_P.Result_t[gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t], gleam_P.Nil_t] {
	parts := strings.SplitN(string(s), string(sep), 2)
	if len(parts) == 2 {
		return gleam_P.Ok_c[gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t], gleam_P.Nil_t]{P_0: gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t]{P_0: gleam_P.String_t(parts[0]), P_1: gleam_P.String_t(parts[1])}}
	}
	return gleam_P.Error_c[gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t], gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
}

func TrimStart(s gleam_P.String_t) gleam_P.String_t {
	return gleam_P.String_t(strings.TrimLeftFunc(string(s), unicode.IsSpace))
}

func TrimEnd(s gleam_P.String_t) gleam_P.String_t {
	return gleam_P.String_t(strings.TrimRightFunc(string(s), unicode.IsSpace))
}

func PopGrapheme(s gleam_P.String_t) gleam_P.Result_t[gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t], gleam_P.Nil_t] {
	r, size := utf8.DecodeRuneInString(string(s))
	if r == utf8.RuneError {
		return gleam_P.Error_c[gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t], gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t], gleam_P.Nil_t]{P_0: gleam_P.Tuple2_t[gleam_P.String_t, gleam_P.String_t]{P_0: gleam_P.String_t(r), P_1: s[size:]}}
}

func unsafeIntToUtfCodepoint(a gleam_P.Int_t) gleam_P.UtfCodepoint_t {
	return gleam_P.UtfCodepoint_t(a)
}

func doToUtfCodepoints(s gleam_P.String_t) gleam_P.List_t[gleam_P.UtfCodepoint_t] {
	return gleam_P.ToList([]gleam_P.UtfCodepoint_t(s)...)
}

func FromUtfCodepoints(xs gleam_P.List_t[gleam_P.UtfCodepoint_t]) gleam_P.String_t {
	var runes []rune
	for {
		switch xsC := xs.(type) {
		case gleam_P.Empty_c[gleam_P.UtfCodepoint_t]:
			return gleam_P.String_t(runes)
		case gleam_P.Nonempty_c[gleam_P.UtfCodepoint_t]:
			runes = append(runes, rune(xsC.P_0))
			xs = xsC.P_1
		default:
			panic(fmt.Sprintf("Invalid list %#v", xs))
		}
	}
}

func UtfCodepointToInt(cp gleam_P.UtfCodepoint_t) gleam_P.Int_t {
	return gleam_P.Int_t(cp)
}

func doInspect[T any](v T) string_tree_P.StringTree_t {
	return string_tree_P.FromString(gleam_P.String_t(fmt.Sprintf("%#v", v)))
}

func ByteSize(s gleam_P.String_t) gleam_P.Int_t {
	return gleam_P.Int_t(len(s))
}

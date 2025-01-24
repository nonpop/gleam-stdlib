package string_tree_P

import (
	"strings"

	gleam_P "example.com/todo/gleam"
	list_P "example.com/todo/gleam_stdlib/gleam/list"
)

type StringTree_t struct {
	value string
}

func (s StringTree_t) Hash() uint32 {
	return gleam_P.String_t(s.value).Hash()
}

func (s StringTree_t) Equal(o StringTree_t) bool {
	return s.value == o.value
}

func AppendTree(tree, suffix StringTree_t) StringTree_t {
	return StringTree_t{tree.value + suffix.value}
}

func FromStrings(strings gleam_P.List_t[gleam_P.String_t]) StringTree_t {
	return list_P.Fold(strings, StringTree_t{}, func(acc StringTree_t, s gleam_P.String_t) StringTree_t {
		return StringTree_t{string(acc.value) + string(s)}
	})
}

func Concat(trees gleam_P.List_t[StringTree_t]) StringTree_t {
	return list_P.Fold(trees, StringTree_t{}, AppendTree)
}

func FromString(s gleam_P.String_t) StringTree_t {
	return StringTree_t{string(s)}
}

func ToString(tree StringTree_t) gleam_P.String_t {
	return gleam_P.String_t(tree.value)
}

func ByteSize(tree StringTree_t) gleam_P.Int_t {
	return gleam_P.Int_t(len(tree.value))
}

func Lowercase(tree StringTree_t) StringTree_t {
	return StringTree_t{strings.ToLower(string(tree.value))}
}

func Uppercase(tree StringTree_t) StringTree_t {
	return StringTree_t{strings.ToUpper(string(tree.value))}
}

func doToGraphemes(s gleam_P.String_t) gleam_P.List_t[gleam_P.String_t] {
	graphemes := []gleam_P.String_t{}
	for _, r := range s {
		graphemes = append(graphemes, gleam_P.String_t(r))
	}
	return gleam_P.ToList(graphemes...)
}

func Split(tree StringTree_t, pattern gleam_P.String_t) gleam_P.List_t[StringTree_t] {
	segments := strings.Split(tree.value, string(pattern))
	elems := make([]gleam_P.String_t, len(segments))
	for i, s := range segments {
		elems[i] = gleam_P.String_t(s)
	}
	return list_P.Map(gleam_P.ToList(elems...), FromString)
}

func Replace(tree StringTree_t, each, with gleam_P.String_t) StringTree_t {
	return StringTree_t{strings.ReplaceAll(tree.value, string(each), string(with))}
}

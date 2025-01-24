package bit_array_P

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode/utf8"

	gleam_P "example.com/todo/gleam"
	list_P "example.com/todo/gleam_stdlib/gleam/list"
	order_P "example.com/todo/gleam_stdlib/gleam/order"
)

func FromString(s gleam_P.String_t) gleam_P.BitArray_t {
	return []byte(s)
}

func ByteSize(b gleam_P.BitArray_t) gleam_P.Int_t {
	return gleam_P.Int_t(len(b))
}

func Slice(bytes gleam_P.BitArray_t, position, length gleam_P.Int_t) gleam_P.Result_t[gleam_P.BitArray_t, gleam_P.Nil_t] {
	start := min(position, position+length)
	end := max(position, position+length)
	if start < 0 || end > gleam_P.Int_t(len(bytes)) {
		return gleam_P.Error_c[gleam_P.BitArray_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.BitArray_t, gleam_P.Nil_t]{P_0: bytes[start:end]}
}

func ToString(b gleam_P.BitArray_t) gleam_P.Result_t[gleam_P.String_t, gleam_P.Nil_t] {
	if !utf8.Valid(b) {
		return gleam_P.Error_c[gleam_P.String_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.String_t, gleam_P.Nil_t]{P_0: gleam_P.String_t(b)}
}

func concat2(x, y gleam_P.BitArray_t) gleam_P.BitArray_t {
	return append(x, y...)
}

func Concat(bs gleam_P.List_t[gleam_P.BitArray_t]) gleam_P.BitArray_t {
	return list_P.Fold(bs, gleam_P.BitArray_t{}, concat2)
}

func Base64Encode(b gleam_P.BitArray_t, _padding gleam_P.Bool_t) gleam_P.String_t {
	return gleam_P.String_t(base64.StdEncoding.EncodeToString(b))
}

func decode64(s gleam_P.String_t) gleam_P.Result_t[gleam_P.BitArray_t, gleam_P.Nil_t] {
	decoded, err := base64.StdEncoding.DecodeString(string(s))
	if err != nil {
		return gleam_P.Error_c[gleam_P.BitArray_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.BitArray_t, gleam_P.Nil_t]{P_0: decoded}
}

func Base16Encode(b gleam_P.BitArray_t) gleam_P.String_t {
	return gleam_P.String_t(strings.ToUpper(hex.EncodeToString(b)))
}

func Base16Decode(s gleam_P.String_t) gleam_P.Result_t[gleam_P.BitArray_t, gleam_P.Nil_t] {
	decoded, err := hex.DecodeString(string(s))
	if err != nil {
		return gleam_P.Error_c[gleam_P.BitArray_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.BitArray_t, gleam_P.Nil_t]{P_0: decoded}
}

func inspectLoop(b gleam_P.BitArray_t, acc gleam_P.String_t) gleam_P.String_t {
	return gleam_P.String_t(fmt.Sprintf("%s%v", acc, b))
}

func Compare(first, second gleam_P.BitArray_t) order_P.Order_t {
	firstLength := len(first)
	secondLength := len(second)
	for i := 0; i < firstLength; i++ {
		if i >= secondLength {
			return order_P.Gt_c{} // first has more items
		}
		f := first[i]
		s := second[i]
		if f > s {
			return order_P.Gt_c{}
		}
		if f < s {
			return order_P.Lt_c{}
		}
	}
	// This means that either first did not have any items
	// or all items in first were equal to second.
	if firstLength == secondLength {
		return order_P.Eq_c{}
	}
	return order_P.Lt_c{} // second has more items
}

func StartsWith(bits gleam_P.BitArray_t, prefix gleam_P.BitArray_t) gleam_P.Bool_t {
	if len(prefix) > len(bits) {
		return false
	}
	for i, b := range prefix {
		if bits[i] != b {
			return false
		}
	}
	return true
}

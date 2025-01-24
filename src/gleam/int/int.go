package intʹ_P

import (
	"strconv"

	gleam_P "example.com/todo/gleam"
)

func Parse(s gleam_P.String_t) gleam_P.Result_t[gleam_P.Int_t, gleam_P.Nil_t] {
	parsed, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return gleam_P.Error_c[gleam_P.Int_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.Int_t, gleam_P.Nil_t]{P_0: gleam_P.Int_t(parsed)}
}

func doBaseParse(s gleam_P.String_t, base gleam_P.Int_t) gleam_P.Result_t[gleam_P.Int_t, gleam_P.Nil_t] {
	parsed, err := strconv.ParseInt(string(s), int(base), 64)
	if err != nil {
		return gleam_P.Error_c[gleam_P.Int_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.Int_t, gleam_P.Nil_t]{P_0: gleam_P.Int_t(parsed)}
}

func ToString(i gleam_P.Int_t) gleam_P.String_t {
	return gleam_P.String_t(strconv.FormatInt(int64(i), 10))
}

func doToBaseString(i gleam_P.Int_t, base gleam_P.Int_t) gleam_P.String_t {
	return gleam_P.String_t(strconv.FormatInt(int64(i), int(base)))
}

func ToFloat(i gleam_P.Int_t) gleam_P.Float_t {
	return gleam_P.Float_t(i)
}

func BitwiseAnd(x, y gleam_P.Int_t) gleam_P.Int_t {
	return x & y
}

func BitwiseNot(x gleam_P.Int_t) gleam_P.Int_t {
	return ^x
}

func BitwiseOr(x, y gleam_P.Int_t) gleam_P.Int_t {
	return x | y
}

func BitwiseExclusiveOr(x, y gleam_P.Int_t) gleam_P.Int_t {
	return x ^ y
}

func BitwiseShiftLeft(x, y gleam_P.Int_t) gleam_P.Int_t {
	return x << y
}

func BitwiseShiftRight(x, y gleam_P.Int_t) gleam_P.Int_t {
	return x >> y
}

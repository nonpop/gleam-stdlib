package float_P

import (
	"math"
	"math/rand/v2"
	"strconv"

	gleam_P "example.com/todo/gleam"
)

func Parse(s gleam_P.String_t) gleam_P.Result_t[gleam_P.Float_t, gleam_P.Nil_t] {
	parsed, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		return gleam_P.Error_c[gleam_P.Float_t, gleam_P.Nil_t]{P_0: gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[gleam_P.Float_t, gleam_P.Nil_t]{P_0: gleam_P.Float_t(parsed)}
}

func ToString(f gleam_P.Float_t) gleam_P.String_t {
	return gleam_P.String_t(strconv.FormatFloat(float64(f), 'f', -1, 64))
}

func Ceiling(f gleam_P.Float_t) gleam_P.Float_t {
	return gleam_P.Float_t(math.Ceil(float64(f)))
}

func Floor(f gleam_P.Float_t) gleam_P.Float_t {
	return gleam_P.Float_t(math.Floor(float64(f)))
}

func Round(f gleam_P.Float_t) gleam_P.Int_t {
	return gleam_P.Int_t(math.Round(float64(f)))
}

func Truncate(f gleam_P.Float_t) gleam_P.Int_t {
	return gleam_P.Int_t(math.Trunc(float64(f)))
}

func doToFloat(i gleam_P.Int_t) gleam_P.Float_t {
	return gleam_P.Float_t(i)
}

func doPower(a, b gleam_P.Float_t) gleam_P.Float_t {
	return gleam_P.Float_t(math.Pow(float64(a), float64(b)))
}

func Random() gleam_P.Float_t {
	return gleam_P.Float_t(rand.Float64())
}

func doLog(a gleam_P.Float_t) gleam_P.Float_t {
	return gleam_P.Float_t(math.Log(float64(a)))
}

func Exponential(a gleam_P.Float_t) gleam_P.Float_t {
	return gleam_P.Float_t(math.Exp(float64(a)))
}

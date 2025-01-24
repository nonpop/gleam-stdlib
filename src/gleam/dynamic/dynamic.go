package dynamic_P

import (
	"fmt"

	gleam_P "example.com/todo/gleam"
	dict_P "example.com/todo/gleam_stdlib/gleam/dict"
	option_P "example.com/todo/gleam_stdlib/gleam/option"
)

type Dynamic_t = gleam_P.Dynamic_t

func From[T gleam_P.Type[T]](x T) Dynamic_t {
	return Dynamic_t{x}
}

type decodeErrors = gleam_P.List_t[DecodeError_t]

func decodeBitArray(x Dynamic_t) gleam_P.Result_t[gleam_P.BitArray_t, decodeErrors] {
	switch xv := x.Value.(type) {
	case gleam_P.BitArray_t:
		return gleam_P.Ok_c[gleam_P.BitArray_t, decodeErrors]{xv}
	default:
		return decoderError[gleam_P.BitArray_t]("BitArray", x)
	}
}

func decodeString(x Dynamic_t) gleam_P.Result_t[gleam_P.String_t, decodeErrors] {
	switch xv := x.Value.(type) {
	case gleam_P.String_t:
		return gleam_P.Ok_c[gleam_P.String_t, decodeErrors]{xv}
	default:
		return decoderError[gleam_P.String_t]("String", x)
	}
}

func decodeInt(x Dynamic_t) gleam_P.Result_t[gleam_P.Int_t, decodeErrors] {
	switch xv := x.Value.(type) {
	case gleam_P.Int_t:
		return gleam_P.Ok_c[gleam_P.Int_t, decodeErrors]{xv}
	default:
		return decoderError[gleam_P.Int_t]("Int", x)
	}
}

func decodeFloat(x Dynamic_t) gleam_P.Result_t[gleam_P.Float_t, decodeErrors] {
	switch xv := x.Value.(type) {
	case gleam_P.Float_t:
		return gleam_P.Ok_c[gleam_P.Float_t, decodeErrors]{xv}
	default:
		return decoderError[gleam_P.Float_t]("Float", x)
	}
}

func decodeBool(x Dynamic_t) gleam_P.Result_t[gleam_P.Bool_t, decodeErrors] {
	switch xv := x.Value.(type) {
	case gleam_P.Bool_t:
		return gleam_P.Ok_c[gleam_P.Bool_t, decodeErrors]{xv}
	default:
		return decoderError[gleam_P.Bool_t]("Bool", x)
	}
}

func decodeList(x Dynamic_t) gleam_P.Result_t[gleam_P.List_t[Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(gleam_P.List_dyn); ok {
		return gleam_P.Ok_c[gleam_P.List_t[Dynamic_t], decodeErrors]{x.ToDynamic()}
	}
	return decoderError[gleam_P.List_t[Dynamic_t]]("List", x)
}

func decodeResult(x Dynamic_t) gleam_P.Result_t[gleam_P.Result_t[Dynamic_t, Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(gleam_P.Result_dyn); ok {
		ok, err, isOk := x.ToDynamic()
		if isOk {
			return gleam_P.Ok_c[gleam_P.Result_t[Dynamic_t, Dynamic_t], decodeErrors]{
				gleam_P.Ok_c[Dynamic_t, Dynamic_t]{ok},
			}
		} else {
			return gleam_P.Ok_c[gleam_P.Result_t[Dynamic_t, Dynamic_t], decodeErrors]{
				gleam_P.Error_c[Dynamic_t, Dynamic_t]{err},
			}
		}
	}
	return decoderError[gleam_P.Result_t[Dynamic_t, Dynamic_t]]("Result", x)
}

func decodeOption(x Dynamic_t) gleam_P.Result_t[option_P.Option_t[Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(option_P.Option_dyn); ok {
		some, isSome := x.ToDynamic()
		if isSome {
			return gleam_P.Ok_c[option_P.Option_t[Dynamic_t], decodeErrors]{
				option_P.Some_c[Dynamic_t]{some},
			}
		} else {
			return gleam_P.Ok_c[option_P.Option_t[Dynamic_t], decodeErrors]{
				option_P.None_c[Dynamic_t]{},
			}
		}
	}
	return decoderError[option_P.Option_t[Dynamic_t]]("Option", x)
}

func decodeField[F gleam_P.Type[F]](x Dynamic_t, field F) gleam_P.Result_t[option_P.Option_t[Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(dict_P.Dict_dyn); ok {
		keys, values := x.ToDynamic()
		for i, k := range keys {
			if k.Equal(Dynamic_t{field}) {
				return gleam_P.Ok_c[option_P.Option_t[Dynamic_t], decodeErrors]{
					option_P.Some_c[Dynamic_t]{values[i]},
				}
			}
		}
		return gleam_P.Ok_c[option_P.Option_t[Dynamic_t], decodeErrors]{option_P.None_c[Dynamic_t]{}}
	}
	return decoderError[option_P.Option_t[Dynamic_t]]("Dict", x)
}

type unknownTuple_t []Dynamic_t

func (u unknownTuple_t) Hash() uint32 {
	h := gleam_P.NewOrderedCollectionHasher()
	for _, elem := range u {
		h.WriteHash(elem.Hash())
	}
	return h.Sum()
}

func (u unknownTuple_t) Equal(o unknownTuple_t) bool {
	if len(u) != len(o) {
		return false
	}
	for i := range u {
		if !u[i].Equal(o[i]) {
			return false
		}
	}
	return true
}

func decodeTuple(x Dynamic_t) gleam_P.Result_t[unknownTuple_t, decodeErrors] {
	if x, ok := x.Value.(gleam_P.Tuple_dyn); ok {
		return gleam_P.Ok_c[unknownTuple_t, decodeErrors]{x.ToDynamic()}
	}
	return decoderError[unknownTuple_t]("Tuple", x)
}

func decodeTuple2(x Dynamic_t) gleam_P.Result_t[gleam_P.Tuple2_t[Dynamic_t, Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(gleam_P.Tuple_dyn); ok {
		values := x.ToDynamic()
		if len(values) == 2 {
			return gleam_P.Ok_c[gleam_P.Tuple2_t[Dynamic_t, Dynamic_t], decodeErrors]{gleam_P.Tuple2_t[Dynamic_t, Dynamic_t]{values[0], values[1]}}
		}
	}
	return decoderError[gleam_P.Tuple2_t[Dynamic_t, Dynamic_t]]("Tuple of 2 elements", x)
}

func decodeTuple3(x Dynamic_t) gleam_P.Result_t[gleam_P.Tuple3_t[Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(gleam_P.Tuple_dyn); ok {
		values := x.ToDynamic()
		if len(values) == 3 {
			return gleam_P.Ok_c[gleam_P.Tuple3_t[Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors]{gleam_P.Tuple3_t[Dynamic_t, Dynamic_t, Dynamic_t]{values[0], values[1], values[2]}}
		}
	}
	return decoderError[gleam_P.Tuple3_t[Dynamic_t, Dynamic_t, Dynamic_t]]("Tuple of 3 elements", x)
}

func decodeTuple4(x Dynamic_t) gleam_P.Result_t[gleam_P.Tuple4_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(gleam_P.Tuple_dyn); ok {
		values := x.ToDynamic()
		if len(values) == 4 {
			return gleam_P.Ok_c[gleam_P.Tuple4_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors]{gleam_P.Tuple4_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t]{values[0], values[1], values[2], values[3]}}
		}
	}
	return decoderError[gleam_P.Tuple4_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t]]("Tuple of 4 elements", x)
}

func decodeTuple5(x Dynamic_t) gleam_P.Result_t[gleam_P.Tuple5_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(gleam_P.Tuple_dyn); ok {
		values := x.ToDynamic()
		if len(values) == 5 {
			return gleam_P.Ok_c[gleam_P.Tuple5_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors]{gleam_P.Tuple5_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t]{values[0], values[1], values[2], values[3], values[4]}}
		}
	}
	return decoderError[gleam_P.Tuple5_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t]]("Tuple of 5 elements", x)
}

func decodeTuple6(x Dynamic_t) gleam_P.Result_t[gleam_P.Tuple6_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(gleam_P.Tuple_dyn); ok {
		values := x.ToDynamic()
		if len(values) == 6 {
			return gleam_P.Ok_c[gleam_P.Tuple6_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t], decodeErrors]{gleam_P.Tuple6_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t]{values[0], values[1], values[2], values[3], values[4], values[5]}}
		}
	}
	return decoderError[gleam_P.Tuple6_t[Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t, Dynamic_t]]("Tuple of 6 elements", x)
}

func tupleGet(x unknownTuple_t, i gleam_P.Int_t) gleam_P.Result_t[Dynamic_t, decodeErrors] {
	if i < 0 || i >= gleam_P.Int_t(len(x)) {
		return decoderError[Dynamic_t]("Tuple index out of bounds", Dynamic_t{i})
	}
	return gleam_P.Ok_c[Dynamic_t, decodeErrors]{x[i]}
}

func tupleLength(x unknownTuple_t) gleam_P.Int_t {
	return gleam_P.Int_t(len(x))
}

func decodeDict(x Dynamic_t) gleam_P.Result_t[dict_P.Dict_t[Dynamic_t, Dynamic_t], decodeErrors] {
	if x, ok := x.Value.(dict_P.Dict_dyn); ok {
		keys, values := x.ToDynamic()
		entries := make([]gleam_P.Tuple2_t[Dynamic_t, Dynamic_t], len(keys))
		for i, k := range keys {
			entries[i] = gleam_P.Tuple2_t[Dynamic_t, Dynamic_t]{k, values[i]}
		}
		return gleam_P.Ok_c[dict_P.Dict_t[Dynamic_t, Dynamic_t], decodeErrors]{dict_P.FromList(gleam_P.ToList(entries...))}
	}
	return decoderError[dict_P.Dict_t[Dynamic_t, Dynamic_t]]("Dict", x)
}

func decoderError[T gleam_P.Type[T]](expected gleam_P.String_t, got Dynamic_t) gleam_P.Result_t[T, decodeErrors] {
	return decoderErrorNoClassify[T](expected, Classify(got))
}

func decoderErrorNoClassify[T gleam_P.Type[T]](expected gleam_P.String_t, got gleam_P.String_t) gleam_P.Result_t[T, decodeErrors] {
	return gleam_P.Error_c[T, decodeErrors]{gleam_P.ToList(DecodeError_t(DecodeError_c{
		gleam_P.String_t(expected),
		gleam_P.String_t(got),
		gleam_P.ToList[gleam_P.String_t](),
	}))}
}

func Classify(data Dynamic_t) gleam_P.String_t {
	switch data := data.Value.(type) {
	case gleam_P.String_t:
		return "String"
	case gleam_P.Bool_t:
		return "Bool"
	case gleam_P.Result_dyn:
		return "Result"
	case gleam_P.List_dyn:
		return "List"
	case gleam_P.BitArray_t:
		return "BitArray"
	case dict_P.Dict_dyn:
		return "Dict"
	case gleam_P.Int_t:
		return "Int"
	case gleam_P.Float_t:
		return "Float"
	case gleam_P.Tuple_dyn:
		{
			elems := len(data.ToDynamic())
			plural := ""
			if elems != 1 {
				plural = "s"
			}
			return gleam_P.String_t(fmt.Sprintf("Tuple of %d element%s", elems, plural))
		}
	case nil:
		return "Null"
	case gleam_P.Nil_t:
		return "Nil"
	default:
		return gleam_P.String_t(fmt.Sprintf("%T", data))
	}
}

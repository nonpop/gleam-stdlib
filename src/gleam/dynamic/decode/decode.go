package decode_P

import (
	"strconv"

	gleam_P "example.com/todo/gleam"
	dict_P "example.com/todo/gleam_stdlib/gleam/dict"
	dynamic_P "example.com/todo/gleam_stdlib/gleam/dynamic"
	option_P "example.com/todo/gleam_stdlib/gleam/option"
)

func strictIndex[T any](data any, key T) gleam_P.Result_t[option_P.Option_t[any], string] {
	if data, ok := data.(gleam_P.Indexable); ok {
		if value, ok := data.GetAt(key); ok {
			return gleam_P.Ok_c[option_P.Option_t[any], string]{option_P.Some_c[any]{value}}
		} else {
			return gleam_P.Ok_c[option_P.Option_t[any], string]{option_P.None_c[any]{}}
		}
	}
	return gleam_P.Error_c[option_P.Option_t[any], string]{"Indexable"}
}

func decodeList[T any](
	data any,
	item func(any) gleam_P.Tuple2_t[T, gleam_P.List_t[DecodeError_t]],
	pushPath func(gleam_P.Tuple2_t[T, gleam_P.List_t[DecodeError_t]], string) gleam_P.Tuple2_t[T, gleam_P.List_t[DecodeError_t]],
	index int64,
	acc gleam_P.List_t[T],
) gleam_P.Tuple2_t[gleam_P.List_t[T], gleam_P.List_t[DecodeError_t]] {
	dynList, ok := data.(gleam_P.List_dyn)
	if !ok {
		err := DecodeError_c{"List", dynamic_P.Classify(data), gleam_P.Empty_c[string]{}}
		return gleam_P.Tuple2_t[gleam_P.List_t[T], gleam_P.List_t[DecodeError_t]]{gleam_P.Empty_c[T]{}, gleam_P.ToList[DecodeError_t](err)}
	}
	var decoded []T
	items := dynList.ToDynamic()
	for {
		switch itemsC := items.(type) {
		case gleam_P.Empty_c[any]:
			return gleam_P.Tuple2_t[gleam_P.List_t[T], gleam_P.List_t[DecodeError_t]]{gleam_P.ToList[T](decoded...), gleam_P.Empty_c[DecodeError_t]{}}

		case gleam_P.Nonempty_c[any]:
			layer := item(itemsC.P_0)
			out := layer.P_0
			errors := layer.P_1

			if errors, ok := errors.(gleam_P.Nonempty_c[DecodeError_t]); ok {
				pushPath(layer, strconv.FormatInt(index, 10))
				return gleam_P.Tuple2_t[gleam_P.List_t[T], gleam_P.List_t[DecodeError_t]]{gleam_P.Empty_c[T]{}, errors}
			}

			decoded = append(decoded, out)
			index++
			items = itemsC.P_1
		}
	}
}

func decodeDict(x any) gleam_P.Result_t[dict_P.Dict_t[any, any], gleam_P.Nil_t] {
	if x, ok := x.(dict_P.Dict_dyn); ok {
		keys, values := x.ToDynamic()
		entries := make([]gleam_P.Tuple2_t[any, any], len(keys))
		for i, k := range keys {
			entries[i] = gleam_P.Tuple2_t[any, any]{k, values[i]}
		}
		return gleam_P.Ok_c[dict_P.Dict_t[any, any], gleam_P.Nil_t]{dict_P.FromList(gleam_P.ToList(entries...))}
	}
	return gleam_P.Error_c[dict_P.Dict_t[any, any], gleam_P.Nil_t]{gleam_P.Nil_c{}}
}

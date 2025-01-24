package dict_P

import (
	gleam_P "example.com/todo/gleam"
	"example.com/todo/gleam_stdlib/gleam/dict/immutable"
)

type Dict_t[K gleam_P.Type[K], V gleam_P.Type[V]] struct {
	m *immutable.Map[K, V]
}

type Dict_dyn interface {
	ToDynamic() ([]gleam_P.Dynamic_t, []gleam_P.Dynamic_t)
	gleam_P.Indexable
}

func (c Dict_t[K, V]) ToDynamic() ([]gleam_P.Dynamic_t, []gleam_P.Dynamic_t) {
	var keys []gleam_P.Dynamic_t
	var values []gleam_P.Dynamic_t

	iter := c.m.Iterator()
	for {
		k, v, ok := iter.Next()
		if !ok {
			return keys, values
		}
		keys = append(keys, gleam_P.Dynamic_t{k})
		values = append(values, gleam_P.Dynamic_t{v})
	}
}

func (c Dict_t[K, V]) GetAt(i any) (any, bool) {
	k, ok := i.(K)
	if !ok {
		return nil, false
	}
	v, ok := c.m.Get(k)
	if !ok {
		return nil, false
	}
	return v, true
}

func (c Dict_t[K, V]) Hash() uint32 {
	h := gleam_P.NewUnorderedCollectionHasher()
	iter := c.m.Iterator()
	for {
		k, v, ok := iter.Next()
		if !ok {
			return h.Sum()
		}
		h.WriteHash(gleam_P.HashTuple(k.Hash(), v.Hash()))
	}
}

func (c Dict_t[K, V]) Equal(o Dict_t[K, V]) bool {
	if c.m.Len() != o.m.Len() {
		return false
	}
	iter := c.m.Iterator()
	for {
		k, v1, ok := iter.Next()
		if !ok {
			return true
		}
		v2, ok := o.m.Get(k)
		if !ok || !v1.Equal(v2) {
			return false
		}
	}
}

func Size[K gleam_P.Type[K], V gleam_P.Type[V]](d Dict_t[K, V]) gleam_P.Int_t {
	return gleam_P.Int_t(d.m.Len())
}

func ToList[K gleam_P.Type[K], V gleam_P.Type[V]](dict Dict_t[K, V]) gleam_P.List_t[gleam_P.Tuple2_t[K, V]] {
	var entries gleam_P.List_t[gleam_P.Tuple2_t[K, V]] = gleam_P.Empty_c[gleam_P.Tuple2_t[K, V]]{}
	iter := dict.m.Iterator()
	for {
		k, v, ok := iter.Next()
		if !ok {
			return entries
		}
		entries = gleam_P.Nonempty_c[gleam_P.Tuple2_t[K, V]]{gleam_P.Tuple2_t[K, V]{k, v}, entries}
	}
}

func New[K gleam_P.Type[K], V gleam_P.Type[V]]() Dict_t[K, V] {
	return Dict_t[K, V]{immutable.NewMap[K, V]()}
}

func Get[K gleam_P.Type[K], V gleam_P.Type[V]](d Dict_t[K, V], k K) gleam_P.Result_t[V, gleam_P.Nil_t] {
	v, ok := d.m.Get(k)
	if !ok {
		return gleam_P.Error_c[V, gleam_P.Nil_t]{gleam_P.Nil_c{}}
	}
	return gleam_P.Ok_c[V, gleam_P.Nil_t]{v}
}

func doInsert[K gleam_P.Type[K], V gleam_P.Type[V]](k K, v V, d Dict_t[K, V]) Dict_t[K, V] {
	return Dict_t[K, V]{d.m.Set(k, v)}
}

func doDelete[K gleam_P.Type[K], V gleam_P.Type[V]](k K, d Dict_t[K, V]) Dict_t[K, V] {
	return Dict_t[K, V]{d.m.Delete(k)}
}

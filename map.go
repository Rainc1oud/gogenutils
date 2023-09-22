package gogenutils

// New collections implementations using generics

type Map[Tkey comparable, Telem any] map[Tkey]Telem

func TypedSlice[Telem any](s []interface{}) (ret []Telem) {
	ret = make([]Telem, len(s))
	for i, e := range s {
		ret[i] = e.(Telem)
	}
	return ret
}

func NewMap[Tkey comparable, Telem any]() Map[Tkey, Telem] {
	return make(Map[Tkey, Telem])
}

func (m Map[Tkey, Telem]) FilterMapByVal(cmpfun func(Telem) bool) (filteredMap Map[Tkey, Telem]) {
	filteredMap = make(Map[Tkey, Telem])
	for k, v := range m {
		if cmpfun(v) {
			filteredMap[k] = v
		}
	}
	return filteredMap
}

func (m Map[Tkey, Telem]) FilterMapByKey(cmpfun func(Tkey) bool) (filteredMap Map[Tkey, Telem]) {
	filteredMap = make(Map[Tkey, Telem])
	for k, v := range m {
		if cmpfun(k) {
			filteredMap[k] = v
		}
	}
	return filteredMap
}

// we could use MapToSlice, but that would mean two loops...
func (m Map[Tkey, Telem]) FilterMapToSlice(cmpfun func(Telem) bool) (filteredSlice []Telem) {
	filteredSlice = make([]Telem, 0, len(m))
	for _, v := range m {
		if cmpfun(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func (m *Map[Tkey, Telem]) ToSlice() (rslice []Telem) {
	rslice = make([]Telem, len(*m))
	i := 0
	for _, v := range *m {
		rslice[i] = v
		i++
	}
	return rslice
}

// AttrSlice is used to get a slice of attributes of Telem that are selected into a user-provided struct using the provided function
// Unfortunately we cannot use generics in the function parameter, so we are stuck with the any return type, which is the same as interface{}
// Therefore this function returns []interface{} which we must convert.
// As a workaround we now do this conversion generic static function TypedSlice.
// More info here? (But behind paywall) https://itnext.io/generic-map-filter-and-reduce-in-go-3845781a591c
func (m Map[Tkey, Telem]) AttrSlice(selfun func(Telem) any) (rslice []any) {
	rslice = make([]any, len(m))
	i := 0
	for _, v := range m {
		rslice[i] = selfun(v)
		i++
	}
	return rslice
}

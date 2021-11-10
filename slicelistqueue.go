package gogenutils

import "reflect"

type TSimpleSLQ struct {
	iterpos int
	items   []interface{}
}

// NewSimpleSLQ contsructor
// SimpleSLQ is a experimental, generic "Slice/List/Queue" hybrid like structure based on slice that implements:
//	Pop/Push, Next/Prev, IndexOf, ...
func NewSimpleSLQ() *TSimpleSLQ {
	ssq := new(TSimpleSLQ)
	ssq.items = make([]interface{}, 0)
	ssq.iterpos = -1
	return ssq
}

func (s *TSimpleSLQ) IndexOf(item interface{}) int {
	return PosInSlice(item, s.items)
}

func (s *TSimpleSLQ) Contains(item interface{}) bool {
	return PosInSlice(item, s.items) > -1
}

func (s *TSimpleSLQ) Length() int {
	return len(s.items)
}

func (s *TSimpleSLQ) Push(item interface{}) interface{} {
	s.items = append([]interface{}{item}, s.items...)
	return typedValue(s.items[0])
}

func (s *TSimpleSLQ) Pop() interface{} {
	if len(s.items) > 0 {
		r := s.items[0]
		s.items[0] = nil
		s.items = s.items[1:]
		return typedValue(r)
	}
	return nil
}

func (s *TSimpleSLQ) Append(item interface{}) interface{} {
	s.items = append(s.items, item)
	return typedValue(s.items[len(s.items)-1])
}

func (s *TSimpleSLQ) InsertAt(item interface{}, pos int) interface{} {
	if pos < 1 {
		return s.Push(item)
	}
	if pos > len(s.items)-1 {
		return typedValue(s.Append(item))
	}
	s.items = append(s.items[:pos], append([]interface{}{item}, s.items[pos:]...)...)
	return typedValue(s.items[pos])
}

func (s *TSimpleSLQ) Next() bool {
	if s.iterpos < len(s.items)-1 && s.iterpos > -1 {
		s.iterpos++
		return true
	}
	s.iterpos = -1 // reset iterator
	return false
}

func (s *TSimpleSLQ) Prev() bool {
	if s.iterpos > 0 && len(s.items) > 0 {
		s.iterpos--
		return true
	}
	s.iterpos = -1 // reset iterator
	return false
}

func (s *TSimpleSLQ) CurrentItem() interface{} {
	return s.ItemAt(s.iterpos)
}

func (s *TSimpleSLQ) ItemAt(pos int) interface{} {
	if pos > -1 && pos < len(s.items) {
		return typedValue(s.items[pos])
	}
	return nil
}

func typedValue(value interface{}) interface{} {
	// t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	if !v.IsValid() || v.IsZero() {
		return typedZero(v.Kind())
	}

	switch v.Kind() {
	case reflect.Int:
		return v.Interface().(int)
	case reflect.Float64:
		return v.Interface().(float64)
	case reflect.Float32:
		return v.Interface().(float32)
	case reflect.Bool:
		return v.Interface().(bool)
	case reflect.Struct:
		return v.Interface().(reflect.StructField)
	default:
		return v.Interface().(string)
	}
}

func typedZero(valkind reflect.Kind) interface{} {
	switch valkind {
	case reflect.Int:
		return *new(int)
	case reflect.Float64:
		return *new(float64)
	case reflect.Float32:
		return *new(float32)
	case reflect.Bool:
		return *new(bool)
	default:
		return *new(string)
	}
}

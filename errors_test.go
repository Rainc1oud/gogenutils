package gogenutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrors_Append(t *testing.T) {
	type fields struct {
		errs []error
	}
	type args struct {
		errs *Errors
	}

	e1 := NewErrors()
	for i := range []int{1, 2, 3} {
		e1.AddIf(fmt.Errorf("error%d", i))
	}

	e2 := NewErrors()
	for i := range []int{4, 5, 6} {
		e2.AddIf(fmt.Errorf("error%d", i))
	}

	ew := NewErrors()
	for _, e := range e1.errs {
		ew.AddIf(e)
	}
	for _, e := range e2.errs {
		ew.AddIf(e)
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Errors
	}{
		{
			name:   "3append3",
			fields: fields{errs: e1.errs},
			args:   args{errs: e2},
			want:   ew,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Errors{
				errs: tt.fields.errs,
			}
			if got := e.Append(tt.args.errs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errors.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

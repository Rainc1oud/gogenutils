package gogenutils

import (
	"fmt"
	"strings"
)

type Errors struct {
	errs []error
}

func NewErrors() *Errors {
	return &Errors{errs: make([]error, 0)}
}

func (e *Errors) AddIf(err error) error {
	if err != nil {
		e.errs = append(e.errs, err)
	}
	return err
}

func (e *Errors) Err() error {
	if len(e.errs) < 1 {
		return nil
	}
	return fmt.Errorf(e.ErrText())
}

func (e *Errors) ErrText() string {
	if len(e.errs) < 1 {
		return ""
	}
	es := make([]string, len(e.errs))
	for i, e := range e.errs {
		es[i] = e.Error()
	}
	return strings.Join(es, "\n")
}

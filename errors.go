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

func (e *Errors) AddMsgf(format string, param ...interface{}) error {
	err := fmt.Errorf(format, param...)
	e.AddIf(err)
	return err
}

func (e *Errors) AddMsgs(message ...string) {
	for _, s := range message {
		e.AddIf(fmt.Errorf(s))
	}
}

func (e *Errors) Err() error {
	if len(e.errs) < 1 {
		return nil
	}
	return fmt.Errorf(e.ErrString())
}

func (e *Errors) ErrStrings() []string {
	if len(e.errs) < 1 {
		return []string{}
	}
	es := make([]string, len(e.errs))
	for i, e := range e.errs {
		es[i] = e.Error()
	}
	return es
}

func (e *Errors) ErrString() string {
	if len(e.errs) < 1 {
		return ""
	}
	return strings.Join(e.ErrStrings(), "\n")
}

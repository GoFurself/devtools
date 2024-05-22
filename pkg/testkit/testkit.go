package testkit

import "testing"

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) True(value bool, msg string) {
	if !value {
		a.t.Errorf("expected true, got true. %s", msg)
	}
}

func (a *Assert) False(value bool, msg string) {
	if value {
		a.t.Errorf("expected false, got true. %s", msg)
	}
}

func (a *Assert) Equal(expected, actual any, msg string) {
	if expected != actual {
		a.t.Errorf("expected equal %v, got %v. %s", expected, actual, msg)
	}
}

func (a *Assert) NotEqual(expected, actual any, msg string) {
	if expected == actual {
		a.t.Errorf("expected not equal %v, got %v. %s", expected, actual, msg)
	}
}

func (a *Assert) Nil(value any, msg string) {
	if value != nil {
		a.t.Errorf("expected nil, got %v. %s.", value, msg)
	}
}

func (a *Assert) NotNil(value any, msg string) {
	if value == nil {
		a.t.Error("expected not nil, got nil", msg)
	}
}

func (a *Assert) Error(err error, msg string) {
	if err == nil {
		a.t.Error("expected error, got nil", msg)
	}
}

func (a *Assert) NoError(err error, msg string) {
	if err != nil {
		a.t.Errorf("expected no error, got %v. %s.", err, msg)
	}
}

func (a *Assert) Panic(f func(), msg string) {
	defer func() {
		if r := recover(); r == nil {
			a.t.Errorf("expected panic, got nil. %s.", msg)
		}
	}()
	f()
}

func (a *Assert) NoPanic(f func(), msg string) {

	defer func() {
		if r := recover(); r != nil {
			a.t.Errorf("expected no panic, got %v.  %s.", r, msg)
		}
	}()
	f()
}

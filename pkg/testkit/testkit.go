package testkit

import "testing"

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) True(value bool) {
	if !value {
		a.t.Error("expected true, got false")
	}
}

func (a *Assert) False(value bool) {
	if value {
		a.t.Error("expected false, got true")
	}
}

func (a *Assert) Equal(expected, actual any) {
	if expected != actual {
		a.t.Errorf("expected %v, got %v", expected, actual)
	}
}

func (a *Assert) NotEqual(expected, actual any) {
	if expected == actual {
		a.t.Errorf("expected %v, got %v", expected, actual)
	}
}

func (a *Assert) Nil(value any) {
	if value != nil {
		a.t.Errorf("expected nil, got %v", value)
	}
}

func (a *Assert) NotNil(value any) {
	if value == nil {
		a.t.Error("expected not nil, got nil")
	}
}

func (a *Assert) Error(err error) {
	if err == nil {
		a.t.Error("expected error, got nil")
	}
}

func (a *Assert) NoError(err error) {
	if err != nil {
		a.t.Errorf("expected no error, got %v", err)
	}
}

func (a *Assert) Panic(f func()) {
	defer func() {
		if r := recover(); r == nil {
			a.t.Error("expected panic, got nil")
		}
	}()
	f()
}

func (a *Assert) NoPanic(f func()) {

	defer func() {
		if r := recover(); r != nil {
			a.t.Errorf("expected no panic, got %v", r)
		}
	}()
	f()
}
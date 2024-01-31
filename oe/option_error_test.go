package oe

import (
	"errors"
	"testing"

	"github.com/808BiTT/options/option"
)

func TestSomeOE(t *testing.T) {
	oe := Some[int](5)
	if !oe.IsSome() {
		t.Errorf("Expected Some to contain a value")
	}
	if oe.Err != nil {
		t.Errorf("Expected Some to not contain an error")
	}
}

func TestNoneOE(t *testing.T) {
	oe := None[int]()
	if !oe.IsNone() {
		t.Errorf("Expected None to not contain a value")
	}
	if !errors.Is(oe.Err, option.ErrNoValue) {
		t.Errorf("Expected None to contain ErrNoValue error")
	}
}

func TestIsSomeOE(t *testing.T) {
	oe := Some[int](5)
	if !oe.IsSome() {
		t.Errorf("Expected IsSome to return true for Some")
	}
}

func TestIsNoneOE(t *testing.T) {
	oe := None[int]()
	if !oe.IsNone() {
		t.Errorf("Expected IsNone to return true for None")
	}
}

func TestGetOE(t *testing.T) {
	oe := Some[int](5)
	value, err := oe.Get()
	if err != nil {
		t.Errorf("Expected Get to not return an error for Some")
	}
	if value != 5 {
		t.Errorf("Expected Get to return the correct value for Some")
	}

	oe = None[int]()
	_, err = oe.Get()
	if err == nil {
		t.Errorf("Expected Get to return an error for None")
	}
}

package option

import (
	"testing"
)

func TestSome(t *testing.T) {
	opt := Some[int](5)
	if !opt.IsSome() {
		t.Errorf("Expected Some to contain a value")
	}
}

func TestNone(t *testing.T) {
	opt := None[int]()
	if !opt.IsNone() {
		t.Errorf("Expected None to not contain a value")
	}
}

func TestIsSome(t *testing.T) {
	opt := Some[int](5)
	if !opt.IsSome() {
		t.Errorf("Expected IsSome to return true for Some")
	}
}

func TestIsNone(t *testing.T) {
	opt := None[int]()
	if !opt.IsNone() {
		t.Errorf("Expected IsNone to return true for None")
	}
}

func TestGet(t *testing.T) {
	opt := Some[int](5)
	value, err := opt.Get()
	if err != nil {
		t.Errorf("Expected Get to not return an error for Some")
	}
	if value != 5 {
		t.Errorf("Expected Get to return the correct value for Some")
	}

	opt = None[int]()
	_, err = opt.Get()
	if err == nil {
		t.Errorf("Expected Get to return an error for None")
	}
}

func TestOrElse(t *testing.T) {
	opt := Some[int](5)
	value := opt.OrElse(10)
	if value != 5 {
		t.Errorf("Expected OrElse to return the value for Some")
	}

	opt = None[int]()
	value = opt.OrElse(10)
	if value != 10 {
		t.Errorf("Expected OrElse to return the default value for None")
	}
}

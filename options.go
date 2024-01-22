package options

import (
	"errors"
)

var ErrNoValue = errors.New("no value")

type Option[T any] struct {
	value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) Get() (T, error) {
	if o.IsNone() {
		var zero T
		return zero, ErrNoValue
	}
	return *o.value, nil
}

func (o Option[T]) OrElse(defaultValue T) T {
	if o.IsSome() {
		return *o.value
	}
	return defaultValue
}

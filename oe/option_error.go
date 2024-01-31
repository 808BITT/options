package oe

import (
	"github.com/808BiTT/options/option"
)

type Option[T any] struct {
	Option option.Option[T]
	Err    error
}

func Some[T any](value T) Option[T] {
	return Option[T]{Option: option.Some(value)}
}

func None[T any]() Option[T] {
	return Option[T]{Option: option.None[T](), Err: nil}
}

func Error[T any](err error) Option[T] {
	return Option[T]{Option: option.None[T](), Err: err}
}

func (o Option[T]) IsError() bool {
	return o.Err != nil
}

func (o Option[T]) IsSome() bool {
	return o.Option.IsSome()
}

func (o Option[T]) IsNone() bool {
	return o.Option.IsNone()
}

func (o Option[T]) Get() (T, error) {
	return o.Option.Get()
}

func (o Option[T]) OrElse(defaultValue T) T {
	return o.Option.OrElse(defaultValue)
}

func (o Option[T]) Error() error {
	return o.Err
}

package sample

import (
	"fmt"

	"github.com/808BiTT/options/oe"
	"github.com/808BiTT/options/option"
)

type SampleStruct struct {
	Title string
	Desc  string
	ID    int
}

var DefaultValue = SampleStruct{
	Title: "default title",
	Desc:  "default description",
	ID:    0,
}

func (s SampleStruct) String() string {
	return fmt.Sprintf("Title: %s, Description: %s, ID: %d", s.Title, s.Desc, s.ID)
}

func SimpleOptionFunc(valFlag, errFlag bool) (option.Option[SampleStruct], error) {
	if errFlag {
		return option.None[SampleStruct](), option.ErrNoValue
	}
	if valFlag {
		res := SampleStruct{
			Title: "title",
			Desc:  "details",
			ID:    1,
		}
		return option.Some(res), nil
	}
	return option.None[SampleStruct](), nil
}

func SimpleOptionErrorFunc(valFlag, errFlag bool) oe.Option[SampleStruct] {
	if errFlag {
		return oe.Error[SampleStruct](option.ErrNoValue)
	}
	if valFlag {
		res := SampleStruct{
			Title: "title",
			Desc:  "details",
			ID:    1,
		}
		return oe.Some(res)
	}
	return oe.None[SampleStruct]()
}

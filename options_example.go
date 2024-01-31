package main

import (
	"fmt"
	"os"

	"github.com/808BiTT/options/oe"
	"github.com/808BiTT/options/sample"
)

func main() {
	runOptionExamples()
	runOeExamples()

	text := readTextFile("em.txt")
	if text.IsError() {
		fmt.Println(text.Error())
	}
	fmt.Println(string(text.OrElse("File is Empty")))

}

func readTextFile(filepath string) oe.Option[string] {
	// If error return Error[string](err)
	// If success return Some[string](data)
	// If no data return None[string]()
	f, err := os.ReadFile(filepath)
	if err != nil {
		return oe.Error[string](err)
	}
	if len(f) == 0 {
		return oe.None[string]()
	}

	return oe.Some(string(f))
}

func runOptionExamples() {
	// Value is returned as expected
	fmt.Println("Expected: Title: title, Description: details, ID: 1")
	val, err := sample.SimpleOptionFunc(true, false)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		res := val.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()

	// Or else value is used as expected
	fmt.Println("Expected: Title: default title, Description: default description, ID: 0")
	val, err = sample.SimpleOptionFunc(false, false)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		res := val.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()

	// Error is returned as expected
	fmt.Println("Expected: no value")
	_, err = sample.SimpleOptionFunc(false, true)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		res := val.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()

	// Error is returned as expected
	fmt.Println("Expected: no value")
	val, err = sample.SimpleOptionFunc(true, true)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		res := val.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()
}

func runOeExamples() {
	// Value is returned as expected
	fmt.Println("Expected: Title: title, Description: details, ID: 1")
	oe := sample.SimpleOptionErrorFunc(true, false)
	if oe.IsError() {
		fmt.Println(oe.Error())
	} else {
		res := oe.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()

	// Or else value is used as expected
	fmt.Println("Expected: Title: default title, Description: default description, ID: 0")
	oe = sample.SimpleOptionErrorFunc(false, false)
	if oe.IsError() {
		fmt.Println(oe.Error())
	} else {
		res := oe.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()

	// Error is returned as expected
	fmt.Println("Expected: no value")
	oe = sample.SimpleOptionErrorFunc(false, true)
	if oe.IsError() {
		fmt.Println(oe.Error())
	} else {
		res := oe.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()

	// Error is returned as expected
	fmt.Println("Expected: no value")
	oe = sample.SimpleOptionErrorFunc(true, true)
	if oe.IsError() {
		fmt.Println(oe.Error())
	} else {
		res := oe.OrElse(sample.DefaultValue)
		fmt.Println(res)
	}
	fmt.Println()
}

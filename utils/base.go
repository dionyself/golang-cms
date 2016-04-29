package utils

import (
	_ "fmt"
)

/*
type Decor func(s string)

func NewDecor(fn Decor) Decor {
	return func(s string) {
		fn.extra(s)
	}
}

func (d Decor) Add(fn Decor) Decor {
	return func(s string) {
		d(s)
		fn(s)
	}
}

func (d Decor) extra(s string) {
	d(s)
	fmt.Printf("extra %s\n", s)
}

// --
var SomeFunc Decor = UndecoratedFunc

func WebHandler(w http.ResponseWriter, r *Request) {
	// Yay! Naive initialization!
	SomeFunc = SomeFunc.Add(SomeDecorator)
	SomeFunc("wohoo!")
}
*/

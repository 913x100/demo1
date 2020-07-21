package main

import (
	"github.com/913x100/demo1/external/outer"
	"github.com/913x100/demo1/internal/inner"
)

func main() {
	o := outer.Outer{}
	i := inner.Inner{}
	o.Print()
	i.Print()
}

package main

import (
	"github.com/alecthomas/kong"
	"github.com/naoyafurudono/opname"
)

var CLI struct {
	Prefix string `arg:""`
}

func main() {
	ctx := kong.Parse(&CLI)
	g, err := opname.New(CLI.Prefix)
	if err != nil {
		ctx.FatalIfErrorf(err)
	}
	println(g.Gen())
}

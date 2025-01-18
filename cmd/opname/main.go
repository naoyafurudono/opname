package main

import (
	"github.com/alecthomas/kong"
	"github.com/naoyafurudono/opname"
)

var CLI struct {
	Prefix string `arg:""`
	Num    int    `long:"num" short:"n" default:"1"`
}

func main() {
	ctx := kong.Parse(&CLI)
	g, err := opname.New(CLI.Prefix)
	if err != nil {
		ctx.FatalIfErrorf(err)
	}
	for range CLI.Num {
		println(g.Gen())
	}
}

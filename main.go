package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/dave/jennifer/jen"
)

const (
	GOMOOL_NAME = "gomool"
	GOMOOL_VERSION = "0.0.1"
)

func parseArgs() {
	var version bool
	flag.BoolVar(&version, "v", false, "Show version")
	flag.BoolVar(&version, "version", false, "Show version")

	flag.Parse()

	if version {
		fmt.Printf("%s v%s\n", GOMOOL_NAME, GOMOOL_VERSION)
	}

	var wg sync.WaitGroup
	for _, arg := range flag.Args() {
		wg.Add(1)
		go compile(&wg, arg)
	}
	wg.Wait()
}

func main() {
	parseArgs()

	f := jen.NewFile("main")

	fmt.Printf("%#v\n", f)
}

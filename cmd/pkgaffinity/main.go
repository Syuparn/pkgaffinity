package main

import (
	"github.com/syuparn/pkgaffinity"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(pkgaffinity.Analyzer) }

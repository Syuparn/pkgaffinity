package main

import (
	"github.com/syuparn/pkgaffinity"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(pkgaffinity.Analyzer) }

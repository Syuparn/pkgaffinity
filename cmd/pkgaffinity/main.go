package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/syuparn/pkgaffinity"
)

func main() { singlechecker.Main(pkgaffinity.Analyzer) }

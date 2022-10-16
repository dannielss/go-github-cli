package main

import (
	"flag"

	"github.com/dannielss/go-github-cli/cmd/outputs"
)

func main() {
	var h bool
	var e bool
	var u string
	var w int

	flag.StringVar(&u, "u", "", "")
	flag.StringVar(&u, "u", "", "")

	flag.BoolVar(&h, "h", false, "")
	flag.BoolVar(&h, "help", false, "")

	flag.BoolVar(&e, "e", false, "")
	flag.BoolVar(&e, "export", false, "")

	flag.IntVar(&w, "w", 1, "")
	flag.IntVar(&w, "w", 1, "")

	setFlag(flag.CommandLine)

	flag.Parse()

	if h {
		outputs.ShowHelp()
		return
	}

	outputs.GetHeader()

	if u != "" && !e {
		outputs.GetRepositoriesInfo(u)
		return
	}

	if u != "" && e {
		if w > 1 {
			outputs.ExportAsCSVConcurrently(u, w)
		} else {
			outputs.ExportAsCSV(u)
		}
	}
}

func setFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		outputs.ShowHelp()
	}
}

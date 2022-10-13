package main

import (
	"flag"
	"log"

	"github.com/dannielss/go-github-cli/cmd/outputs"
)

func main() {
	var h bool
	var e bool
	var user string

	flag.StringVar(&user, "u", "", "")
	flag.StringVar(&user, "user", "", "")

	flag.BoolVar(&h, "h", false, "")
	flag.BoolVar(&h, "help", false, "")

	flag.BoolVar(&e, "e", false, "")
	flag.BoolVar(&e, "export", false, "")

	setFlag(flag.CommandLine)

	flag.Parse()

	if h {
		outputs.ShowHelp()
		return
	}

	outputs.GetHeader()

	if user != "" {
		outputs.GetRepositoriesInfo(user)
		return
	}

	if e {
		log.Printf("Exporting as csv...")
	}
}

func setFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		outputs.ShowHelp()
	}
}

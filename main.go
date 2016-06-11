package main

import (
	"fmt"
	"os"

	"github.com/micahhausler/loco/archive"
	flag "github.com/spf13/pflag"
)

const Version = "0.0.1"

var version = flag.Bool("version", false, "print version and exit")

var registry = flag.StringP("registry", "r", "https://index.docker.io/v1/", "Specify a specific registry")
var username = flag.StringP("username", "u", "", "The user to login as")
var password = flag.StringP("password", "p", "", "The password ")
var outfile = flag.StringP("output", "o", "docker.tgz", "The file to create. Use \"-\" for Stdout.")

func main() {

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Docker Login Compressor\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Printf("loco %s\n", Version)
		os.Exit(0)
	}

	lc := archive.LoginConfig{
		Registry:   *registry,
		Username:   *username,
		Password:   *password,
		OutputFile: *outfile,
	}
	lc.CreateArchive()
}

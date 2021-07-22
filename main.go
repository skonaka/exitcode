package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func flagUsage() {
	usageText := `exitcode returns designated exit code
	Usage:
	exitcode [number]
	
	Mac, linux:
	$ ./exitcode 123
	$ echo $?
	123

	Windows:
	C:\ exitcode.ext 123
	C:\> echo %ERRORLEVEL%
	123 `

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func main() {
	flag.Usage = flagUsage

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	flag.Parse()
	ec, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}
	os.Exit(ec)
}

// Command csvstat profiles the columns of a CSV file.
package main

import (
	"os"

	"csvstat/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}

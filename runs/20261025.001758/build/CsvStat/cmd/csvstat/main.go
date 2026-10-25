// Command csvstat profiles one CSV file. Its behavior is defined by internal/cli.
package main

import (
	"os"

	"csvstat/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}

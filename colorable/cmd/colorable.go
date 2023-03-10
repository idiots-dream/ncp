package main

import (
	"io"
	"os"

	"ncp/colorable"
)

func main() {
	io.Copy(colorable.NewColorableStdout(), os.Stdin)
}

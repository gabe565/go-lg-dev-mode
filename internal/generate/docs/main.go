package main

import (
	"os"

	"gabe565.com/lg-dev-mode/cmd"
	"gabe565.com/utils/cobrax"
	"github.com/spf13/cobra/doc"
)

func main() {
	output := "./docs"

	if err := os.MkdirAll(output, 0o755); err != nil {
		panic(err)
	}

	root := cmd.New(cobrax.WithVersion("beta"))
	if err := doc.GenMarkdownTree(root, output); err != nil {
		panic(err)
	}
}

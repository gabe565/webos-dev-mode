package main

import (
	"net/http"
	"os"

	"gabe565.com/lg-dev-mode/cmd"
	"gabe565.com/utils/cobrax"
	"gabe565.com/utils/httpx"
)

var version = "beta"

func main() {
	root := cmd.New(cobrax.WithVersion(version))
	http.DefaultTransport = httpx.NewUserAgentTransport(nil, cobrax.BuildUserAgent(root))
	if err := root.Execute(); err != nil {
		root.PrintErrln(root.ErrPrefix(), err)
		os.Exit(1)
	}
}

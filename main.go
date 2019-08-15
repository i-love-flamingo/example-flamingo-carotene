package main

import (
	"flamingo.me/dingo"
	"flamingo.me/example-hello-flamingo-carotene/src/helloworld"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/flamingo/v3/core/zap"
	"flamingo.me/pugtemplate"
)

// main is our entry point
func main() {
	flamingo.App([]dingo.Module{
		new(zap.Module),           // log formatter
		new(requestlogger.Module), // requestlogger show request logs
		new(helloworld.Module),
		new(pugtemplate.Module),
	})
}

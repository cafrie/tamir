package main

import (
	"github.com/cafrie/tamir/api/server/http"
	"github.com/orchestd/dependencybundler/bundler"
	"github.com/orchestd/dependencybundler/depBundler"
)

type Configuration struct {
	depBundler.DependencyBundlerConfiguration
}

var appConf Configuration

func main() {
	bundler.CreateApplication(&appConf, http.InitHandlers, deps()...)
}

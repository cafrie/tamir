package main

import (
	"github.com/cafrie/tamir/api/server/monolith"
	"github.com/cafrie/tamir/application/defaultApp"
)

func deps() []interface{} {
	return append(internalDeps(), externalDeps()...)
}

func internalDeps() []interface{} {
	return []interface{}{
		defaultApp.NewTamirApp,
		monolith.NewTamirInterface,
	}
}

func externalDeps() []interface{} {
	return []interface{}{}
}

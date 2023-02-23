/* This file is auto-generated and should not be modified */

package http

import (
	"github.com/cafrie/tamir/api"
	"github.com/cafrie/tamir/api/server/monolith"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/transport"
)

func InitHandlers(router transport.IRouter, m monolith.TamirInterface, c configuration.Config) {
	router.POST(api.TestMethod, transport.HandleFunc(m.Test))
}

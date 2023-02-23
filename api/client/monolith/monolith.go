/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/cafrie/tamir/api"
	"github.com/cafrie/tamir/api/server/monolith"
	. "github.com/orchestd/servicereply"
)

type TamirMonolithClient struct {
	MonolithInterface monolith.TamirInterface
}

func NewTamirMonolithClient(monolithInterface monolith.TamirInterface) api.TamirApi {
	return TamirMonolithClient{MonolithInterface: monolithInterface}
}

func (p TamirMonolithClient) Test(c context.Context, req api.TestRequest) ServiceReply {
	return p.MonolithInterface.Test(c, req)
}

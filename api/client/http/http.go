/* This file is auto-generated and should not be modified */

package http

import (
	"context"
	"github.com/cafrie/tamir/api"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	. "github.com/orchestd/servicereply"
)

const serviceName = "tamir"

type tamirHTTPClient struct {
	client transport.HttpClient
}

func NewTamirApiHTTPClient(client transport.HttpClient) api.TamirApi {
	return tamirHTTPClient{client: client}
}

func (h tamirHTTPClient) Test(c context.Context, req api.TestRequest) ServiceReply {
	if reply := h.client.Call(c, req, serviceName, api.TestMethod, nil, nil); !reply.IsSuccess() {
		return reply
	}
	return nil
}

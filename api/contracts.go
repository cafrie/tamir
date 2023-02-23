/* This file is auto-generated and should not be modified */

package api

import (
	"context"
	. "github.com/orchestd/servicereply"
)

type TamirApi interface {
	Test(c context.Context, req TestRequest) ServiceReply
}

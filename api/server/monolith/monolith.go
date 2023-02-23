/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/cafrie/tamir/api"
	"github.com/cafrie/tamir/domain/application"
	"github.com/orchestd/dependencybundler/interfaces/validations"
	. "github.com/orchestd/servicereply"
)

type TamirInterface struct {
	tamirApp   application.TamirApp
	validation validations.Validations
}

func NewTamirInterface(tamirApp application.TamirApp, validation validations.Validations) TamirInterface {
	return TamirInterface{tamirApp: tamirApp, validation: validation}
}

func (i TamirInterface) Test(c context.Context, req api.TestRequest) ServiceReply {
	if vErr := i.validation.Validate(req); vErr != nil && !vErr.IsSuccess() {
		return vErr
	}
	err := i.tamirApp.Test(c)
	return err
}

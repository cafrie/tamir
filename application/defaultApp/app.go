package defaultApp

import (
	"github.com/cafrie/tamir/domain/application"
)

type tamirApp struct {
}

func NewTamirApp() application.TamirApp {
	return &tamirApp{}
}

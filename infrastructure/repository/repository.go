package repository

import (
	"github.com/cafrie/tamir/domain/repository"
)

type tamirRepository struct {
}

func NewTamirRepository() repository.TamirRepository {
	return tamirRepository{}
}

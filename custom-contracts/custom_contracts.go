package customcontracts

import (
	"github.com/google/uuid"
)

type CustomContracts struct {
	Name string
	ID   uuid.UUID
}

func NewCustomContracts(name string) CustomContracts {
	return CustomContracts{Name: name, ID: uuid.New()}
}

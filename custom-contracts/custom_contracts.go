package customcontracts

import (
	"time"

	"github.com/google/uuid"
)

type CustomContracts struct {
	Name      string
	ID        uuid.UUID
	ID2       uuid.UUID
	CreatedAt time.Time
}

func NewCustomContracts(name string) CustomContracts {
	return CustomContracts{Name: name, ID: uuid.New(), CreatedAt: time.Now()}
}

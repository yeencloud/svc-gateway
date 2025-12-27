package ports

import (
	"context"
)

type Usecases interface {
	// User
	RegisterUser(context context.Context) error
}

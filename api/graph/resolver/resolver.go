package resolver

//go:generate go run -mod=mod github.com/99designs/gqlgen generate
import (
	repo "github.com/PwrFr/gqlgen/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RepoDB repo.RepoDB
	// rooms   []*model.Room
	// account []*model.Account
}

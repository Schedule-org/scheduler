package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Account struct {
		Email string `json:"email"`
		Cnpj  string `json:"cnpj"`
	}

	AccountUseCase interface {
		// Add creates a new account.
		Add(ctx context.Context, payload *Account) (*Account, *core.Exception)
	}

	AccountRepository interface {
		// Add creates a new account.
		Add(ctx context.Context, payload *Account) (*Account, *core.Exception)
	}

	AccountController interface {
		// Add handles the HTTP request to create a new account.
		Add(ctx *gin.Context)
	}
)

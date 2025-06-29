package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	Account struct {
		Name        string       `json:"name"`
		Email       string       `json:"email"`
		Password    string       `json:"password"`
		ActivatedAt sql.NullTime `json:"activate_at"`
		CreatedAt   time.Time    `json:"created_at"`
		UpdatedAt   time.Time    `json:"updated_at"`
	}

	// AccountUseCase defines the business logic related to Accounts.
	AccountUseCase interface {
		// Add creates a new Account.
		Add(ctx context.Context, payload *Account) (*Account, *core.Exception)
		// FindAccountById retrieves a Account by their ID.
		FindAccountById(ctx context.Context, id string) (*Account, *core.Exception)
		// FindAllAccounts returns a list of all Accounts.
		FindAllAccounts(ctx context.Context) ([]Account, *core.Exception)
		// FindAllEstablishmentsByAccountId retrieves all establishments associated with a specific Account.
		FindAllEstablishmentsByAccountId(ctx context.Context, Account_id string) ([]Establishment, *core.Exception)
	}

	// AccountRepository defines the data access layer for Accounts.
	AccountRepository interface {
		// Add inserts a new Account into the database.
		Add(ctx context.Context, Account *Account) (*Account, error)
		// FindAccountByEmail retrieves a Account by their email address.
		FindAccountByEmail(ctx context.Context, email string) (*Account, error)
		// FindAccountById retrieves a Account by their ID from the database.
		FindAccountById(ctx context.Context, id string) (*Account, error)
		// FindAllEstablishmentsByAccountId fetches all establishments linked to a Account.
		FindAllEstablishmentsByAccountId(ctx context.Context, Account_id string) ([]Establishment, error)
		// FindAllAccounts retrieves all Accounts from the database.
		FindAllAccounts(ctx context.Context) ([]Account, error)
	}

	// AccountController defines the HTTP handlers for managing Accounts.
	AccountController interface {
		// Add handles the HTTP request to create a new Account.
		Add(ctx *gin.Context)
		// FindAccountById handles the HTTP request to retrieve a Account by ID.
		FindAccountById(ctx *gin.Context)
		// FindAllAccounts handles the HTTP request to return a list of all Accounts.
		FindAllAccounts(ctx *gin.Context)
		// FindAllEstablishmentsByAccountId handles the HTTP request to fetch all establishments linked to a Account.
		FindAllEstablishmentsByAccountId(ctx *gin.Context)
	}
)

package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

type (
	User struct {
		Name        string       `json:"name"`
		Email       string       `json:"email"`
		Cnpj        string       `gorm:"size:100;unique;not null"`
		Password    string       `json:"password"`
		ActivatedAt sql.NullTime `json:"activate_at"`
		CreatedAt   time.Time    `json:"created_at"`
		UpdatedAt   time.Time    `json:"updated_at"`
	}

	// UserUseCase defines the business logic related to users.
	UserUseCase interface {
		// Add creates a new user.
		Add(ctx context.Context, payload *User) (*User, *core.Exception)
		// FindUserById retrieves a user by their ID.
		FindUserById(ctx context.Context, id string) (*User, *core.Exception)
		// FindAllUsers returns a list of all users.
		FindAllUsers(ctx context.Context) ([]User, *core.Exception)
		// FindAllEstablishmentsByUserId retrieves all establishments associated with a specific user.
		FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]Establishment, *core.Exception)
	}

	// UserRepository defines the data access layer for users.
	UserRepository interface {
		// Add inserts a new user into the database.
		Add(ctx context.Context, user *User) (*User, error)
		// FindUserByEmail retrieves a user by their email address.
		FindUserByEmail(ctx context.Context, email string) (*User, error)
		// FindUserById retrieves a user by their ID from the database.
		FindUserById(ctx context.Context, id string) (*User, error)
		// FindAllEstablishmentsByUserId fetches all establishments linked to a user.
		FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]Establishment, error)
		// FindAllUsers retrieves all users from the database.
		FindAllUsers(ctx context.Context) ([]User, error)
	}

	// UserController defines the HTTP handlers for managing users.
	AccountController interface {
		// Add handles the HTTP request to create a new user.
		Add(ctx *gin.Context)
		// FindUserById handles the HTTP request to retrieve a user by ID.
		FindUserById(ctx *gin.Context)
		// FindAllUsers handles the HTTP request to return a list of all users.
		FindAllUsers(ctx *gin.Context)
		// FindAllEstablishmentsByUserId handles the HTTP request to fetch all establishments linked to a user.
		FindAllEstablishmentsByUserId(ctx *gin.Context)
	}
)

package domains

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
		Password    string       `json:"password"`
		ActivatedAt sql.NullTime `json:"activate_at"`
		CreatedAt   time.Time    `json:"created_at"`
		UpdatedAt   time.Time    `json:"updated_at"`
	}

	UserUseCase interface {
		Add(ctx context.Context, payload *User) (*User, *core.Exception)

		FindUserById(ctx context.Context, id string) (*User, *core.Exception)

		FindAllUsers(ctx context.Context) ([]User, *core.Exception)

		FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]Establishment, *core.Exception)
	}

	UserRepository interface {
		Add(ctx context.Context, user *User) (*User, error)

		FindUserByEmail(ctx context.Context, email string) (*User, error)

		FindUserById(ctx context.Context, id string) (*User, error)

		FindAllEstablishmentsByUserId(ctx context.Context, user_id string) ([]Establishment, error)

		FindAllUsers(ctx context.Context) ([]User, error)
	}

	UserController interface {
		Add(ctx *gin.Context)

		FindUserById(ctx *gin.Context)

		FindAllUsers(ctx *gin.Context)

		FindAllEstablishmentsByUserId(ctx *gin.Context)
	}
)

package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/core"
)

// represent client in the system
type (
	Client struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
	}

	ClientUseCase interface {
		// Add creates a new Client.
		Add(ctx context.Context, payload *Client) (*Client, *core.Exception)
	}

	ClientRepository interface {
		// Add creates a new Client.
		Add(ctx context.Context, payload *Client) (*Client, *core.Exception)
	}

	ClientController interface {
		// Add handles the HTTP request to create a new Client.
		Add(ctx *gin.Context)
	}
)

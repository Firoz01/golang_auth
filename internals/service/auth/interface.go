package auth

import (
	"context"
	"github.com/go-frame/internals/entity"
	"github.com/go-frame/internals/entity/proto"
	"time"
)

type Repository interface {
	LogLoginAttempt(ctx context.Context, attempt *entity.LoginAttempt) error
	CountFailedAttempts(ctx context.Context, email string, since time.Time) (int, error)
	GetFirstFailedAttemptTime(ctx context.Context, email string, since time.Time) (time.Time, error)
	UpdateLoginAttempt(ctx context.Context, attemptID int) error
}

type UsersServiceRepo interface {
	VerifyPasswordAndGetClaimData(ctx context.Context, in *proto.VerifyPasswordRequest) (*proto.VerifyPasswordResponse, error)
	GetClaimDataByUserID(ctx context.Context, in *proto.ClaimDataByUserIDRequest) (*proto.UserClaimData, error)
}

// LoginAttemptLogger is an interface for logging login attempts.
type LoginAttemptLogger interface {
	LogLoginAttempt(ctx context.Context, attempt *entity.LoginAttempt) error
	UpdateLoginAttempt(ctx context.Context, attemptID int) error
}

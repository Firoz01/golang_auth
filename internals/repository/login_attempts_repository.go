package repository

import (
	"context"
	"github.com/go-frame/internals/entity"
	"github.com/uptrace/bun"
	"time"
)

type LoginAttemptRepo struct {
	db *bun.DB
}

func NewLoginAttemptRepo(db *bun.DB) *LoginAttemptRepo {
	return &LoginAttemptRepo{
		db: db,
	}
}

func (repo *LoginAttemptRepo) LogLoginAttempt(ctx context.Context, attempt *entity.LoginAttempt) error {
	_, err := repo.db.NewInsert().Model(attempt).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *LoginAttemptRepo) CountFailedAttempts(ctx context.Context, email string, since time.Time) (int, error) {

	count, err := repo.db.NewSelect().
		Model((*entity.LoginAttempt)(nil)).
		Where("email = ?", email).
		Where("successful = ?", false).
		Where("attempt_time > ?", since).ScanAndCount(ctx)
	return count, err
}

func (repo *LoginAttemptRepo) GetFirstFailedAttemptTime(ctx context.Context, email string, since time.Time) (time.Time, error) {
	var attempt entity.LoginAttempt
	err := repo.db.NewSelect().
		Model(&attempt).
		Where("email = ?", email).
		Where("successful = ?", false).
		Where("attempt_time > ?", since).
		Order("attempt_time ASC").
		Limit(1).
		Scan(ctx)
	return attempt.AttemptTime, err
}

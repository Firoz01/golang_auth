package repository

import (
	"context"
	"fmt"
	"github.com/go-frame/config"
	"github.com/go-frame/internals/entity/proto"
	"time"
)

type UsersServiceRepo struct {
	client proto.UserServiceClient
}

func NewUsersServiceRepo(client proto.UserServiceClient) *UsersServiceRepo {
	return &UsersServiceRepo{
		client: client,
	}
}

func (repo *UsersServiceRepo) VerifyPasswordAndGetClaimData(ctx context.Context, in *proto.VerifyPasswordRequest) (*proto.VerifyPasswordResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, config.GRPC_TIMEOUT_SECS*time.Second)
	defer cancel()
	res, err := repo.client.VerifyPasswordAndGetClaimData(ctx, in)
	fmt.Print(err)
	return res, err
}

func (repo *UsersServiceRepo) GetClaimDataByUserID(ctx context.Context, in *proto.ClaimDataByUserIDRequest) (*proto.UserClaimData, error) {
	ctx, cancel := context.WithTimeout(ctx, config.GRPC_TIMEOUT_SECS*time.Second)
	defer cancel()
	res, err := repo.client.GetClaimDataByUserID(ctx, in)
	fmt.Print(err)
	return res, err
}

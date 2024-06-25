package cmd

import (
	"log"
	"net"

	"github.com/go-frame/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func setupGRPCServer(conf *config.Config, db *bun.DB) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	return s
}

func startGRPCServer(g *grpc.Server, grpcConfig config.GRPC, db *bun.DB) {
	lis, err := net.Listen("tcp", grpcConfig.GrpcPort)
	if err != nil {
		log.Fatalf("GRPC: failed to listen: %v", err)
	}

	log.Printf("GRPC: server listening at %v", lis.Addr())
	if err := g.Serve(lis); err != nil {
		log.Fatalf("GRPC: failed to serve: %v", err)
	}
}

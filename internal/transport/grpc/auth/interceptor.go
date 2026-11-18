package auth

import (
	"context"

	"github.com/dockloud/service-auth/core"
	"github.com/dockloud/service-auth/grpcauth"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func NewKeyringInterceptor(keyring Keyring) (grpc.UnaryServerInterceptor, error) {
	verifier, err := serviceauth.NewVerifier(keyring, 0)
	if err != nil {
		return nil, err
	}

	interceptor, err := grpcauth.NewServerInterceptor(verifier, 0)
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if info.FullMethod == healthpb.Health_Check_FullMethodName {
			return handler(ctx, request)
		}
		return interceptor(ctx, request, info, handler)
	}, nil
}

func NewSignInterceptor(secret []byte) (grpc.UnaryClientInterceptor, error) {
	signingKey := serviceauth.Key{
		ID:      "lnk-gateway-v1",
		Service: "dockloud-link",
		Secret:  secret,
	}

	signer, err := serviceauth.NewSigner(signingKey)
	if err != nil {
		return nil, err
	}
	interceptor, err := grpcauth.NewClientInterceptor(signer, 0)
	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

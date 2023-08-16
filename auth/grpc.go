package auth

import (
	"context"

	"github.com/mainflux/mainflux/things/policies"
)

var _ AuthServiceClient = (*grpcAuthClient)(nil)

type grpcAuthClient struct {
	authClient policies.AuthServiceClient
}

func NewGrpcAuthClient(authClient policies.AuthServiceClient) AuthServiceClient {
	return &grpcAuthClient{
		authClient: authClient,
	}
}

// Authorize implements AuthServiceClient.
func (gc *grpcAuthClient) Authorize(ctx context.Context, in *policies.AuthorizeReq) (*policies.AuthorizeRes, error) {
	return gc.authClient.Authorize(ctx, in)
}

// Identify implements AuthServiceClient.
func (gc *grpcAuthClient) Identify(ctx context.Context, in *policies.IdentifyReq) (*policies.IdentifyRes, error) {
	return gc.authClient.Identify(ctx, in)
}

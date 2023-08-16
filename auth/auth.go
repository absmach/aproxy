package auth

import (
	"context"

	"github.com/mainflux/mainflux/things/policies"
)

type AuthServiceClient interface {
	// Authorize checks if the subject is authorized to perform
	// the action on the object.
	Authorize(ctx context.Context, in *policies.AuthorizeReq) (*policies.AuthorizeRes, error)
	// Identify returns the ID of the thing has the given secret.
	Identify(ctx context.Context, in *policies.IdentifyReq) (*policies.IdentifyRes, error)
}

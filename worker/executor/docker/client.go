package docker

import (
	"context"
)

type Client interface {
	PullImage(ctx context.Context, image string) error
}

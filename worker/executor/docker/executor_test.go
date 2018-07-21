package docker

import (
	"context"
	"errors"
	"testing"

	"github.com/daisaru11/cider/worker/executor/docker/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPullImageForExistingImage(t *testing.T) {
	image := "existing:latest"
	client := &mocks.Client{}
	ctx := context.Background()

	client.On("PullImage", ctx, "existing:latest").
		Return(nil).
		Once()

	e := executor{
		ctx:    ctx,
		client: client,
	}
	err := e.pullImage(image)
	assert.Nil(t, err)
}

func TestPullImageForNotExistingImage(t *testing.T) {
	image := "not_existing:latest"

	client := &mocks.Client{}
	ctx := context.Background()

	client.On("PullImage", ctx, "not_existing:latest").
		Return(errors.New("not found")).
		Once()

	e := executor{
		ctx:    ctx,
		client: client,
	}
	err := e.pullImage(image)
	assert.Error(t, err)
}

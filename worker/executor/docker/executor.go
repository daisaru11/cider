package docker

import (
	"context"
	"log"

	"github.com/daisaru11/cider/worker/executor/common"
)

type executor struct {
	client Client
	ctx    context.Context
}

// NewExecutor creates a new executor instance
func NewExecutor(ctx context.Context) (common.Executor, error) {
	client, err := NewClientImpl()
	if err != nil {
		return nil, err
	}

	e := &executor{
		client: client,
		ctx:    ctx,
	}
	return e, nil
}

func (e *executor) Run() {
	log.Print("docker executor Run Start")

	image := "golang:1.10-alpine"
	e.pullImage(image)
	log.Print("docker executor Run End")
}

func (e *executor) pullImage(image string) error {
	err := e.client.PullImage(e.ctx, image)
	if err != nil {
		return err
	}

	return nil
}

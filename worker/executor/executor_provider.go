package executor

import (
	"context"

	"github.com/daisaru11/cider/worker/executor/common"
	"github.com/daisaru11/cider/worker/executor/docker"
)

// ExecutorProvider instantiates executor instances
type ExecutorProvider struct {
}

// NewExecutorProvider creates a new provider instance
func NewExecutorProvider() *ExecutorProvider {
	return &ExecutorProvider{}
}

// Provide returns a executor implementation
func (p *ExecutorProvider) Provide(ctx context.Context) (common.Executor, error) {
	return docker.NewExecutor(ctx)
}

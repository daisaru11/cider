package main

import (
	"context"
	"log"

	"github.com/daisaru11/cider/worker/executor"
)

func main() {
	log.Print("Start")

	provider := executor.NewExecutorProvider()

	runContext, runCancel := context.WithCancel(context.Background())
	defer runCancel()

	executor, error := provider.Provide(runContext)
	if error != nil {
		return
	}
	executor.Run()

	log.Print("End")
}

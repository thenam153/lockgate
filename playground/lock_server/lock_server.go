package main

import (
	"fmt"
	"os"

	"github.com/thenam153/lockgate/pkg/distributed_locker/optimistic_locking_store"

	"github.com/thenam153/lockgate/pkg/distributed_locker"
)

func run() error {
	store := optimistic_locking_store.NewInMemoryStore()
	backend := distributed_locker.NewOptimisticLockingStorageBasedBackend(store)
	return distributed_locker.RunHttpBackendServer("0.0.0.0", "55589", backend)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

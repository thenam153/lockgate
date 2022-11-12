package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/werf/kubedog/pkg/kube"

	"lockgate/pkg/distributed_locker/optimistic_locking_store"

	"lockgate/pkg/distributed_locker"
)

func run() error {
	if err := kube.Init(kube.InitOptions{}); err != nil {
		return err
	}

	store := optimistic_locking_store.NewKubernetesResourceAnnotationsStore(
		kube.DynamicClient, schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "configmaps",
		}, "mycm", "myns",
	)
	backend := distributed_locker.NewOptimisticLockingStorageBasedBackend(store)
	return distributed_locker.RunHttpBackendServer("0.0.0.0", "55589", backend)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

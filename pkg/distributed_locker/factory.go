package distributed_locker

import (
	"github.com/thenam153/lockgate/pkg/distributed_locker/optimistic_locking_store"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func NewKubernetesLocker(kubernetesInterface dynamic.Interface, gvr schema.GroupVersionResource, resourceName string, namespace string) *DistributedLocker {
	store := optimistic_locking_store.NewKubernetesResourceAnnotationsStore(kubernetesInterface, gvr, resourceName, namespace)
	backend := NewOptimisticLockingStorageBasedBackend(store)
	return NewDistributedLocker(backend)
}

func NewHttpLocker(urlEndpoint string) *DistributedLocker {
	backend := NewHttpBackend(urlEndpoint)
	return NewDistributedLocker(backend)
}

func NewHttpBackendHandlerWithInMemoryStore() *HttpBackendHandler {
	store := optimistic_locking_store.NewInMemoryStore()
	backend := NewOptimisticLockingStorageBasedBackend(store)
	return NewHttpBackendHandler(backend)
}

func NewHttpBackendHandlerWithKubernetesStore(kubernetesInterface dynamic.Interface, gvr schema.GroupVersionResource, resourceName string, namespace string) *HttpBackendHandler {
	store := optimistic_locking_store.NewKubernetesResourceAnnotationsStore(kubernetesInterface, gvr, resourceName, namespace)
	backend := NewOptimisticLockingStorageBasedBackend(store)
	return NewHttpBackendHandler(backend)
}

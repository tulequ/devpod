// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// SpaceInstanceLister helps list SpaceInstances.
// All objects returned here must be treated as read-only.
type SpaceInstanceLister interface {
	// List lists all SpaceInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1.SpaceInstance, err error)
	// SpaceInstances returns an object that can list and get SpaceInstances.
	SpaceInstances(namespace string) SpaceInstanceNamespaceLister
	SpaceInstanceListerExpansion
}

// spaceInstanceLister implements the SpaceInstanceLister interface.
type spaceInstanceLister struct {
	listers.ResourceIndexer[*storagev1.SpaceInstance]
}

// NewSpaceInstanceLister returns a new SpaceInstanceLister.
func NewSpaceInstanceLister(indexer cache.Indexer) SpaceInstanceLister {
	return &spaceInstanceLister{listers.New[*storagev1.SpaceInstance](indexer, storagev1.Resource("spaceinstance"))}
}

// SpaceInstances returns an object that can list and get SpaceInstances.
func (s *spaceInstanceLister) SpaceInstances(namespace string) SpaceInstanceNamespaceLister {
	return spaceInstanceNamespaceLister{listers.NewNamespaced[*storagev1.SpaceInstance](s.ResourceIndexer, namespace)}
}

// SpaceInstanceNamespaceLister helps list and get SpaceInstances.
// All objects returned here must be treated as read-only.
type SpaceInstanceNamespaceLister interface {
	// List lists all SpaceInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1.SpaceInstance, err error)
	// Get retrieves the SpaceInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*storagev1.SpaceInstance, error)
	SpaceInstanceNamespaceListerExpansion
}

// spaceInstanceNamespaceLister implements the SpaceInstanceNamespaceLister
// interface.
type spaceInstanceNamespaceLister struct {
	listers.ResourceIndexer[*storagev1.SpaceInstance]
}

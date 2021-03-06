// Code generated by main. DO NOT EDIT.

package fake

import (
	v1 "github.com/ysicing/k3sbuild/pkg/generated/clientset/versioned/typed/k3s.cattle.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeK3sV1 struct {
	*testing.Fake
}

func (c *FakeK3sV1) Addons(namespace string) v1.AddonInterface {
	return &FakeAddons{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeK3sV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

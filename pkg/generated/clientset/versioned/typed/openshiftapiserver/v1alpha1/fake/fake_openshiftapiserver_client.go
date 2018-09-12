// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/cluster-openshift-apiserver-operator/pkg/generated/clientset/versioned/typed/openshiftapiserver/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOpenshiftapiserverV1alpha1 struct {
	*testing.Fake
}

func (c *FakeOpenshiftapiserverV1alpha1) KubeApiserverOperatorConfigs() v1alpha1.KubeApiserverOperatorConfigInterface {
	return &FakeKubeApiserverOperatorConfigs{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOpenshiftapiserverV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenShiftAPIServers implements OpenShiftAPIServerInterface
type FakeOpenShiftAPIServers struct {
	Fake *FakeOperatorV1
}

var openshiftapiserversResource = v1.SchemeGroupVersion.WithResource("openshiftapiservers")

var openshiftapiserversKind = v1.SchemeGroupVersion.WithKind("OpenShiftAPIServer")

// Get takes name of the openShiftAPIServer, and returns the corresponding openShiftAPIServer object, and an error if there is any.
func (c *FakeOpenShiftAPIServers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(openshiftapiserversResource, name), &v1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OpenShiftAPIServer), err
}

// List takes label and field selectors, and returns the list of OpenShiftAPIServers that match those selectors.
func (c *FakeOpenShiftAPIServers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OpenShiftAPIServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(openshiftapiserversResource, openshiftapiserversKind, opts), &v1.OpenShiftAPIServerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OpenShiftAPIServerList{ListMeta: obj.(*v1.OpenShiftAPIServerList).ListMeta}
	for _, item := range obj.(*v1.OpenShiftAPIServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openShiftAPIServers.
func (c *FakeOpenShiftAPIServers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(openshiftapiserversResource, opts))
}

// Create takes the representation of a openShiftAPIServer and creates it.  Returns the server's representation of the openShiftAPIServer, and an error, if there is any.
func (c *FakeOpenShiftAPIServers) Create(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.CreateOptions) (result *v1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(openshiftapiserversResource, openShiftAPIServer), &v1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OpenShiftAPIServer), err
}

// Update takes the representation of a openShiftAPIServer and updates it. Returns the server's representation of the openShiftAPIServer, and an error, if there is any.
func (c *FakeOpenShiftAPIServers) Update(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.UpdateOptions) (result *v1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(openshiftapiserversResource, openShiftAPIServer), &v1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OpenShiftAPIServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenShiftAPIServers) UpdateStatus(ctx context.Context, openShiftAPIServer *v1.OpenShiftAPIServer, opts metav1.UpdateOptions) (*v1.OpenShiftAPIServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(openshiftapiserversResource, "status", openShiftAPIServer), &v1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OpenShiftAPIServer), err
}

// Delete takes name of the openShiftAPIServer and deletes it. Returns an error if one occurs.
func (c *FakeOpenShiftAPIServers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(openshiftapiserversResource, name, opts), &v1.OpenShiftAPIServer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenShiftAPIServers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(openshiftapiserversResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OpenShiftAPIServerList{})
	return err
}

// Patch applies the patch and returns the patched openShiftAPIServer.
func (c *FakeOpenShiftAPIServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OpenShiftAPIServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(openshiftapiserversResource, name, pt, data, subresources...), &v1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OpenShiftAPIServer), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied openShiftAPIServer.
func (c *FakeOpenShiftAPIServers) Apply(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OpenShiftAPIServer, err error) {
	if openShiftAPIServer == nil {
		return nil, fmt.Errorf("openShiftAPIServer provided to Apply must not be nil")
	}
	data, err := json.Marshal(openShiftAPIServer)
	if err != nil {
		return nil, err
	}
	name := openShiftAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("openShiftAPIServer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(openshiftapiserversResource, *name, types.ApplyPatchType, data), &v1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OpenShiftAPIServer), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeOpenShiftAPIServers) ApplyStatus(ctx context.Context, openShiftAPIServer *operatorv1.OpenShiftAPIServerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OpenShiftAPIServer, err error) {
	if openShiftAPIServer == nil {
		return nil, fmt.Errorf("openShiftAPIServer provided to Apply must not be nil")
	}
	data, err := json.Marshal(openShiftAPIServer)
	if err != nil {
		return nil, err
	}
	name := openShiftAPIServer.Name
	if name == nil {
		return nil, fmt.Errorf("openShiftAPIServer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(openshiftapiserversResource, *name, types.ApplyPatchType, data, "status"), &v1.OpenShiftAPIServer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.OpenShiftAPIServer), err
}
package genericoperatorclient

import (
	"context"

	"github.com/imdario/mergo"

	operatorv1 "github.com/openshift/api/operator/v1"
	applyoperatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	"github.com/openshift/library-go/pkg/apiserver/jsonpatch"
	"github.com/openshift/library-go/pkg/operator/v1helpers"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/rest"
	"k8s.io/utils/clock"
)

func NewStaticPodOperatorClient(clock clock.PassiveClock, config *rest.Config, gvr schema.GroupVersionResource, gvk schema.GroupVersionKind, extractApplySpec StaticPodOperatorSpecExtractorFunc, extractApplyStatus StaticPodOperatorStatusExtractorFunc) (v1helpers.StaticPodOperatorClient, dynamicinformer.DynamicSharedInformerFactory, error) {
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return newClusterScopedOperatorClient(clock, dynamicClient, gvr, gvk, defaultConfigName,
		extractApplySpec, extractApplyStatus)
}

func (c dynamicOperatorClient) GetStaticPodOperatorState() (*operatorv1.StaticPodOperatorSpec, *operatorv1.StaticPodOperatorStatus, string, error) {
	uncastInstance, err := c.informer.Lister().Get("cluster")
	if err != nil {
		return nil, nil, "", err
	}
	instance := uncastInstance.(*unstructured.Unstructured)

	return getStaticPodOperatorStateFromInstance(instance)
}

func getStaticPodOperatorStateFromInstance(instance *unstructured.Unstructured) (*operatorv1.StaticPodOperatorSpec, *operatorv1.StaticPodOperatorStatus, string, error) {
	spec, err := getStaticPodOperatorSpecFromUnstructured(instance.UnstructuredContent())
	if err != nil {
		return nil, nil, "", err
	}
	status, err := getStaticPodOperatorStatusFromUnstructured(instance.UnstructuredContent())
	if err != nil {
		return nil, nil, "", err
	}

	return spec, status, instance.GetResourceVersion(), nil
}

func (c dynamicOperatorClient) GetStaticPodOperatorStateWithQuorum(ctx context.Context) (*operatorv1.StaticPodOperatorSpec, *operatorv1.StaticPodOperatorStatus, string, error) {
	instance, err := c.client.Get(ctx, "cluster", metav1.GetOptions{})
	if err != nil {
		return nil, nil, "", err
	}

	return getStaticPodOperatorStateFromInstance(instance)
}

func (c dynamicOperatorClient) UpdateStaticPodOperatorSpec(ctx context.Context, resourceVersion string, spec *operatorv1.StaticPodOperatorSpec) (*operatorv1.StaticPodOperatorSpec, string, error) {
	uncastOriginal, err := c.informer.Lister().Get("cluster")
	if err != nil {
		return nil, "", err
	}
	original := uncastOriginal.(*unstructured.Unstructured)

	copy := original.DeepCopy()
	copy.SetResourceVersion(resourceVersion)
	if err := setStaticPodOperatorSpecFromUnstructured(copy.UnstructuredContent(), spec); err != nil {
		return nil, "", err
	}

	ret, err := c.client.Update(ctx, copy, metav1.UpdateOptions{})
	if err != nil {
		return nil, "", err
	}
	retSpec, err := getStaticPodOperatorSpecFromUnstructured(ret.UnstructuredContent())
	if err != nil {
		return nil, "", err
	}

	return retSpec, ret.GetResourceVersion(), nil
}

func (c dynamicOperatorClient) UpdateStaticPodOperatorStatus(ctx context.Context, resourceVersion string, status *operatorv1.StaticPodOperatorStatus) (*operatorv1.StaticPodOperatorStatus, error) {
	uncastOriginal, err := c.informer.Lister().Get("cluster")
	if err != nil {
		return nil, err
	}
	original := uncastOriginal.(*unstructured.Unstructured)

	copy := original.DeepCopy()
	copy.SetResourceVersion(resourceVersion)
	if err := setStaticPodOperatorStatusFromUnstructured(copy.UnstructuredContent(), status); err != nil {
		return nil, err
	}

	ret, err := c.client.UpdateStatus(ctx, copy, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	retStatus, err := getStaticPodOperatorStatusFromUnstructured(ret.UnstructuredContent())
	if err != nil {
		return nil, err
	}

	return retStatus, nil
}

func (c dynamicOperatorClient) ApplyStaticPodOperatorSpec(ctx context.Context, fieldManager string, desiredConfiguration *applyoperatorv1.StaticPodOperatorSpecApplyConfiguration) (err error) {
	return c.applyOperatorSpec(ctx, fieldManager, desiredConfiguration)
}

func (c dynamicOperatorClient) ApplyStaticPodOperatorStatus(ctx context.Context, fieldManager string, desiredConfiguration *applyoperatorv1.StaticPodOperatorStatusApplyConfiguration) (err error) {
	return c.applyOperatorStatus(ctx, fieldManager, desiredConfiguration)
}

func (c dynamicOperatorClient) PatchStaticOperatorStatus(ctx context.Context, jsonPatch *jsonpatch.PatchSet) (err error) {
	return c.patchOperatorStatus(ctx, jsonPatch)
}

func getStaticPodOperatorSpecFromUnstructured(obj map[string]interface{}) (*operatorv1.StaticPodOperatorSpec, error) {
	uncastSpec, exists, err := unstructured.NestedMap(obj, "spec")
	if !exists {
		return &operatorv1.StaticPodOperatorSpec{}, nil
	}
	if err != nil {
		return nil, err
	}

	ret := &operatorv1.StaticPodOperatorSpec{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(uncastSpec, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func setStaticPodOperatorSpecFromUnstructured(obj map[string]interface{}, spec *operatorv1.StaticPodOperatorSpec) error {
	// we cannot simply set the entire map because doing so would stomp unknown fields, like say a static pod operator spec when cast as an operator spec
	newUnstructuredSpec, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return err
	}

	originalUnstructuredSpec, exists, err := unstructured.NestedMap(obj, "spec")
	if !exists {
		return unstructured.SetNestedMap(obj, newUnstructuredSpec, "spec")
	}
	if err != nil {
		return err
	}
	if err := mergo.Merge(&originalUnstructuredSpec, newUnstructuredSpec, mergo.WithOverride); err != nil {
		return err
	}

	return unstructured.SetNestedMap(obj, originalUnstructuredSpec, "spec")
}

func getStaticPodOperatorStatusFromUnstructured(obj map[string]interface{}) (*operatorv1.StaticPodOperatorStatus, error) {
	uncastStatus, exists, err := unstructured.NestedMap(obj, "status")
	if !exists {
		return &operatorv1.StaticPodOperatorStatus{}, nil
	}
	if err != nil {
		return nil, err
	}

	ret := &operatorv1.StaticPodOperatorStatus{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(uncastStatus, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func setStaticPodOperatorStatusFromUnstructured(obj map[string]interface{}, spec *operatorv1.StaticPodOperatorStatus) error {
	// we cannot simply set the entire map because doing so would stomp unknown fields, like say a static pod operator spec when cast as an operator spec
	newUnstructuredStatus, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return err
	}

	originalUnstructuredStatus, exists, err := unstructured.NestedMap(obj, "status")
	if !exists {
		return unstructured.SetNestedMap(obj, newUnstructuredStatus, "status")
	}
	if err != nil {
		return err
	}
	if err := mergo.Merge(&originalUnstructuredStatus, newUnstructuredStatus, mergo.WithOverride); err != nil {
		return err
	}

	return unstructured.SetNestedMap(obj, originalUnstructuredStatus, "status")
}

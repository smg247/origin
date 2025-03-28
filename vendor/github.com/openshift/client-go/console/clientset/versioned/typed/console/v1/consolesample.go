// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	consolev1 "github.com/openshift/api/console/v1"
	applyconfigurationsconsolev1 "github.com/openshift/client-go/console/applyconfigurations/console/v1"
	scheme "github.com/openshift/client-go/console/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ConsoleSamplesGetter has a method to return a ConsoleSampleInterface.
// A group's client should implement this interface.
type ConsoleSamplesGetter interface {
	ConsoleSamples() ConsoleSampleInterface
}

// ConsoleSampleInterface has methods to work with ConsoleSample resources.
type ConsoleSampleInterface interface {
	Create(ctx context.Context, consoleSample *consolev1.ConsoleSample, opts metav1.CreateOptions) (*consolev1.ConsoleSample, error)
	Update(ctx context.Context, consoleSample *consolev1.ConsoleSample, opts metav1.UpdateOptions) (*consolev1.ConsoleSample, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*consolev1.ConsoleSample, error)
	List(ctx context.Context, opts metav1.ListOptions) (*consolev1.ConsoleSampleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *consolev1.ConsoleSample, err error)
	Apply(ctx context.Context, consoleSample *applyconfigurationsconsolev1.ConsoleSampleApplyConfiguration, opts metav1.ApplyOptions) (result *consolev1.ConsoleSample, err error)
	ConsoleSampleExpansion
}

// consoleSamples implements ConsoleSampleInterface
type consoleSamples struct {
	*gentype.ClientWithListAndApply[*consolev1.ConsoleSample, *consolev1.ConsoleSampleList, *applyconfigurationsconsolev1.ConsoleSampleApplyConfiguration]
}

// newConsoleSamples returns a ConsoleSamples
func newConsoleSamples(c *ConsoleV1Client) *consoleSamples {
	return &consoleSamples{
		gentype.NewClientWithListAndApply[*consolev1.ConsoleSample, *consolev1.ConsoleSampleList, *applyconfigurationsconsolev1.ConsoleSampleApplyConfiguration](
			"consolesamples",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *consolev1.ConsoleSample { return &consolev1.ConsoleSample{} },
			func() *consolev1.ConsoleSampleList { return &consolev1.ConsoleSampleList{} },
		),
	}
}
